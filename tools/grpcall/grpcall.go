package grpcall

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/w3liu/bull/client"
	"google.golang.org/grpc"
	"io"
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

func (g *grpcall) Call(server, svcName, methodName, data string) (*Response, error) {
	var module = g.getModule(svcName)
	if module == "" {
		return nil, fmt.Errorf("svcName %q format invalid", svcName)
	}
	svc, _ := g.findService(module, svcName)
	if svc == nil {
		return nil, fmt.Errorf("svcName %q not found", svcName)
	}
	var mtd = svc.FindMethodByName(methodName)
	if mtd == nil {
		return nil, fmt.Errorf("service %q does not include a method named %q", svc, methodName)
	}

	var cli = g.getClient(server)

	if cli == nil {
		return nil, fmt.Errorf("server %q does not include a client", server)
	}

	msgFactory := dynamic.NewMessageFactoryWithExtensionRegistry(dynamic.NewExtensionRegistryWithDefaults())
	req := msgFactory.NewMessage(mtd.GetInputType())
	resp := msgFactory.NewMessage(mtd.GetOutputType())
	var inData io.Reader
	inData = strings.NewReader(data)

	rf, err := g.requestParserFor(inData)
	if err != nil {
		return nil, err
	}
	err = rf.Next(req)
	if err != nil {
		return nil, err
	}
	err = cli.Invoke(context.TODO(), fmt.Sprintf("/%s/%s", mtd.GetService().GetFullyQualifiedName(), mtd.GetName()), req, resp)

	formatter, err := g.newFormatter(true)
	if err != nil {
		return nil, err
	}

	respText, err := formatter(resp)
	if err != nil {
		return nil, err
	}

	var result = &Response{
		IsStream: false,
		Data:     respText,
	}

	return result, err
}

func (g *grpcall) CallWithCtx(ctx context.Context, server, svcName, methodName, data string) (*Response, error) {
	return nil, nil
}

// 解析proto
func (g *grpcall) parseProto(module string) (*desc.FileDescriptor, error) {
	var content = g.opts.Resource.GetProtoFileContent(module)
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
	if _, ok := g.opts.FdMap[module]; !ok {
		fd, err := g.parseProto(module)
		if err != nil {
			return nil, err
		}
		g.RLock()
		g.opts.FdMap[module] = fd
		g.RUnlock()
	}
	var svc = g.opts.FdMap[module].FindService(svcName)
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

// 获取client
func (g *grpcall) getClient(server string) *grpc.ClientConn {
	g.RLock()
	_, ok := g.opts.ClientMap[server]
	g.RUnlock()
	if !ok {
		cli := client.NewClient(
			client.Registry(g.opts.Registry),
			client.Service(server))
		g.Lock()
		g.opts.ClientMap[server] = cli.Instance().(*grpc.ClientConn)
		g.Unlock()
	}
	return g.opts.ClientMap[server]
}

func (g *grpcall) anyResolver() (jsonpb.AnyResolver, error) {
	files := make([]*desc.FileDescriptor, 0)
	for _, fd := range g.opts.FdMap {
		files = append(files, fd)
	}

	var er dynamic.ExtensionRegistry
	for _, fd := range files {
		er.AddExtensionsFromFile(fd)
	}
	mf := dynamic.NewMessageFactoryWithExtensionRegistry(&er)
	return dynamic.AnyResolver(mf, files...), nil
}

func (g *grpcall) requestParserFor(in io.Reader) (RequestParser, error) {
	resolver, err := g.anyResolver()
	if err != nil {
		return nil, fmt.Errorf("error creating message resolver: %v", err)
	}

	return NewJSONRequestParser(in, resolver), nil
}

func (g *grpcall) newFormatter(emitFields bool) (Formatter, error) {
	resolver, err := g.anyResolver()
	if err != nil {
		return nil, fmt.Errorf("error creating message resolver: %v", err)
	}
	return NewJSONFormatter(emitFields, resolver), err
}
