package grpcall

type Resource interface {
	GetProtoFileContent(module string) string
}
