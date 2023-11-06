package http

import (
	"bytes"
	"html/template"
	"net/http"
	"net/url"
	"strings"
	"web_proxy/http/builder"
)

// body response to main page
var bodyMainPage []byte

func (s Server) makeBodyMainPage() {
	var temp = `
		<!DOCTYPE html>
		<html>
			<head>
				<meta charset="utf-8">
				<meta http-equiv="X-UA-Compatible" content="IE=edge">
				<title>Web Proxy</title>
				<meta name="description" content="empty">
				<meta name="viewport" content="width=device-width, initial-scale=1">
				<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN" crossorigin="anonymous">
			</head>
			<body class="bg-dark">
				<div class="d-flex flex-column flex-md-row p-4 gap-4 md-4 align-items-center justify-content-center">
					<ul class="dropdown-menu position-static d-grid gap-1 p-2 rounded-3 mx-0 shadow w-75" data-bs-theme="dark">
						<li><h3 class="dropdown-item rounded-2">Urls</h3></li>
						<li><hr class="dropdown-divider"></li>
						{{range $name, $val := .Service}}
							<li><a class="dropdown-item rounded-2" href="{{ url $val.ServiceUrl }}">{{$name}}</a></li>
						{{end}}
					</ul>
				</div>
			</body>
		</html>
	`
	templateFunc := template.FuncMap{
		"url": func(u string) string {
			uri, err := url.Parse(s.config.Listen)
			if err != nil {
				s.log.Panic(err)
			}

			port := uri.Port()
			if port == "" {
				pat := strings.Split(s.config.Listen, ":")
				switch len(pat) {
				case 1:
					port = "80"
				case 2:
					port = pat[1]
				case 3:
					port = pat[2]
				}
			}

			return "http://" + u + ":" + port
		},
	}

	html, err := template.New("template").Funcs(templateFunc).Parse(temp)
	if err != nil {
		s.log.Panic(err)
	}

	var buf bytes.Buffer
	err = html.Execute(&buf, s.config)
	if err != nil {
		s.log.Panic(err)
	}

	bodyMainPage = buf.Bytes()
}

// setup default template service
//
// in main page show you urls
func (s Server) defaultService() builder.Handler {
	s.makeBodyMainPage()
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		_, _ = w.Write(bodyMainPage)
	}
}
