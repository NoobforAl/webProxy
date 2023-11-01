package http

import (
	"bytes"
	"html/template"
	"net"
	"net/http"
	"strings"
	"web_proxy/config"
	"web_proxy/core/contract"
	"web_proxy/http/builder"
)

type Server struct {
	http   *http.ServeMux
	config config.ConfigWebProxy
	log    contract.Logger
}

func New(log contract.Logger, conf config.ConfigWebProxy) *Server {
	ser := &Server{
		http:   http.NewServeMux(),
		config: conf,
		log:    log,
	}

	return ser
}

func (s Server) setUpMainRout() {
	s.log.Info("Setup subdomain")

	conf := []*config.UrlRequester{}
	for k, v := range s.config.Service {
		s.log.Info("setup service: ", k)
		conf = append(conf, &config.UrlRequester{
			Name:      v.Url,
			Domain:    v.Domain,
			ProxyList: v.Proxy,
		})
	}

	builder := builder.New(conf, s.log)
	s.http.HandleFunc("/", s.handleSubdomain(builder))
}

func (s Server) defaultService() builder.Handelr {

	html, err := template.New("template").Funcs(template.FuncMap{
		"url": func(u string) string {
			port := strings.Split(s.config.Listen, ":")[1]
			return "http://" + u + ":" + port
		},
	}).Parse(temp)
	if err != nil {
		s.log.Panic(err)
	}

	var buf bytes.Buffer
	err = html.Execute(&buf, s.config)
	if err != nil {
		s.log.Panic(err)
	}

	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		_, _ = w.Write(buf.Bytes())
	}
}

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

func (s Server) handleSubdomain(bu builder.BuilderHttp) builder.Handelr {
	s.log.Info("Setup Subdomain checker")
	defaultService := s.defaultService()

	return func(w http.ResponseWriter, r *http.Request) {
		hostParts := strings.Split(r.Host, ":")
		subdomain := hostParts[0]
		path := r.URL.RequestURI()
		ip := s.getIp(r)

		s.log.Infof("%s | %s | %s | %s", ip, r.Method, subdomain, path)

		switch httpHandeler, ok := bu[subdomain]; ok {
		case true:
			httpHandeler(w, r)

		default:
			defaultService(w, r)
		}
	}
}

func (s Server) Run(addr string) error {
	s.setUpMainRout()
	s.log.Info("Listen And Serve: ", addr)
	return http.ListenAndServe(addr, s.http)
}
