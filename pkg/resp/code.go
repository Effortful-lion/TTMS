package resp



type ResCode int64

// 定义响应码 和 响应信息的映射

var codeMsgMap = map[ResCode]string {
	CodeSuccess: "success",
	CodeInvalidParams: "请求参数错误",
	CodeError: "服务器内部错误",
	CodeNeedLogin: "需要登录",
	CodeUserExist: "用户已存在",
	CodeUserNotExist: "用户不存在",
	CodeInvalidPassword: "密码错误",
	CodeNoAuth: "没有权限",
}

// 系统响应状态码
const (
	CodeSuccess ResCode = 1000 + iota // 成功
	CodeInvalidParams                 // 参数错误
	CodeError                         // 服务器内部错误
	CodeNeedLogin                     // 需要登录
	CodeNoAuth                        // 没有权限
)

// 用户业务状态码
const (
	CodeUserExist       ResCode = 2000 + iota // 用户已存在
	CodeUserNotExist                          // 用户不存在
	CodeInvalidPassword                        // 密码错误
)

func (c ResCode) Msg() string { 
	msg, ok := codeMsgMap[c] 
	if !ok {                 
		msg = codeMsgMap[CodeError]
	}
	return msg
}