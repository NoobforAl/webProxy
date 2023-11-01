package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Auth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Proxy struct {
	Addr string `yaml:"addr"`
	Auth Auth   `yaml:"auth"`
}

type Service struct {
	Domain string  `yaml:"domain"`
	Url    string  `yaml:"url"`
	Proxy  []Proxy `yaml:"proxy"`
}

type ConfigWebProxy struct {
	Debug   bool               `yaml:"debug"`
	LogFile string             `yaml:"logfile"`
	Listen  string             `yaml:"listen"`
	Service map[string]Service `yaml:"service"`
}

var Arg ConfigWebProxy

func init() {
	b, err := os.ReadFile("webProxConf.yaml")
	if err != nil {
		errMsg := "Error load config file: " + err.Error()
		panic(errMsg)
	}

	err = yaml.Unmarshal(b, &Arg)
	if err != nil {
		errMsg := "Error load config file: " + err.Error()
		panic(errMsg)
	}
}
