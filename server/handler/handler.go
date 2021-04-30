package handler

import (
	"context"
	"github.com/w3liu/bull-gateway/pkg/log"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"
)

type Handler struct {
	tr *http.Transport
}

func New() *Handler {
	return &Handler{
		tr: &http.Transport{},
	}
}

func (s *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "HEAD" && r.URL.Path == "/" {
		return
	}
	//_, _ = w.Write([]byte("hello world"))

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://www.baidu.com", r.Body)

	r.URL.Path = "http://www.baidu.com"

	resp, err := s.tr.RoundTrip(req)
	if err != nil {
		log.Error("err", zap.Error(err))
		return
	}
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Error("err", zap.Error(err))
		return
	}
}
