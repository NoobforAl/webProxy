package interactor

import (
	"context"
	"io"
	"net/http"
)

func (wb *WebPage) SetPath(path string) *WebPage {
	wb.url = wb.url + path
	wb.log.Debug("fix url fix http: ", wb.url)
	return wb
}

func (wb *WebPage) NewRequest(ctx context.Context, method string, bodyReq io.Reader) (*WebPage, error) {
	wb.log.Info("create New reaquest | with method: ", method)
	req, err := http.NewRequestWithContext(ctx, method, wb.url, bodyReq)
	if err != nil {
		return wb, err
	}

	wb.log.Info("create New reaquest | Setup header")
	for k, v := range wb.header {
		req.Header[k] = v
	}

	req.Header.Del("Referer")
	req.Header.Set("Origin", "https://ieeexplore.ieee.org")
	req.Header.Set("Access-Control-Allow-Origin", "*")

	wb.req = req
	return wb, err
}

func (wb *WebPage) DoReq() (*http.Response, error) {
	wb.log.Info("Do request")
	res, err := wb.client.Do(wb.req)
	return res, err
}
