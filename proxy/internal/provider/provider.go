package provider

import (
	"io"
	"net/http"
)

type Handler struct {
}

func InitHandler() *Handler {
	handler := &Handler{}

	return handler
}

func (h *Handler) Proxy(w http.ResponseWriter, r *http.Request) {
	// удаляем заголовок Proxy-Connection
	r.Header.Del("Proxy-Connection")
	// заменяем путь на относительный
	r.RequestURI = ""

	// низкоуровнево отправляем запрос на r.Host
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	for name, values := range resp.Header {
		for _, v := range values {
			w.Header().Add(name, v)
		}
	}

	defer resp.Body.Close()
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
}
