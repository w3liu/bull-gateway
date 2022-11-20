package types

import "errors"

// 请求方法，请求和后端请求通用
type HTTPMethod string

const (
	HTTPMethodGet     HTTPMethod = "GET"
	HTTPMethodPost    HTTPMethod = "POST"
	HTTPMethodPatch   HTTPMethod = "PATCH"
	HTTPMethodPut     HTTPMethod = "PUT"
	HTTPMethodDelete  HTTPMethod = "DELETE"
	HTTPMethodHead    HTTPMethod = "HEAD"
	HTTPMethodOptions HTTPMethod = "OPTIONS"
	HTTPMethodAny     HTTPMethod = "ANY"
)

// 返回ContentType
type ContentType string

const (
	ContentTypeJson   ContentType = "application/json;charset=utf-8"
	ContentTypeText   ContentType = "text/plain;charset=utf-8"
	ContentTypeBinary ContentType = "application/octet-stream;charset=utf-8"
	ContentTypeXml    ContentType = "application/xml;charset=utf-8"
	ContentTypeHtml   ContentType = "text/html;charset=utf-8"
)

// api上线状态
type ApiStatusType int32

const (
	ApiStatusOffline ApiStatusType = iota + 1 // 下线
	ApiStatusOnline                           // 上线
)

func (t ApiStatusType) IsValid() error {
	switch t {
	case ApiStatusOffline, ApiStatusOnline:
		return nil
	}
	return errors.New("invalid type")
}

// api入参请求模式
type ApiInputRequestMode string

const (
	ApiInputRequestModeMapping            ApiInputRequestMode = "MAPPING"              // 入参映射（过滤未知参数）
	ApiInputRequestModeMappingPassThrough ApiInputRequestMode = "MAPPING_PASS_THROUGH" // 入参映射（透传未知参数）
	ApiInputRequestModePassThrough        ApiInputRequestMode = "PASS_THROUGH"         // 入参透传
)

// 后端服务类型
type BackendType string

const (
	BackendTypeHttp BackendType = "HTTP"
	BackendTypeRpc  BackendType = "RPC"
	BackendTypeMock BackendType = "MOCK"
)

// 后端content-type类型
type ServiceContentTypeCategory string

const (
	ServiceContentTypeCatagoryClient  ServiceContentTypeCategory = "CLIENT"  // 透传客户端ContentType头，CUSTOM 自定义，DEFAULT API网关默认
	ServiceContentTypeCatagoryCustom  ServiceContentTypeCategory = "CUSTOM"  // 自定义，DEFAULT API网关默认
	ServiceContentTypeCatagoryDefault ServiceContentTypeCategory = "DEFAULT" // API网关默认

	ServiceContentTypeValueDefault = "application/x-www-form-urlencoded; charset=UTF-8" // 后端content-type，API网关默认
)

// 参数位置
type ParameterLocation int32

const (
	ParameterLocationPath  ParameterLocation = iota + 1 // Parameter Path
	ParameterLocationHead                               // Head
	ParameterLocationQuery                              // Query
	ParameterLocationBody                               // Body
)

func (t ParameterLocation) String() string {
	switch t {
	case ParameterLocationPath:
		return "Parameter Path"
	case ParameterLocationHead:
		return "Head"
	case ParameterLocationQuery:
		return "Query"
	case ParameterLocationBody:
		return "Body"
	}
	return "unknown"
}

func (t ParameterLocation) IsValid() error {
	switch t {
	case
		ParameterLocationPath,
		ParameterLocationHead,
		ParameterLocationQuery,
		ParameterLocationBody:
		return nil
	}
	return errors.New("invalid type")
}

// 参数数据类型
type ParamDataType string

const (
	ParamDataTypeQuery ParamDataType = "query"
	ParamDataTypeHead  ParamDataType = "head"
	ParamDataTypeForm  ParamDataType = "form"
	ParamDataTypeJson  ParamDataType = "json"
)

