# webProxy

See content website on localhost.

### How run it ?
first setup config file like this:  
Config file example(file name: wpconf.yaml):
```yaml
listen: "localhost:2020"
logfile: ""
debug: false

service:
  example1:
    servicedomain: "http://example.com"
    serviceurl: "example.localhost"
    proxylist:
      - addr: "localhost:2020"
        username: "user"
        password: "passs"
      - addr: "localhost:2020"
        username: "user"
        password: "passs"

  example2:
    servicedomain: "http://example.com"
    serviceurl: "example.localhost"
    proxylist:
      - addr: "localhost:2020"
        username: "user"
        password: "passs"
      - addr: "localhost:2020"
        username: "user"
        password: "passs"

```

And run code:
> go run main.go

or download executable file in [release](https://github.com/NoobforAl/webProxy/releases/tag/v0.1.0) page.

#### How build ?
> go build -o prWeb.exe main.go

or build release (you need istall cmake):
> make release 
