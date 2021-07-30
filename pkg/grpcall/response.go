package grpcall

import (
	"context"
	"google.golang.org/grpc/metadata"
)

type Response struct {
	ResultChan chan string
	SendChan   chan []byte
	DoneChan   chan error
	Data       string
	RespHeader metadata.MD
	IsStream   bool
	Cancel     context.CancelFunc
}
