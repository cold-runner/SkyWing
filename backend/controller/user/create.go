package user

import (
	"Skywing/controller"
	"Skywing/models"
	res "Skywing/models/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *UserController) Create(c *gin.Context) {
	var err error
	// 获取请求参数
	fo := models.RegisterForm{
		StuNum:          c.PostForm("stuNum"),
		StuName:         c.PostForm("stuName"),
		StuGender:       c.PostForm("stuGender"),
		Major:           c.PostForm("major"),
		Qq:              c.PostForm("qq"),
		Mobile:          c.PostForm("mobile"),
		Province:        c.PostForm("province"),
		Introduce:       c.PostForm("introduce"),
		Password:        c.PostForm("password"),
		ConfirmPassword: c.PostForm("confirm_password"),
		//SmsCode:         c.PostForm("smsCode"),
	}
	// 校验数据有效性
	if err = controller.Val.Struct(fo); err != nil {
		zap.L().Info("参数校验失败", zap.Error(err))
		res.ResponseError(c, res.CodeInvalidParams)
		return
	}

	// 校验手机验证码，签名和模板过审后上线
	//randCode, err := aliyunMsg.GetVerificationCode(fo.Mobile)
	//if err != nil {
	//	zap.L().Error("手机验证码获取失败", zap.Error(err))
	//	return
	//}
	//if randCode != fo.SmsCode {
	//	res.ResponseError(c, res.CodeInvalidRandCode)
	//	return
	//}

	// 校验照片文件
	_, _, err = c.Request.FormFile("photo")
	if err != nil {
		zap.L().Error("照片文件校验失败", zap.Error(err))
		res.ResponseError(c, res.CodeInvalidPhoto)
		return
	}

	// 注册用户
	genCreateI, err := u.Srv.Users().Create(c, &fo)
	if err != nil {
		zap.L().Error("注册用户失败", zap.Error(err))
		res.ResponseError(c, res.CodeUserExist)
		return
	}
	// 萌新角色信息插入数据库
	if err = u.Srv.Roles().Create(genCreateI); err != nil {
		zap.L().Error("角色信息创建失败", zap.Error(err))
		return
	}
	// 创建授权策略
	if err = u.Srv.Policies().Create(genCreateI); err != nil {
		zap.L().Error("创建授权策略失败", zap.Error(err))
	}
	res.ResponseSuccess(c, nil)
}
