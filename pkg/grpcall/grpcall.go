package grpcall

import (
	"context"
	"fmt"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"strings"
	"sync"
)

type grpcall struct {
	sync.RWMutex
	opts Options
}

func newGrpCall(opts ...Option) GrpClient {
	var options = newOptions(opts...)
	return &grpcall{opts: options}
}

func (g *grpcall) Init(opts ...Option) error {
	for _, o := range opts {
		o(&g.opts)
	}
	return nil
}

func (g *grpcall) Options() Options {
	return g.opts
}

func (g *grpcall) Call(svcName, methodName, data string) (*Response, error) {
	var module = g.getModule(svcName)
	if module == "" {
		return nil, fmt.Errorf("svcName %q format invalid", svcName)
	}
	var svc = g.opts.fdMap[module].FindService(svcName)
	if svc == nil {
		return nil, fmt.Errorf("svcName %q not found", svcName)
	}
	var mtd = svc.FindMethodByName(methodName)
	if mtd == nil {
		return nil, fmt.Errorf("service %q does not include a method named %q", svc, methodName)
	}
	//msgFactory := dynamic.NewMessageFactoryWithExtensionRegistry(dynamic.NewExtensionRegistryWithDefaults())
	//var cc *grpc.ClientConn
	//req := msgFactory.NewMessage(mtd.GetInputType())
	//resp := msgFactory.NewMessage(mtd.GetOutputType())
	//var inData io.Reader
	//inData = strings.NewReader(data)
	//rf, err := s.requestParserFor(inData)
	//if err != nil {
	//	return nil, err
	//}
	//err = rf.Next(req)
	//if err != nil {
	//	return nil, err
	//}
	//err = cc.Invoke(context.TODO(), fmt.Sprintf("/%s/%s", mtd.GetService().GetFullyQualifiedName(), mtd.GetName()), req, resp)
	//return resp, err
	return nil, nil
}

func (g *grpcall) CallWithCtx(ctx context.Context, serviceName, methodName, data string) (*Response, error) {
	return nil, nil
}

// 解析proto
func (g *grpcall) parseProto(module string) (*desc.FileDescriptor, error) {
	var content = g.opts.resource.GetProtoFileContent(module)
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

// 查找Service
func (g *grpcall) findService(module, svcName string) (*desc.ServiceDescriptor, error) {
	if _, ok := g.opts.fdMap[module]; !ok {
		fd, err := g.parseProto(module)
		if err != nil {
			return nil, err
		}
		g.RLock()
		defer g.RUnlock()
		g.opts.fdMap[module] = fd
	}
	var svc = g.opts.fdMap[module].FindService(svcName)
	if svc == nil {
		return nil, fmt.Errorf("service %v not found", svcName)
	}
	return svc, nil
}

// 获取模块
func (g *grpcall) getModule(svcName string) string {
	var svcs = strings.Split(svcName, ".")
	if len(svcs) > 0 {
		return svcs[0]
	}
	return ""
}
