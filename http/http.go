package http

import (
	"net/http"
	"web_proxy/config"
	"web_proxy/contract"
)

type Server struct {
	http   *http.ServeMux
	config config.ConfigWebProxy
	log    contract.Logger
}

func New(log contract.Logger, conf config.ConfigWebProxy) *Server {
	return &Server{
		http:   http.NewServeMux(),
		config: conf,
		log:    log,
	}
}

func (s Server) Run(addr string) error {
	s.setUpRouter()
	s.log.Info("Listen And Serve: ", addr)
	return http.ListenAndServe(addr, s.http)
}
