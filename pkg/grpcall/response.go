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

func (r *Response) Read() {
}

func (r *Response) Write() {
}

func (r *Response) IsError() {
}

func (r *Response) IsClose() {
}

func (r *Response) Close() {
}
