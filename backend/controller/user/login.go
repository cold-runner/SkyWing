package user

import (
	"Skywing/controller"
	"Skywing/models"
	"Skywing/models/response"
	"Skywing/store/redis"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *UserController) Login(c *gin.Context) {
	var fo models.LoginUser
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
	// 图片验证码校验
	if success := redis.RdbClient.Verify(fo.Captcha.CaptchaId, fo.Captcha.CaptchaVal, true); !success {
		response.ResponseError(c, response.CodeCaptchaFailed)
		return
	}
	// 登录逻辑
	userInfo, err := u.Srv.Users().Login(c, &fo)
	if err != nil {
		zap.L().Error("登录失败", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidPassword)
		return
	}

	response.ResponseSuccess(c, userInfo)
}

func (u *UserController) Logout(c *gin.Context) {
	// 登出逻辑直接让前端把token删除了算了233333
}
