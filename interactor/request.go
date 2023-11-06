package interactor

import (
	"context"
	"io"
	"net/http"
)

// fix request path
//
// set main url before path
func (wp *WebProxy) MakeAddr(path string) *WebProxy {
	wp.reqUrl = wp.mainUrl + path
	wp.log.Debug("setup path to main url: ", wp.reqUrl)
	return wp
}

// setup client for new request
func (wp *WebProxy) SetClient() *WebProxy {
	client := &http.Client{}
	wp.client = client
	return wp
}

// create new request content with body, method
func (wb *WebProxy) SetRequest(
	ctx context.Context,
	method string,
	bodyReq io.Reader,
) (*WebProxy, error) {
	wb.log.Infof("create new request | %s | %s ", method, wb.reqUrl)
	req, err := http.NewRequestWithContext(ctx, method, wb.reqUrl, bodyReq)
	if err != nil {
		return wb, err
	}

	wb.req = req
	return wb, err
}

// send new request with client and request body
func (wb *WebProxy) Do() (*http.Response, error) {
	wb.log.Debug("start send request")
	res, err := wb.client.Do(wb.req)
	wb.log.Debugf("finished request, error? :", err)
	return res, err
}
