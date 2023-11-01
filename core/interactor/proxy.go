package interactor

import (
	"crypto/tls"
	"encoding/base64"
	"math/rand"
	"net/http"
	"net/url"
	"web_proxy/config"
)

type ProxyFunc func(req *http.Request) (*url.URL, error)

func (_ WebPage) setupProxy(proxyURL string) ProxyFunc {
	return func(_ *http.Request) (*url.URL, error) {
		u, err := url.Parse(proxyURL)
		return u, err
	}
}

func (wp WebPage) chooseProxyAdrr() *config.Proxy {
	if wp.sizeListProxy < 1 {
		return nil
	}
	n := rand.Intn(int(wp.sizeListProxy))
	return &wp.proxyList[n]
}

func (wp *WebPage) SetupProxy() *WebPage {
	proxyConf := wp.chooseProxyAdrr()
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}
	wp.client = client

	if proxyConf == nil {
		wp.log.Info("setup proxy client | No set proxy, proxy list is empty ")
		return wp
	}

	wp.log.Info("setup proxy client | with: ", proxyConf.Addr)
	transport.ProxyConnectHeader = http.Header{}
	userPass := proxyConf.Auth.Username + ":" + proxyConf.Auth.Password
	userPass = "Basic " + base64.StdEncoding.EncodeToString([]byte(userPass))
	transport.ProxyConnectHeader.Add("Proxy-Authorization", userPass)
	transport.Proxy = wp.setupProxy(proxyConf.Addr)
	wp.header.Add("Proxy-Authorization", userPass)
	return wp
}
