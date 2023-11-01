package interactor

import (
	"net/http"
	"web_proxy/config"
	"web_proxy/core/contract"
)

type WebPage struct {
	req    *http.Request
	client *http.Client

	header http.Header

	urlDomain,
	url string

	proxyList     []config.Proxy
	sizeListProxy int

	log contract.Logger
}

func New(log contract.Logger, url string, proxyList []config.Proxy) *WebPage {
	return &WebPage{
		log:           log,
		url:           url,
		urlDomain:     url,
		sizeListProxy: len(proxyList),
	}
}
