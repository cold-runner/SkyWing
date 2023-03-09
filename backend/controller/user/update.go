package user

import (
	"Skywing/controller"
	"Skywing/models"
	"Skywing/models/response"
	"Skywing/store/redis"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *UserController) Update(c *gin.Context) {
	// 图片验证码校验
	capId := c.PostForm("captchaID")
	capVal := c.PostForm("captchaCode")
	if capId == "" || capVal == "" {
		response.ResponseError(c, response.CodeCaptchaFailed)
		return
	}
	if success := redis.RdbClient.Verify(capId, capVal, true); !success {
		response.ResponseError(c, response.CodeCaptchaFailed)
		return
	}
	// 获取请求参数
	updateObj := models.UpdateForm{
		Introduce: c.PostForm("introduce"),
	}
	// 参数校验
	if err := controller.Val.Struct(updateObj); err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		response.ResponseErrorWithMsg(c, response.CodeInvalidParams, err.Error())
		return
	}
	// 校验照片
	_, _, err := c.Request.FormFile("photo")
	if err != nil {
		zap.L().Error("照片文件校验错误", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidPhoto)
		return
	}
	// 绑定照片属性
	_, phoHead, _ := c.Request.FormFile("photo")
	updateObj.Photo = phoHead

	uuid := c.Param("uuid")
	err = u.Srv.Users().Update(uuid, &updateObj)
	if err != nil {
		zap.L().Error("修改用户信息失败", zap.Error(err))
		response.ResponseError(c, response.CodeUpdateFailed)
		return
	}
	response.ResponseSuccess(c, nil)
}
