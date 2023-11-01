# webProxy

See content website on localhost.

### How run it ?
first setup config file like this:  
Config file example(file name: webProxConf.yaml):
```yaml
listen: "localhost:2020"
logfile: "./log.log"
debug: false

service:
  Example1:
    domain: "https://example1.org"
    url: "subdomain.domain"
    
    proxy:
      - addr: "http://localhost:2022"
        auth: 
          username: "test"
          password: "test"

  Example2:
    domain: "https://example1.org"
    url: "subdomain.domain"

    proxy:
      - addr: "http://localhost:2021"
        auth: 
          username: "test"
          password: "test"
```

And run code:
> go run main.go

or download executable file in [release](https://github.com/NoobforAl/webProxy/releases/tag/v0.1.0) page.

#### How build ?
> go build -o prWeb.exe main.go

or build release (you need istall cmake):
> make release 
