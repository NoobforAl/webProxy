package http

import (
	"net"
	"net/http"
	"strings"
	"web_proxy/http/builder"
)

func (s Server) setUpRouter() {
	s.log.Info("Setup subdomain")
	builder := builder.New(s.config, s.log)
	s.http.HandleFunc("/", s.handleSubdomain(builder))
}

// handle subdomain
func (s Server) handleSubdomain(bu builder.BuilderHttp) builder.Handler {
	s.log.Info("Setup Subdomain checker")
	defaultService := s.defaultService()

	return func(w http.ResponseWriter, r *http.Request) {
		hostParts := strings.Split(r.Host, ":")
		subdomain := hostParts[0]
		path := r.URL.RequestURI()
		ip := s.getIp(r)

		s.log.Infof("%s | %s | %s | %s", ip, r.Method, subdomain, path)
		switch httpHandler, ok := bu[subdomain]; ok {
		case true:
			httpHandler(w, r)

		default:
			defaultService(w, r)
		}
	}
}

// get ip client
func (_ Server) getIp(r *http.Request) (ip string) {
	ip = r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = r.RemoteAddr
	}

	ip, _, _ = net.SplitHostPort(ip)
	return ip
}
