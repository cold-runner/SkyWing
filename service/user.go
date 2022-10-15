package service

import (
	"Skywing/models"
	"Skywing/models/response"
	"Skywing/pkg/jwt"
	"Skywing/pkg/oss"
	snow "Skywing/pkg/snowflake"
	"Skywing/store"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// UserSrv defines functions used to handle user request.
type UserSrv interface {
	Create(*gin.Context, *models.RegisterForm) (*models.GenCreateInfo, error)
	Update(string, *models.UpdateForm) error
	Delete(stuNum string) error
	DeleteCollection(stuNum []string) error
	Get(stuNum string) (*models.User, error)
	GetCount() (int, error)
	List() ([]models.User, error)
	Login(loginUser *models.LoginUser) (*models.LoginedUser, error)
	//ChangePassword() error
}
type UserService struct {
	store store.Factory
}

var _ UserSrv = (*UserService)(nil)

func newUsers(srv *Service) *UserService {
	return &UserService{store: srv.Store}
}
func (u *UserService) Create(c *gin.Context, passUser *models.RegisterForm) (*models.GenCreateInfo, error) {
	res, _ := u.store.Users().GetByStuNum(passUser.StuNum)
	if res.StuNum == passUser.StuNum {
		return nil, fmt.Errorf("用户已存在")
	}
	// 生成uuid
	snowId, err := snow.GetID()
	if err != nil {
		zap.L().Error("UUID生成失败", zap.Error(err))
	}
	// 生成加密密码
	passUser.Password = models.EyPasswd(passUser.Password)

	// 照片上传至oss
	_, fileHeader, err := c.Request.FormFile("photo")
	photoUrl, err := oss.UploadToQiNiu(fileHeader, strconv.FormatUint(snowId, 10))
	if err != nil {
		response.ResponseError(c, response.CodeInvalidPhoto)
	}
	// user表中插入数据
	user := &models.User{
		UserID:     snowId,
		CreateTime: time.Now(),
		StuNum:     passUser.StuNum,
		StuName:    passUser.StuName,
		StuGender:  passUser.StuGender,
		Major:      passUser.Major,
		Qq:         passUser.Qq,
		Mobile:     passUser.Mobile,
		Province:   passUser.Province,
		Photo:      photoUrl,
		Introduce:  passUser.Introduce,
		Password:   passUser.Password,
	}
	if err = u.store.Users().Create(user); err != nil {
		return nil, err
	}
	// 生成带有uuid和stuNum的业务结构体
	mic := &models.GenCreateInfo{
		Uuid:   snowId,
		StuNum: passUser.StuNum,
	}
	if err != nil {
		return nil, err
	}
	return mic, nil
}

func (u *UserService) Update(uuid string, user *models.UpdateForm) error {
	// 获取uuid
	existInfo, err := u.store.Users().GetByUuid(uuid)
	if err != nil {
		zap.L().Error("不存在的用户", zap.Error(err))
	}
	// 删除原来照片
	processKey := fmt.Sprintf("userPhoto/%s.jpg", uuid)
	if err = oss.DeleteFileFromQiniu(processKey); err != nil {
		zap.L().Error("删除照片失败！照片可能已经被删除", zap.Error(err))
	}
	// 新的照片
	photoUrl, err := oss.UploadToQiNiu(user.Photo, strconv.FormatUint(existInfo.UserID, 10))
	if err != nil {
		zap.L().Error("更新照片失败！", zap.Error(err))
		return err
	}
	newInfo := &models.User{
		UserID:     existInfo.UserID,
		UpdateTime: time.Now(),
		Photo:      photoUrl,
		Introduce:  user.Introduce,
	}

	if err = u.store.Users().Update(newInfo); err != nil {
		zap.L().Error("更新数据失败", zap.Error(err))
		return err
	}
	return nil
}
func (u *UserService) GetCount() (int, error) {
	count, err := u.store.Users().GetCount()
	if err != nil {
		zap.L().Error("查询报名人数错误", zap.Error(err))
		return -1, err
	}
	return count, nil
}
func (u *UserService) Delete(stuNum string) error {
	if err := u.store.Users().Delete(stuNum); err != nil {
		// 返回数据库内部错误的相应码
	}
	return nil
}
func (u *UserService) DeleteCollection(stuNum []string) error {
	if err := u.store.Users().DeleteCollection(stuNum); err != nil {
		// 返回数据库内部错误的相应码
	}

	return nil
}
func (u *UserService) Get(uuid string) (*models.User, error) {
	user, err := u.store.Users().GetByUuid(uuid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// List returns user list in the storage. This function has a good performance.
func (u *UserService) List() ([]models.User, error) {
	return u.store.Users().List()
}
func (u *UserService) Login(lUser *models.LoginUser) (userInfo *models.LoginedUser, err error) {
	// 在数据库中查询是否存在用户
	get, err := u.store.Users().GetByStuNum(lUser.StuNum)
	if err != nil {
		return nil, fmt.Errorf("用户不存在")
	}
	// 用户登录密码和数据库密码进行匹配
	enPassword := models.EyPasswd(lUser.Password)
	if enPassword != get.Password {
		return nil, fmt.Errorf("密码错误")
	}
	// 登录成功根据uuid和角色信息颁发accessToken
	aToken, err := jwt.GenToken(&models.CustomClaims{
		BaseClaims: models.BaseClaims{
			Uuid:   get.UserID,
			StuNum: get.StuNum,
		},
	})
	if err != nil {
		zap.L().Error("token生成失败！", zap.Error(err))
	}
	loginUser := &models.LoginedUser{
		StuNum:    get.StuNum,
		StuName:   get.StuName,
		StuGender: get.StuGender,
		Major:     get.Major,
		Qq:        get.Qq,
		Mobile:    get.Mobile,
		Province:  get.Province,
		Introduce: get.Introduce,
		Atoken:    aToken,
		Photo:     get.Photo,
	}
	return loginUser, nil
}

//
//func (u *UserService) ChangePassword(ctx context.Context, user *v1.User) error {
//	// Save changed fields.
//	if err := u.store.Users().Update(ctx, user, metav1.UpdateOptions{}); err != nil {
//		return errors.WithCode(code.ErrDatabase, err.Error())
//	}
//
//	return nil
//}
