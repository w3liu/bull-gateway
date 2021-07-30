package grpcall

import (
	"fmt"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"sync"
)

type protoSet struct {
	sync.RWMutex
	resource Resource
	fdMap    map[string]*desc.FileDescriptor
}

func newProtoSet(resource Resource) *protoSet {
	var p = &protoSet{
		resource: resource,
		fdMap:    make(map[string]*desc.FileDescriptor),
	}
	return p
}

func (p *protoSet) parseProto(module string) (*desc.FileDescriptor, error) {
	var content = p.resource.GetProtoFileContent(module)
	var parser = protoparse.Parser{
		Accessor: protoparse.FileContentsFromMap(map[string]string{
			module: content,
		}),
	}
	fds, err := parser.ParseFiles(module)
	if err != nil {
		return nil, err
	}
	if fds == nil || len(fds) == 0 {
		return nil, fmt.Errorf("ParseFiles module %v Not found", module)
	}
	return fds[0], nil
}

func (p *protoSet) findService(module, svcName string) (*desc.ServiceDescriptor, error) {
	if _, ok := p.fdMap[module]; !ok {
		fd, err := p.parseProto(module)
		if err != nil {
			return nil, err
		}
		p.RLock()
		defer p.RUnlock()
		p.fdMap[module] = fd
	}
	var svc = p.fdMap[module].FindService(svcName)
	if svc == nil {
		return nil, fmt.Errorf("service %v not found", svcName)
	}
	return svc, nil
}
