package interactor

import (
	"crypto/tls"
	"math/rand"
	"net/http"
	"net/url"
	"time"
	"web_proxy/config"
)

// skip verify tls if is not valid
var tlsConf = &tls.Config{InsecureSkipVerify: true}

// chose random proxy form list
//
// if size of proxy is zero return nil
func (wp WebProxy) chooseProxyAdrr() *config.Proxy {
	if wp.sizeListProxy < 1 {
		return nil
	}
	n := rand.Intn(int(wp.sizeListProxy))
	return &wp.proxyList[n]
}

// setup new proxy if exists
func (wp *WebProxy) SetProxy() *WebProxy {
	transport := &http.Transport{
		TLSClientConfig:       tlsConf,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	wp.client.Transport = transport

	proxyConf := wp.chooseProxyAdrr()
	if proxyConf == nil {
		wp.log.Info("setup proxy client | No set proxy, proxy list is empty ")
		return wp
	}

	proxyUrl, err := url.Parse(proxyConf.Addr)
	if err != nil {
		return wp
	}

	wp.log.Info("setup proxy client | ", proxyConf.Addr)
	proxyUrl.User = url.UserPassword(proxyConf.Username, proxyConf.Password)
	transport.Proxy = http.ProxyURL(proxyUrl)
	return wp
}
