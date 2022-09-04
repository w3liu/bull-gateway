package grpcall

import (
	"github.com/w3liu/bull/registry"
	"testing"
)

type resource struct {
	content map[string]string
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

func (r *resource) GetProtoFileContent(module string) string {
	if _, ok := r.content[module]; ok {
		return r.content[module]
	} else {
		return ""
	}
}

func TestGrpcall_Call(t *testing.T) {
	var res = &resource{content: getTestData()}
	r := registry.NewRegistry(registry.Addrs([]string{"127.0.0.1:2379"}...))
	var g = newGrpCall(Registry(r), Resource(res))
	resp, err := g.Call("hello.svc", "person.Person", "SayHello", `{"name": "hello world"}`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("resp", resp)
}
