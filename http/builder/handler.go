package builder

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"web_proxy/config"
	"web_proxy/core/contract"
	"web_proxy/core/interactor"
)

type httpSub struct {
	name,
	domain string
	proxyList []config.Proxy
	log       contract.Logger
}

type Handelr func(http.ResponseWriter, *http.Request)

type BuilderHttp map[string]Handelr

func New(conf []*config.UrlRequester, log contract.Logger) BuilderHttp {
	builderHttp := make(BuilderHttp)
	for _, v := range conf {
		http := httpSub{
			name:      v.Name,
			domain:    v.Domain,
			proxyList: v.ProxyList,
			log:       log,
		}
		builderHttp[v.Name] = http.getWebPage()
	}
	return builderHttp
}

func (hs httpSub) getWebPage() Handelr {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.RequestURI()
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		ina := interactor.New(hs.log, hs.domain, hs.proxyList).
			SetHeader(r.Header).
			SetupProxy().
			SetPath(path)

		ina, err := ina.NewRequest(ctx, r.Method, r.Body)
		if err != nil {
			errMsg := err.Error()
			hs.setError(w, errMsg)
			return
		}

		res, err := ina.DoReq()
		if err != nil {
			errMsg := "request error: " + err.Error()
			hs.setError(w, errMsg)
			return
		}

		defer res.Body.Close()
		ina.SetResponcHeader(&w, res.Header)

		contentType := res.Header.Get("Content-Type")
		if strings.Count(contentType, "text/html") > 0 {
			b, err := io.ReadAll(res.Body)
			if err != nil {
				errMsg := "Error copy conetent:" + err.Error()
				hs.setError(w, errMsg)
				return
			}

			b = bytes.Replace(b, []byte(hs.domain), []byte(""), -1)
			_, err = w.Write(b)
			if err != nil {
				errMsg := "Error copy conetent:" + err.Error()
				hs.setError(w, errMsg)
				return
			}
		} else {
			_, err = io.Copy(w, res.Body)
			if err != nil {
				errMsg := "Error copy conetent:" + err.Error()
				hs.setError(w, errMsg)
				return
			}
		}

	}
}

func (hs httpSub) setError(w http.ResponseWriter, errMsg string) {
	http.Error(w, errMsg, http.StatusInternalServerError)
	hs.log.Error(errMsg)
}
