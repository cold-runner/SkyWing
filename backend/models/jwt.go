package models

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	BaseClaims
	StandardClaims jwt.RegisteredClaims
	BufferTime     int64
}

func (c *CustomClaims) Valid() error {
	// 使用库中自带的标准校验函数
	err := c.StandardClaims.Valid()
	if err != nil {
		return err
	}
	return nil
}

type BaseClaims struct {
	Uuid   uint64
	StuNum string
}
