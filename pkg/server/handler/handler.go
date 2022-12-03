package handler

import (
	"bytes"
	"github.com/w3liu/bull-gateway/infra/grpcall"
	"github.com/w3liu/bull-gateway/infra/log"
	"github.com/w3liu/bull/registry"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type Handler struct {
	tr *http.Transport
}

func New() *Handler {
	return &Handler{
		tr: &http.Transport{},
	}
}

func getTestData() map[string]string {
	return map[string]string{
		"helloworld": `syntax = "proto3";

		package helloworld;
		
		// The greeting service definition.
		service Greeter {
			// Sends a greeting
			rpc SayHello (HelloRequest) returns (HelloReply) {}
			rpc SayGirl (HelloRequest) returns (HelloReply) {}
		}
		
		// The request message containing the user's name.
		message HelloRequest {
			string name = 1;
		}
		
		// The response message containing the greetings
		message HelloReply {
			string message = 1;
		}
		
		service Notify {
			// Sends notify
			rpc SayWorld (NotifyReq) returns (NotifyReply) {}
			rpc SayHaha (NotifyReq) returns (NotifyReply) {}
			rpc SayHi (NotifyReq) returns (NotifyReply) {}
		}
		
		// The request message containing the user's name.
		message NotifyReq{
			string nick = 1;
			string addr = 2;
		}
		
		// The response message containing the greetings
		message NotifyReply {
			string message = 1;
			string phone = 2;
			string car = 3;
		}
		
		service BidiStreamService {
			rpc BidiRPC (stream SimpleData) returns (stream SimpleData) {}
		}
		
		message SimpleData {
			string msg = 1;
		}
		
		service ServerStreamService {
			rpc StreamRpc(ServerStreamData) returns (stream ServerStreamData) {}
		}
		
		message ServerStreamData{
			string msg = 1;
		}`,
		"person": `syntax = "proto3";

		package person;
		
		service Person {
		  rpc SayHello(SayHelloRequest) returns (SayHelloResponse) {}
		}
		
		message SayHelloRequest {
		  string name = 1;
		}
		
		message SayHelloResponse {
		  string msg = 2;
		}`,
	}
}

type resource struct {
	content map[string]string
}

func (r *resource) GetProtoFileContent(module string) string {
	if _, ok := r.content[module]; ok {
		return r.content[module]
	} else {
		return ""
	}
}

func (s *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "HEAD" && r.URL.Path == "/" {
		return
	}
	var res = &resource{content: getTestData()}
	reg := registry.NewRegistry(registry.Addrs([]string{"127.0.0.1:2379"}...))

	cli := grpcall.NewClient(grpcall.Registry(reg), grpcall.Resource(res))
	resp, err := cli.Call("hello.svc", "person.Person", "SayHello", `{"name": "hello world"}`)
	if err != nil {
		log.Error("err", zap.Error(err))
		return
	}
	_, err = io.Copy(w, bytes.NewBufferString(resp.Data))
	if err != nil {
		log.Error("err", zap.Error(err))
		return
	}
}
