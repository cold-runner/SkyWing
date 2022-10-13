package jwt

import (
	"Skywing/models"
	"Skywing/settings"
	"go.uber.org/zap"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return []byte(settings.Conf.JwtConf.SigningKey), nil
}

// GenaToken 生成Token
func GenToken(custom *models.CustomClaims) (token string, err error) {
	// 创建自定义声明
	claim := models.CustomClaims{
		BaseClaims: models.BaseClaims{Uuid: custom.Uuid, StuNum: custom.StuNum},
		StandardClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * settings.Conf.JwtConf.ExpiresTime)),
			Issuer:    settings.Conf.JwtConf.Issuer,
		},
		// 指定缓存时间
		BufferTime: time.Now().Add(time.Minute * settings.Conf.JwtConf.BufferTime).Unix(),
	}
	// 加密并获得完整的编码后的字符串token，其中密钥类型必须是[]byte
	secretKey := settings.Conf.JwtConf.SigningKey
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, &claim).SignedString([]byte(secretKey))
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*models.CustomClaims, error) {
	claims := new(models.CustomClaims)
	_, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		zap.L().Error("token解析失败", zap.Error(err))
		return nil, err
	}
	// 校验token有效性
	if err := claims.Valid(); err != nil {
		zap.L().Info("token已失效", zap.Error(err))
		return nil, err
	}
	return claims, nil
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
//func CreateTokenByOldToken(oldToken string, claims *models.CustomClaims) (string, error) {
//	concurrencyControl := &singleflight.Group{}
//	v, err, _ := concurrencyControl.Do("JWT:"+oldToken,
//		func() (interface{}, error) {
//			return GenToken(claims)
//		})
//	return v.(string), err
//}
