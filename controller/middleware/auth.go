package middleware

import (
	"Skywing/models/response"
	"Skywing/pkg/jwt"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		// 解析并校验token
		mc, err := jwt.ParseToken(authHeader)
		if err != nil {
			response.ResponseError(c, response.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的uuid和角色名信息保存到请求的上下文c上
		c.Set("uuid", mc.Uuid)
		c.Set("stuNum", mc.StuNum)
		c.Next() // 后续的处理函数可以用过c.Get("uuid")来获取当前请求的用户信息
	}
}
