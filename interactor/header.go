package interactor

import (
	"net/http"
)

func (wp *WebProxy) SetHeader(h http.Header) *WebProxy {
	wp.log.Debugf("Set Request Header: %v", h)
	h.Del("Referer")
	h.Set("Origin", wp.mainUrl)
	h.Set("Access-Control-Allow-Origin", "*")

	wp.req.Header = h
	return wp
}

func (wb WebProxy) SetResponseHeader(w http.ResponseWriter, h http.Header) {
	h.Del("Referer")
	h.Set("Origin", wb.mainUrl)
	h.Set("Access-Control-Allow-Origin", "*")
	for k, v := range h {
		w.Header()[k] = v
	}
}
