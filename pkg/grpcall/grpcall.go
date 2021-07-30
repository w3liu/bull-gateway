package grpcall

import "context"

type GrpCall interface {
	Init(...Option) error
	Options() Options
	Call(serviceName, methodName, data string) (*Response, error)
	CallWithCtx(ctx context.Context, serviceName, methodName, data string) (*Response, error)
}

type Option func(*Options)
