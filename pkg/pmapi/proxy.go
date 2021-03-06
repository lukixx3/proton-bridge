// Copyright (c) 2020 Proton Technologies AG
//
// This file is part of ProtonMail Bridge.
//
// ProtonMail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ProtonMail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ProtonMail Bridge.  If not, see <https://www.gnu.org/licenses/>.

package pmapi

import (
	"crypto/tls"
	"encoding/base64"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/miekg/dns"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	proxyRevertTime    = 24 * time.Hour
	proxySearchTimeout = 30 * time.Second
	proxyQueryTimeout  = 10 * time.Second
	proxyLookupWait    = 5 * time.Second
	proxyQuery         = "dMFYGSLTQOJXXI33ONVQWS3BOMNUA.protonpro.xyz"
)

var dohProviders = []string{ //nolint[gochecknoglobals]
	"https://dns11.quad9.net/dns-query",
	"https://dns.google/dns-query",
}

// globalAllowDoH controls whether or not to enable use of DoH/Proxy in pmapi.
var globalAllowDoH = false // nolint[golint]

// globalProxyMutex allows threadsafe modification of proxy state.
var globalProxyMutex = sync.RWMutex{} // nolint[golint]

// globalOriginalURL backs up the original API url so it can be restored later.
var globalOriginalURL = RootURL // nolint[golint]

// globalIsDoHAllowed returns whether or not to use DoH.
func globalIsDoHAllowed() bool { // nolint[golint]
	globalProxyMutex.RLock()
	defer globalProxyMutex.RUnlock()

	return globalAllowDoH
}

// GlobalAllowDoH enables DoH.
func GlobalAllowDoH() { // nolint[golint]
	globalProxyMutex.Lock()
	defer globalProxyMutex.Unlock()

	globalAllowDoH = true
}

// GlobalDisallowDoH disables DoH and sets the RootURL back to what it was.
func GlobalDisallowDoH() { // nolint[golint]
	globalProxyMutex.Lock()
	defer globalProxyMutex.Unlock()

	globalAllowDoH = false
	RootURL = globalOriginalURL
}

// globalSetRootURL sets the global RootURL.
func globalSetRootURL(url string) { // nolint[golint]
	globalProxyMutex.Lock()
	defer globalProxyMutex.Unlock()

	RootURL = url
}

// GlobalGetRootURL returns the global RootURL.
func GlobalGetRootURL() (url string) { // nolint[golint]
	globalProxyMutex.RLock()
	defer globalProxyMutex.RUnlock()

	return RootURL
}

// isProxyEnabled returns whether or not we are currently using a proxy.
func isProxyEnabled() bool { // nolint[golint]
	return globalOriginalURL != GlobalGetRootURL()
}

// proxyManager manages known proxies.
type proxyManager struct {
	// dohLookup is used to look up the given query at the given DoH provider, returning the TXT records>
	dohLookup func(query, provider string) (urls []string, err error)

	providers  []string // List of known doh providers.
	query      string   // The query string used to find proxies.
	proxyCache []string // All known proxies, cached in case DoH providers are unreachable.

	useDuration                time.Duration // How much time to use the proxy before returning to the original API.
	findTimeout, lookupTimeout time.Duration // Timeouts for DNS query and proxy search.

	lastLookup time.Time // The time at which we last attempted to find a proxy.
}

// newProxyManager creates a new proxyManager that queries the given DoH providers
// to retrieve DNS records for the given query string.
func newProxyManager(providers []string, query string) (p *proxyManager) { // nolint[unparam]
	p = &proxyManager{
		providers:     providers,
		query:         query,
		useDuration:   proxyRevertTime,
		findTimeout:   proxySearchTimeout,
		lookupTimeout: proxyQueryTimeout,
	}

	// Use the default DNS lookup method; this can be overridden if necessary.
	p.dohLookup = p.defaultDoHLookup

	return
}

// findProxy returns a new proxy domain which is not equal to the current RootURL.
// It returns an error if the process takes longer than ProxySearchTime.
func (p *proxyManager) findProxy() (proxy string, err error) {
	if time.Now().Before(p.lastLookup.Add(proxyLookupWait)) {
		return "", errors.New("not looking for a proxy, too soon")
	}

	p.lastLookup = time.Now()

	proxyResult := make(chan string)
	errResult := make(chan error)
	go func() {
		if err = p.refreshProxyCache(); err != nil {
			logrus.WithError(err).Warn("Failed to refresh proxy cache, cache may be out of date")
		}

		for _, proxy := range p.proxyCache {
			if proxy != stripProtocol(GlobalGetRootURL()) && p.canReach(proxy) {
				proxyResult <- proxy
				return
			}
		}

		errResult <- errors.New("no proxy available")
	}()

	select {
	case <-time.After(p.findTimeout):
		logrus.Error("Timed out finding a proxy server")
		return "", errors.New("timed out finding a proxy")

	case proxy = <-proxyResult:
		logrus.WithField("proxy", proxy).Info("Found proxy server")
		return

	case err = <-errResult:
		logrus.WithError(err).Error("Failed to find available proxy server")
		return
	}
}

