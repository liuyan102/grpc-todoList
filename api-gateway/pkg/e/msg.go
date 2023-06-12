package e

var MsgFlags = map[uint]string{
	Success:                    "ok",
	Error:                      "fail",
	InvalidParams:              "请求的参数错误",
	ErrorAuthCheckTokenFail:    "token 错误",
	ErrorAuthCheckTokenTimeout: "token 验证超时",
}

// GetMsg 获取状态码对应信息
func GetMsg(code uint) string {
	if msg, ok := MsgFlags[code]; ok {
		return msg
	}
	return MsgFlags[Error]
}