func (t ParamDataType) String() string {
	return string(t)
}

// 入参数据类型
type ParameterType int32

const (
	ParameterTypeString  ParameterType = iota + 1 // String
	ParameterTypeInt                              // Int
	ParameterTypeLong                             // Long
	ParameterTypeFloat                            // Float
	ParameterTypeDouble                           // Double
	ParameterTypeBoolean                          // Boolean
	ParameterTypeArray                            // Array
)

func (t ParameterType) String() string {
	switch t {
	case ParameterTypeString:
		return "String"
	case ParameterTypeInt:
		return "Int"
	case ParameterTypeLong:
		return "Long"
	case ParameterTypeFloat:
		return "Float"
	case ParameterTypeDouble:
		return "Double"
	case ParameterTypeBoolean:
		return "Boolean"
	case ParameterTypeArray:
		return "Array"
	}
	return "unknown"
}

func (t ParameterType) IsValid() error {
	switch t {
	case
		ParameterTypeString,
		ParameterTypeInt,
		ParameterTypeLong,
		ParameterTypeFloat,
		ParameterTypeDouble,
		ParameterTypeBoolean,
		ParameterTypeArray:
		return nil
	}
	return errors.New("invalid type")
}

// 入参是否必填
type ParamIsRequired int32

const (
	ParamIsRequiredYes ParamIsRequired = iota + 1 // 是
	ParamIsRequiredNo                             // 否
)

func (t ParamIsRequired) String() string {
	switch t {
	case ParamIsRequiredYes:
		return "是"
	case ParamIsRequiredNo:
		return "否"
	}
	return "unknown"
}

func (t ParamIsRequired) IsValid() error {
	switch t {
	case ParamIsRequiredYes, ParamIsRequiredNo:
		return nil
	}
	return errors.New("invalid type")
}

// 文档可见
type DocShow int32

const (
	DocShowNo  DocShow = iota + 1 // 不可见
	DocShowYes                    // 可见
)

// 入参数组字段类型
type ParamArrayItemsType string

const (
	ParamArrayItemsTypeString  ParamArrayItemsType = "String"  // String
	ParamArrayItemsTypeInt     ParamArrayItemsType = "Int"     // Int
	ParamArrayItemsTypeLong    ParamArrayItemsType = "Long"    // Long
	ParamArrayItemsTypeFloat   ParamArrayItemsType = "Float"   // Float
	ParamArrayItemsTypeDouble  ParamArrayItemsType = "Double"  // Double
	ParamArrayItemsTypeBoolean ParamArrayItemsType = "Boolean" // Boolean
	ParamArrayItemsTypeFile    ParamArrayItemsType = "File"    // File
)

// 插件类型
type PluginType int32

const (
	PluginTypeIp             PluginType = iota + 1 // IP访问
	PluginTypeSign                                 // 签名
	PluginTypeJwt                                  // JWT
	PluginTypeCors                                 // 跨域访问
	PluginTypeCache                                // 缓存
	PluginTypeCodeMapping                          // 错误码映射
	PluginTypeTrafficControl                       // 流量控制
	PluginTypeOAuth2                               // OAuth2
	PluginTypeBreaker                              // 熔断

)

func (t PluginType) IsValid() error {
	switch t {
	case
		PluginTypeIp,
		PluginTypeSign,
		PluginTypeJwt,
		PluginTypeCors,
		PluginTypeCache,
		PluginTypeCodeMapping,
		PluginTypeTrafficControl,
		PluginTypeOAuth2,
		PluginTypeBreaker:
		return nil
	}
	return errors.New("invalid type")
}

// 执行权重，值越大越先执行
func (t PluginType) Weight() int32 {
	switch t {
	case PluginTypeIp:
		return 10000
	case PluginTypeTrafficControl:
		return 9000
	case PluginTypeBreaker:
		return 8000
	case PluginTypeSign, PluginTypeJwt:
		return 7000
	case PluginTypeCache:
		return 6000
	}
	return 0
}