// useProxy sets the proxy server to use. It returns to the original RootURL after 24 hours.
func (p *proxyManager) useProxy(proxy string) {
	if !isProxyEnabled() {
		p.disableProxyAfter(p.useDuration)
	}

	globalSetRootURL(https(proxy))
}

// disableProxyAfter disables the proxy after the given amount of time.
func (p *proxyManager) disableProxyAfter(d time.Duration) {
	go func() {
		<-time.After(d)
		globalSetRootURL(globalOriginalURL)
	}()
}

// refreshProxyCache loads the latest proxies from the known providers.
func (p *proxyManager) refreshProxyCache() error {
	logrus.Info("Refreshing proxy cache")

	for _, provider := range p.providers {
		if proxies, err := p.dohLookup(p.query, provider); err == nil {
			p.proxyCache = proxies

			// We also want to allow bridge to switch back to the standard API at any time.
			p.proxyCache = append(p.proxyCache, globalOriginalURL)

			logrus.WithField("proxies", proxies).Info("Available proxies")

			return nil
		}
	}

	return errors.New("lookup failed with all DoH providers")
}

// canReach returns whether we can reach the given url.
// NOTE: we skip cert verification to stop it complaining that cert name doesn't match hostname.
func (p *proxyManager) canReach(url string) bool {
	pinger := resty.New().
		SetHostURL(https(url)).
		SetTimeout(p.lookupTimeout).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) // nolint[gosec]

	if _, err := pinger.R().Get("/tests/ping"); err != nil {
		return false
	}

	return true
}

// defaultDoHLookup is the default implementation of the proxy manager's DoH lookup.
// It looks up DNS TXT records for the given query URL using the given DoH provider.
// It returns a list of all found TXT records.
// If the whole process takes more than ProxyQueryTime then an error is returned.
func (p *proxyManager) defaultDoHLookup(query, dohProvider string) (data []string, err error) {
	dataResult := make(chan []string)
	errResult := make(chan error)
	go func() {
		// Build new DNS request in RFC1035 format.
		dnsRequest := new(dns.Msg).SetQuestion(dns.Fqdn(query), dns.TypeTXT)

		// Pack the DNS request message into wire format.
		rawRequest, err := dnsRequest.Pack()
		if err != nil {
			errResult <- errors.Wrap(err, "failed to pack DNS request")
			return
		}

		// Encode wire-format DNS request message as base64url (RFC4648) without padding chars.
		encodedRequest := base64.RawURLEncoding.EncodeToString(rawRequest)

		// Make DoH request to the given DoH provider.
		rawResponse, err := resty.New().R().SetQueryParam("dns", encodedRequest).Get(dohProvider)
		if err != nil {
			errResult <- errors.Wrap(err, "failed to make DoH request")
			return
		}

		// Unpack the DNS response.
		dnsResponse := new(dns.Msg)
		if err = dnsResponse.Unpack(rawResponse.Body()); err != nil {
			errResult <- errors.Wrap(err, "failed to unpack DNS response")
			return
		}

		// Pick out the TXT answers.
		for _, answer := range dnsResponse.Answer {
			if t, ok := answer.(*dns.TXT); ok {
				data = append(data, t.Txt...)
			}
		}

		dataResult <- data
	}()

	select {
	case <-time.After(p.lookupTimeout):
		logrus.WithField("provider", dohProvider).Error("Timed out querying DNS records")
		return []string{}, errors.New("timed out querying DNS records")

	case data = <-dataResult:
		logrus.WithField("data", data).Info("Received TXT records")
		return

	case err = <-errResult:
		logrus.WithField("provider", dohProvider).WithError(err).Error("Failed to query DNS records")
		return
	}
}

func stripProtocol(url string) string {
	if strings.HasPrefix(url, "https://") {
		return strings.TrimPrefix(url, "https://")
	}

	if strings.HasPrefix(url, "http://") {
		return strings.TrimPrefix(url, "http://")
	}

	return url
}

func https(url string) string {
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "https://" + url
	}

	return url
}
