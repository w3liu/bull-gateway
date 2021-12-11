package model

import (
	"github.com/w3liu/bull-gateway/pkg/settings/types"
	"time"
)

// 插件
type Plugin struct {
	Id     int64            `xorm:"BIGINT(20) not null pk autoincr"`
	Name   string           `xorm:"VARCHAR(100) not null comment('名称')"`
	Type   types.PluginType `xorm:"TINYINT(2) not null comment('插件类型：1 IP访问，2 签名，3 JWT，4 跨域访问，5 缓存，6 错误码映射，7 流量控制，8 OAuth2，9 熔断，10 crycx auth app，11 crycx auth h5，12 crycx real name，13 crycx view count，14 身份认证')"`
	Des    string           `xorm:"VARCHAR(255) comment('备注')"`
	Config string           `xorm:"TEXT comment('配置')"`
	Weight int32            `xorm:"INT default 0 comment('执行顺序权重')"`

	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

// 绑定api
type PluginBindApi struct {
	Id       int64 `xorm:"BIGINT(20) not null pk autoincr"`
	PluginId int64 `xorm:"BIGINT(20) not null comment('插件id')"`
	ApiId    int64 `xorm:"BIGINT(20) not null comment('api id')"`
}
