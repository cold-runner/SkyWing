package user

import (
	"Skywing/models/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *UserController) GetInfo(c *gin.Context) {
	// 图片验证码校验
	//capId := c.PostForm("captchaID")
	//capVal := c.PostForm("captchaCode")
	//if capId == "" || capVal == "" {
	//	response.ResponseError(c, response.CodeCaptchaFailed)
	//	return
	//}
	//if success := redis.RdbClient.Verify(capId, capVal, true); !success {
	//	response.ResponseError(c, response.CodeCaptchaFailed)
	//	return
	//}
	// 查询信息
	get, err := u.Srv.Users().Get(c.Param("uuid"))
	if err != nil {
		zap.L().Error("查询信息失败", zap.Error(err))
		response.ResponseError(c, response.CodeUserNotExist)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"stuNum":    get.StuNum,
		"stuName":   get.StuName,
		"stuGender": get.StuGender,
		"mobile":    get.Mobile,
		"qq":        get.Qq,
		"introduce": get.Introduce,
		"major":     get.Major,
		"province":  get.Province,
		"photo":     get.Photo,
	})
}
