package interactor

import (
	"net/http"
	"web_proxy/config"
	"web_proxy/contract"
)

type WebProxy struct {
	// data request ( header, url, method, body )
	req *http.Request

	// sender request, in this step setup proxy
	client *http.Client

	// host url (ex: http://example.com)
	mainUrl string

	// host url with path (ex: http://example.com/test)
	reqUrl string

	// proxy list and size of
	proxyList     []config.Proxy
	sizeListProxy int

	log contract.Logger
}

// create new web proxy interactor
func New(log contract.Logger, service config.Service) *WebProxy {
	return &WebProxy{
		log:           log,
		reqUrl:        service.ServiceDomain,
		mainUrl:       service.ServiceDomain,
		proxyList:     service.ProxyList,
		sizeListProxy: len(service.ProxyList),
	}
}
