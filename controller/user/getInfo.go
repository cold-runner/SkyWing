package user

import (
	"Skywing/models/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
)

func (u *UserController) GetInfo(c *gin.Context) {
	get, err := u.Srv.Users().Get(c.Query("stuNum"))
	if err != nil {
		zap.L().Error("查询信息失败", zap.Error(err))
		return
	}
	// 处理照片

	path := fmt.Sprintf("photo/%s.jpg", get.StuName)
	photoFile, err := ioutil.ReadFile(path)
	if err != nil {
		zap.L().Error("读取照片文件失败, ", zap.Error(err))
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
		"photo":     photoFile,
	})
}
