package jwt

import (
	"Skywing/models"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
)

func TestGenToken(t *testing.T) {
	testStruct := models.CustomClaims{
		BaseClaims: models.BaseClaims{
			Uuid:   1234,
			StuNum: "22999001",
		},
		StandardClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
			Issuer:    "zz",
		},
		BufferTime: 0,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &testStruct).SignedString([]byte("iii"))

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("token生成成功\n%v", token)
		fmt.Println(token)
	}
}
func TestParseToken(t *testing.T) {
	claims := new(models.CustomClaims)
	tok := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVdWlkIjoyMDc3NzU0MzUwNjM0OTgyLCJTdHVOdW0iOiIyMjk5OTAwNCIsIlN0YW5kYXJkQ2xhaW1zIjp7ImlzcyI6IlNreUxhYiIsImV4cCI6MTY2NTY2NzE4OX0sIkJ1ZmZlclRpbWUiOjE2NjU2Nzc5ODl9.9SGlwLb3hbgiHT1XmL6gcswkUVokCwgPekXGlvUdcuU"
	_, err := jwt.ParseWithClaims(tok,
		claims,
		func(_ *jwt.Token) (interface{}, error) {
			return []byte("sky2022Wel@02@"), nil
		})
	if err != nil {
		t.Error("token解析失败", err)
	}

}
