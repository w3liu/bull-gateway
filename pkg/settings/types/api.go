package types

import (
	"time"
)

type Api struct {
	Id          int64         `xorm:"BIGINT(20) not null pk autoincr"`
	GroupId     int64         `xorm:"BIGINT(20) not null index comment('分组id')"`
	Name        string        `xorm:"VARCHAR(100) not null comment('名称')"`
	Description string        `xorm:"VARCHAR(255) comment('描述')"`
	Status      ApiStatusType `xorm:"TINYINT(2) default 1 comment('上线状态：1 下线，2 上线')"`

	// 定义API请求
	ReqPath          string              `xorm:"VARCHAR(255) comment('请求Path')"`
	ReqHttpMethod    HTTPMethod          `xorm:"VARCHAR(50) not null comment('请求方法：GET，POST，PATCH，PUT，DELETE，HEAD，OPTIONS，ANY')"`
	InputRequestMode ApiInputRequestMode `xorm:"VARCHAR(100) not null comment('入参请求模式：MAPPING 入参映射（过滤未知参数），MAPPING_PASS_THROUGH 入参映射（透传未知参数），PASS_THROUGH 入参透传')"`

	// 定义API后端服务
	BackendType                BackendType                `xorm:"VARCHAR(100) not null comment('HTTP，RPC，MOCK')"`
	ServiceAddress             string                     `xorm:"VARCHAR(255) comment('后端服务地址')"` // 格式：http(s)://host:port
	ServicePath                string                     `xorm:"VARCHAR(255) comment('后端请求Path')"`
	ServiceHttpMethod          HTTPMethod                 `xorm:"VARCHAR(50) comment('请求方法：GET，POST，PATCH，PUT，DELETE，HEAD，OPTIONS，ANY')"`
	ServiceTimeout             int32                      `xorm:"INT(11) default 0 comment('后端超时，单位ms')"`
	ServiceContentTypeCategory ServiceContentTypeCategory `xorm:"VARCHAR(255) comment('后端content-type类型：CLIENT 透传客户端ContentType头，CUSTOM 自定义，DEFAULT API网关默认')"`
	ServiceContentTypeValue    string                     `xorm:"VARCHAR(255) comment('后端content-type'')"`     // 后端content-type，API网关默认application/x-www-form-urlencoded; charset=UTF-8
	ServiceConstParams         string                     `xorm:"VARCHAR(2000) comment('后端服务常量参数，json数组存储')"`  // ApiServiceParamConstParam
	ServiceParametersMap       string                     `xorm:"VARCHAR(2000) comment('入参和后端参数映射，json数组存储')"` // ApiServiceParameterMap 数组

	// mock方式配置
	MockResult         string `xorm:"TEXT comment('Mock返回结果')"`
	MockHTTPStatusCode int32  `xorm:"INT(11) default 0 comment('Mock HTTP Status Code')"`

	// 返回结果定义
	ResultType       ContentType `xorm:"VARCHAR(100) not null comment('返回ContentType：JSON，TEXT，BINARY，XML，HTML，PASSTHROUGH（透传后端Content-Type）')"`
	ResultSample     string      `xorm:"TEXT comment('返回结果示例')"`
	FailResultSample string      `xorm:"TEXT comment('失败返回结果示例')"`

	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

// API入参定义
type ApiRequestParam struct {
	ApiId             int64             `xorm:"BIGINT(20) not null index comment('api id')"`
	Order             int32             `xorm:"INT(11) default 0 comment('顺序')"`
	Name              string            `xorm:"VARCHAR(100) not null comment('参数名')"`
	ParameterLocation ParameterLocation `xorm:"TINYINT(2) not null comment('参数位置：1 Parameter Path，2 Head，3 Query，4 Body')"`
	ParameterType     ParameterType     `xorm:"TINYINT(2) not null comment('类型：1 String，2 Int，3 Long，4 Float，5 Double，6 Boolean 7 Array')"`
	IsRequired        ParamIsRequired   `xorm:"TINYINT(2) default 1 comment('是否必填：1 是，2 否')"`
	DefaultValue      string            `xorm:"VARCHAR(255) comment('默认值')"`
	DemoValue         string            `xorm:"VARCHAR(255) comment('示例')"`
	Description       string            `xorm:"VARCHAR(255) comment('描述')"`
	DocShow           DocShow           `xorm:"TINYINT(2) default 1 comment('文档可见：1 不可见，2 可见')"`
	DocOrder          int32             `xorm:"INT(11) default 0 comment('文档中顺序')"`
	Extra             string            `xorm:"VARCHAR(2000) comment('附加设置，json格式存储')"` // 下面的附加设置结构体序列化json后存储
	CreatedAt         time.Time         `xorm:"created"`
	UpdatedAt         time.Time         `xorm:"updated"`
}

// 入参类型String附加设置
type ApiRequestParamString struct {
	MaxLength         int32  `json:"maxLength"`         // 最大长度
	EnumValue         string `json:"enumValue"`         // 枚举，用逗号(,)隔开
	RegularExpression string `json:"regularExpression"` // 参数验证，正则表达式
}

// 入参类型Int、Long、Float、Double附加设置
type ApiRequestParamIntLongFloatDouble struct {
	EnumValue string `json:"enumValue"` // 枚举，用逗号(,)隔开
	MinValue  int64  `json:"minValue"`  // 最小值
	MaxValue  int64  `json:"maxValue"`  // 最大值
}

// 入参类型Array附加设置
type ApiRequestParamArray struct {
	ArrayItemsType    ParamArrayItemsType `json:"arrayItemsType"`    // 数组字段类型：String，Int，Long，Float，Double，Boolean，File
	MaxLength         int32               `json:"maxLength"`         // 最大长度
	EnumValue         string              `json:"enumValue"`         // 枚举，用逗号(,)隔开
	RegularExpression string              `json:"regularExpression"` // 参数验证，正则表达式
}

// 后端服务参数配置
type ApiServiceParam struct {
	ApiId             int64             `xorm:"BIGINT(20) not null index comment('api id')"`
	Order             int32             `xorm:"INT(11) default 0 comment('顺序')"`
	Name              string            `xorm:"VARCHAR(100) not null comment('参数名')"`
	ParameterLocation ParameterLocation `xorm:"TINYINT(2) not null comment('参数位置：1 Parameter Path，2 Head，3 Query，4 Body')"`
}

// 后端服务常量参数
type ApiServiceParamConstParam struct {
	Name          string            `json:"name"`          // 参数名称
	Value         string            `json:"value"`         // 参数值
	Location      ParameterLocation `json:"location"`      // 参数位置，2 Head，3 Query
	ParameterType ParameterType     `json:"parameterType"` // 参数类型
	Description   string            `json:"description"`   // 描述
}

// 请求入参和后端参数映射
type ApiServiceParameterMap struct {
	RequestParameterName string `json:"requestParameterName"` // 请求入参名称
	ServiceParameterName string `json:"serviceParameterName"` // 后端服务参数名称

}
