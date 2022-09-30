package user

import (
	"Skywing/controller"
	"Skywing/models"
	res "Skywing/models/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *UserController) Update(c *gin.Context) {
	// 获取请求参数
	var newInfo models.RegisterForm
	// 参数解析
	if err := c.ShouldBindJSON(&newInfo); err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		res.ResponseErrorWithMsg(c, res.CodeInvalidParams, err.Error())
		return
	}
	// 参数校验
	if err := controller.Val.Struct(newInfo); err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		res.ResponseErrorWithMsg(c, res.CodeInvalidParams, err.Error())
		return
	}
	err := u.Srv.Users().Update(&newInfo)
	if err != nil {
		zap.L().Error("修改用户信息失败", zap.Error(err))
		res.ResponseError(c, res.CodeUpdateFailed)
		return
	}
	res.ResponseSuccess(c, nil)
}
