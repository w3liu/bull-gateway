package grpcall

type ProtoResource interface {
	GetProtoFileContent(module string) string
}
