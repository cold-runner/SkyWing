package user

import (
	"Skywing/controller"
	"Skywing/models"
	"Skywing/models/response"
	"Skywing/pkg/captcha"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *UserController) Login(c *gin.Context) {
	var fo models.LoginUser
	// 验证码校验
	if success := captcha.Store.Verify(fo.CaptchaId, fo.Captcha, true); !success {
		response.ResponseError(c, response.CodeCaptchaFailed)
		return
	}
	// 参数解析
	if err := c.ShouldBindJSON(&fo); err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		response.ResponseErrorWithMsg(c, response.CodeInvalidParams, err.Error())
		return
	}
	// 参数校验
	if err := controller.Val.Struct(fo); err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		response.ResponseErrorWithMsg(c, response.CodeInvalidParams, err.Error())
		return
	}
	// 登录逻辑
	userInfo, err := u.Srv.Users().Login(c, &fo)
	if err != nil {
		zap.L().Error("登录失败", zap.Error(err))
		response.ResponseError(c, response.CodeUserNotExist)
		return
	}
	response.ResponseSuccess(c, userInfo)
}

//func RefreshTokenHandler(c *gin.Context) {
//	rt := c.Query("refresh_token")
//	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
//	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
//	// 这里的具体实现方式要依据你的实际业务情况决定
//	authHeader := c.Request.Header.Get("Authorization")
//	if authHeader == "" {
//		ResponseErrorWithMsg(c, CodeInvalidToken, "请求头缺少Auth Token")
//		c.Abort()
//		return
//	}
//	// 按空格分割
//	parts := strings.SplitN(authHeader, " ", 2)
//	if !(len(parts) == 2 && parts[0] == "Bearer") {
//		ResponseErrorWithMsg(c, CodeInvalidToken, "Token格式不对")
//		c.Abort()
//		return
//	}
//	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
//	fmt.Println(err)
//	c.JSON(http.StatusOK, gin.H{
//		"access_token":  aToken,
//		"refresh_token": rToken,
//	})
//}
