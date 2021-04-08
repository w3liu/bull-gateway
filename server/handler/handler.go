package handler

import "net/http"

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (s *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "HEAD" && r.URL.Path == "/" {
		return
	}
	_, _ = w.Write([]byte("hello world"))
}
