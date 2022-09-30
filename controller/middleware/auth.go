package middleware

import (
	"Skywing/models/response"
	"Skywing/pkg/jwt"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey = "userID"
)

var (
	ErrorUserNotLogin = errors.New("当前用户未登录")
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// Token放在Header的Authorization中
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.ResponseErrorWithMsg(c, response.CodeInvalidToken, "请求头缺少Auth Token")
			c.Abort()
			return
		}
		// 按.分割
		parts := strings.SplitN(authHeader, ".", 3)
		if !(len(parts) == 3) {
			response.ResponseErrorWithMsg(c, response.CodeInvalidToken, "Token格式不对")
			c.Abort()
			return
		}
		// 解析JWT
		mc, err := jwt.ParseToken(authHeader)
		if err != nil {
			fmt.Println(err)
			response.ResponseError(c, response.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(ContextUserIDKey, mc.UserID)
		c.Next() // 后续的处理函数可以用过c.Get("userID")来获取当前请求的用户信息
	}
}
