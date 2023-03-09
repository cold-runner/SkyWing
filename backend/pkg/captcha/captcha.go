package captcha

import (
	"Skywing/models/response"
	"Skywing/store/redis"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

func Captcha(c *gin.Context) {
	// 验证码配置
	// 生成默认数字的driver，redis实现Store接口
	driver := base64Captcha.NewDriverDigit(80, 180, 4, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, redis.RdbClient)
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
