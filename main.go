package main

import (
	"web_proxy/config"
	"web_proxy/core/logger"
	server "web_proxy/http"
)

func main() {
	conf := config.Arg
	log := logger.New(conf.Debug, conf.LogFile)
	serv := server.New(log, conf)
	log.Fatal(serv.Run(config.Arg.Listen))
}
