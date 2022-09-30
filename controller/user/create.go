package user

import (
	"Skywing/controller"
	"Skywing/models"
	res "Skywing/models/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *UserController) Create(c *gin.Context) {
	// 获取请求参数
	var fo models.RegisterForm
	if err := c.ShouldBindJSON(&fo); err != nil {
		zap.L().Info("参数校验失败", zap.Error(err))
		res.ResponseErrorWithMsg(c, res.CodeInvalidParams, err.Error())
		return
	}
	// 校验数据有效性
	if err := controller.Val.Struct(fo); err != nil {
		zap.L().Info("参数校验失败", zap.Error(err))
		res.ResponseErrorWithMsg(c, res.CodeInvalidParams, err.Error())
		return
	}

	// 注册用户
	err := u.Srv.Users().Create(c, &fo)
	if err != nil {
		zap.L().Error("注册用户失败", zap.Error(err))
		res.ResponseError(c, res.CodeUserExist)
		return
	}
	res.ResponseSuccess(c, nil)
}
