package builder

import (
	"net/http"
	"web_proxy/config"
	"web_proxy/contract"
)

// http handler function interface
type Handler func(http.ResponseWriter, *http.Request)

// http subdomain
type httpSub struct {
	service config.Service
	log     contract.Logger
}

// builder http service
type BuilderHttp map[string]Handler

// create new service handler (provider)
func New(conf config.ConfigWebProxy, log contract.Logger) BuilderHttp {
	builderHttp := make(BuilderHttp)
	for name, v := range conf.Service {
		log.Info("setup service: ", name)
		http := httpSub{service: v, log: log}
		builderHttp[v.ServiceUrl] = http.handelService()
	}
	return builderHttp
}
