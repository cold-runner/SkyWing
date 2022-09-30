package captcha

import (
	"Skywing/models/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var Store = base64Captcha.DefaultMemStore

func Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(80, 180, 4, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, Store)
	id, b64s, err := cp.Generate()
	if err != nil {
		zap.L().Error("验证码获取失败!", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParams)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"CaptchaId":     id,
		"PictureString": b64s,
		"CaptchaLength": 4,
	})
}
