package grpcall

import (
	"github.com/jhump/protoreflect/desc"
)

type Options struct {
	resource ProtoResource
	fdMap    map[string]*desc.FileDescriptor
}

func newOptions(opts ...Option) Options {
	var options = Options{}
	return options
}

func Resource(resource ProtoResource) Option {
	return func(options *Options) {
		options.resource = resource
	}
}

func FdMap(fdMap map[string]*desc.FileDescriptor) Option {
	return func(options *Options) {
		options.fdMap = fdMap
	}
}
