package interactor

import (
	"crypto/tls"
	"encoding/base64"
	"math/rand"
	"net/http"
	"net/url"
	"web_proxy/config"
)

// type func proxy for returned function
type proxyFunc func(req *http.Request) (*url.URL, error)

// pars address proxy to ulr struct
func (_ WebProxy) parsProxyUrl(proxyURL string) proxyFunc {
	return func(_ *http.Request) (*url.URL, error) {
		u, err := url.Parse(proxyURL)
		return u, err
	}
}

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
func (wp *WebProxy) SetupProxy() *WebProxy {
	// skip verify tls if is not valid
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	wp.client.Transport = transport

	proxyConf := wp.chooseProxyAdrr()
	if proxyConf == nil {
		wp.log.Info("setup proxy client | No set proxy, proxy list is empty ")
		return wp
	}

	wp.log.Info("setup proxy client | ", proxyConf.Addr)
	userPass := proxyConf.Username + ":" + proxyConf.Password
	userPass = "Basic " + base64.StdEncoding.EncodeToString([]byte(userPass))

	wp.req.Header.Add("Proxy-Authorization", userPass)
	transport.Proxy = wp.parsProxyUrl(proxyConf.Addr)
	return wp
}
