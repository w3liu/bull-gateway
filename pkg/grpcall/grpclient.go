package grpcall

import "context"

type GrpClient interface {
	// 初始化
	Init(...Option) error
	// 返回当前的 options
	Options() Options
	// 调用方法
	Call(serviceName, methodName, data string) (*Response, error)
	// 调用方法（自定义context）
	CallWithCtx(ctx context.Context, serviceName, methodName, data string) (*Response, error)
}

type Option func(*Options)
