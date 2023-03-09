package response

type MyCode int64

const (
	CodeDbFailed        MyCode = 5555
	CodeSuccess         MyCode = 2333
	CodeInvalidParams   MyCode = 1001
	CodeUserExist       MyCode = 1002
	CodeUserNotExist    MyCode = 1003
	CodeInvalidPassword MyCode = 1004
	CodeServerBusy      MyCode = 1005

	CodeInvalidToken      MyCode = 1006
	CodeInvalidAuthFormat MyCode = 1007
	CodeNotLogin          MyCode = 1008
	CodeUpdateFailed      MyCode = 1009
	CodeCaptchaFailed     MyCode = 1010
	CodePolicyFailed      MyCode = 1011
	CodeInvalidPhoto      MyCode = 1012
	CodeInvalidPhone      MyCode = 1013
	CodeServerError       MyCode = 1014
	CodeInvalidRandCode   MyCode = 1015
)

var msgFlags = map[MyCode]string{
	CodeDbFailed:        "数据库错误",
	CodeSuccess:         "success",
	CodeInvalidParams:   "请求参数错误",
	CodeUserExist:       "已经报名过",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",

	CodeInvalidToken:      "无效的Token",
	CodeInvalidAuthFormat: "认证格式有误",
	CodeNotLogin:          "未登录",
	CodeUpdateFailed:      "修改用户信息失败",
	CodeCaptchaFailed:     "验证码错误",
	CodePolicyFailed:      "权限不足",
	CodeInvalidPhoto:      "照片文件上传错误",
	CodeInvalidPhone:      "无效的手机号码",
	CodeServerError:       "服务器内部错误",
	CodeInvalidRandCode:   "手机验证码不匹配",
}

func (c MyCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
