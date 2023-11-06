package builder

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"strings"
	"web_proxy/interactor"
)

func (hs httpSub) handelService() Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		path := r.URL.RequestURI()

		webProxy, err := interactor.New(hs.log, hs.service).
			MakeAddr(path).
			SetClient().
			SetRequest(ctx, r.Method, r.Body)

		if err != nil {
			hs.setError(w, err.Error())
			return
		}

		// this line is optional you can ignore it!
		webProxy = webProxy.SetHeader(r.Header).SetProxy()

		res, err := webProxy.Do()
		if err != nil {
			hs.setError(w, err.Error())
			return
		}

		defer res.Body.Close()
		webProxy.SetResponseHeader(w, res.Header)
		hs.makeResponseBody(w, res)
	}
}

func (hs httpSub) makeResponseBody(
	w http.ResponseWriter,
	res *http.Response,
) {
	contentType := res.Header.Get("Content-Type")
	if strings.Count(contentType, "text/html") > 0 {
		hs.log.Debug("body is text/html")
		hs.responseHtmlFile(w, res)
	} else {
		hs.log.Debug("body is not text/html use io copy")
		hs.responseByteFile(w, res)
	}
}

func (hs httpSub) responseHtmlFile(w http.ResponseWriter, res *http.Response) {
	errMsgCopy := "Error copy content:"

	hs.log.Debug("read body content for response")
	b, err := io.ReadAll(res.Body)
	if err != nil {
		errMsgCopy += err.Error()
		hs.setError(w, errMsgCopy)
		return
	}

	hs.log.Debug("check gzip file for response")
	//b = hs.gzipByteDecode(b)
	//b = hs.gzipByteEncode(b)
	b = bytes.Replace(b, []byte(hs.service.ServiceDomain), []byte(""), -1)
	_, err = w.Write(b)
	if err != nil {
		errMsgCopy += err.Error()
		hs.setError(w, errMsgCopy)
	}
}

func (hs httpSub) responseByteFile(w http.ResponseWriter, res *http.Response) {
	_, err := io.Copy(w, res.Body)
	if err != nil {
		errMsg := "Error copy conetent:" + err.Error()
		hs.setError(w, errMsg)
		return
	}
}

func (hs httpSub) gzipByteDecode(b []byte) []byte {
	buf := bytes.NewBuffer(b)
	gz, err := gzip.NewReader(buf)
	if err != nil {
		return b
	}

	hs.log.Debug("is a gzip file, extract")
	defer gz.Close()

	var bufRes bytes.Buffer
	_, _ = io.Copy(&bufRes, gz)
	return bufRes.Bytes()
}

func (hs httpSub) gzipByteEncode(b []byte) []byte {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)

	_, err := gz.Write(b)
	if err != nil {
		return b
	}

	hs.log.Debug("gzip data")
	err = gz.Close()
	if err != nil {
		return b
	}

	return buf.Bytes()
}

// set error message in response
func (hs httpSub) setError(w http.ResponseWriter, errMsg string) {
	http.Error(w, errMsg, http.StatusInternalServerError)
	hs.log.Error(errMsg)
}
