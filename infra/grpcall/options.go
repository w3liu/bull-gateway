package grpcall

import (
	"time"

	"github.com/jhump/protoreflect/desc"
	"github.com/w3liu/bull/registry"
	"google.golang.org/grpc"
)

type Options struct {
	TimeOut   time.Duration
	Registry  registry.Registry
	Resource  ProtoResource
	ClientMap map[string]*grpc.ClientConn
	FdMap     map[string]*desc.FileDescriptor
}

func newOptions(opts ...Option) Options {
	var options = Options{
		TimeOut:   DefaultTimeout,
		ClientMap: make(map[string]*grpc.ClientConn),
		FdMap:     make(map[string]*desc.FileDescriptor),
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

func Registry(reg registry.Registry) Option {
	return func(options *Options) {
		options.Registry = reg
	}
}

func ClientMap(clientMap map[string]*grpc.ClientConn) Option {
	return func(options *Options) {
		options.ClientMap = clientMap
	}
}

func Resource(resource ProtoResource) Option {
	return func(options *Options) {
		options.Resource = resource
	}
}

func FdMap(fdMap map[string]*desc.FileDescriptor) Option {
	return func(options *Options) {
		options.FdMap = fdMap
	}
}
