package errcode

var (
	Success								= NewError(0, "成功")
	ServerError							= NewError(10000000, "服务内部错误")
	InvalidParams						= NewError(10000001, "参数错误")
	NotFound							= NewError(10000002, "无法找到")
	UnauthorizedAuthNotExist			= NewError(10000003, "鉴权失败，找不到对应的AppKey或AppSecret")
	UnauthorizedTokenError				= NewError(10000004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeOut			= NewError(10000005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate			= NewError(10000006, "鉴权失败，Token生成失败")
	TooManyRequests						= NewError(10000007, "请求次数过多")

)
