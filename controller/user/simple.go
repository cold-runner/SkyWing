package user

import (
	"Skywing/models/response"
	"Skywing/pkg/aliyunMsg"
	"Skywing/store/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"regexp"
	"time"
)

func (u *UserController) SendSmsCode(c *gin.Context) {
	phone := c.GetString("phone")
	// 校验手机号
	ok, _ := regexp.MatchString("^(?:(?:\\+|00)86)?1\\d{10}$", phone)
	if !ok {
		response.ResponseError(c, response.CodeInvalidPhone)
		c.Abort()
	}
	// 发送短信
	randCode := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	err := aliyunMsg.SendVerificationCode(phone, randCode)
	if err != nil {
		return
	}
	// 存储进redis，有效时间3分钟
	redis.RdbClient.Client.Set(phone, randCode, time.Minute*3)
	zap.L().Info("验证码发送成功", zap.String("code", randCode))
	response.ResponseSuccess(c, nil)
}

// GetCount 获取实时已报名人数
func (u *UserController) GetCount(c *gin.Context) {
	count, err := u.Srv.Users().GetCount()
	if err != nil {
		response.ResponseError(c, response.CodeServerError)
		return
	}
	response.ResponseSuccess(c, gin.H{"count": count})
}
