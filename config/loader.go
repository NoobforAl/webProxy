package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

// proxy configure
//
// ```yaml
// proxy:
//   - addr: "localhost:2020"
//     username: "user"
//     password: "pass"
//   - addr: "localhost:2020"
//     username: "user"
//     password: "pass"
//
// ```
type Proxy struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// service configure
//
// ```yaml
// servicedomain: "http://example.com"
// serviceurl: "example.localhost"
// proxy:
//   - addr: "localhost:2020"
//     username: "user"
//     password: "passs"
//
// ```
type Service struct {
	// service domain when you want request
	ServiceDomain string `yaml:"servicedomain"`

	// service url you want request in local address
	ServiceUrl string `yaml:"serviceurl"`

	// list of proxy
	ProxyList []Proxy `yaml:"proxylist"`
}

// base config app
//
// ```yaml
// listen: "localhost:2020"
// logfile: ""
// debug: false
//
// service:
//
//	   example:
//			servicedomain: "http://example.com"
//			serviceurl: "example.localhost"
//			proxylist:
//	 	 		- addr: "localhost:2020"
//	   			  username: "user"
//	    		  password: "passs"
//	 	 		- addr: "localhost:2020"
//	   			  username: "user"
//	    		  password: "passs"
//
// ```
type ConfigWebProxy struct {
	// run debug mode app
	Debug bool `yaml:"debug"`

	// path of log file
	LogFile string `yaml:"logfile"`

	// list ip:port
	Listen string `yaml:"listen"`

	// list of service
	Service map[string]Service `yaml:"service"`
}

var (
	// all config in yaml file load in this variable
	Arg ConfigWebProxy

	// program load config errors
	errLoadFile   = errors.New("Error load config file: ")
	errParsConfig = errors.New("Error pars config file: ")
)

// loaded config file when app start working
func init() {
	b, err := os.ReadFile("wpconf.yaml")
	if err != nil {
		panic(errors.Join(errLoadFile, err))
	}

	err = yaml.Unmarshal(b, &Arg)
	if err != nil {
		panic(errors.Join(errParsConfig, err))
	}
}
