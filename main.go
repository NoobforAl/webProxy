package main

import (
	"web_proxy/config"
	server "web_proxy/http"
	"web_proxy/logger"
)

func main() {
	conf := config.Arg
	log := logger.New(conf.Debug, conf.LogFile)

	serve := server.New(log, conf)
	log.Fatal(serve.Run(config.Arg.Listen))
}
