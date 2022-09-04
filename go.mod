module github.com/w3liu/bull-gateway

go 1.14

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.8
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/coreos/bbolt v0.0.0-00010101000000-000000000000 // indirect
	github.com/gin-gonic/gin v1.7.7
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/jhump/protoreflect v1.6.0
	github.com/novalagung/gubrak v1.0.0
	github.com/w3liu/bull v0.1.0
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/genproto v0.0.0-20210406143921-e86de6bf7a46 // indirect
	google.golang.org/grpc v1.37.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	xorm.io/xorm v1.0.7
)
