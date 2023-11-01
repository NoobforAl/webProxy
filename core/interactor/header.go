package interactor

import (
	"net/http"

	"github.com/kr/pretty"
)

func (wb *WebPage) SetHeader(h http.Header) *WebPage {
	wb.log.Debugf("Set Request Header: %# v", pretty.Formatter(h))
	h.Del("Referer")
	h.Set("Origin", wb.urlDomain)
	h.Set("Access-Control-Allow-Origin", "*")

	wb.header = h
	return wb
}

func (wb WebPage) SetResponcHeader(w *http.ResponseWriter, header http.Header) {
	for k, v := range header {
		(*w).Header()[k] = v
	}

	(*w).Header().Del("Referer")
	(*w).Header().Set("Origin", wb.urlDomain)
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
