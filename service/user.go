package service

import (
	"Skywing/models"
	"Skywing/pkg/jwt"
	snow "Skywing/pkg/snowflake"
	"Skywing/store"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
)

// UserSrv defines functions used to handle user request.
type UserSrv interface {
	Create(c *gin.Context, user *models.RegisterForm) error
	Update(user *models.User) error
	Delete(stuNum string) error
	DeleteCollection(stuNum []string) error
	Get(username string) (*models.User, error)
	List() ([]models.User, error)

	Login(ctx *gin.Context, loginUser *models.LoginUser) (*models.LoginedUser, error)
	//ChangePassword() error
}

type userService struct {
	store store.Factory
}

var _ UserSrv = (*userService)(nil)

func newUsers(srv *service) *userService {
	return &userService{store: srv.store}
}

// List returns user list in the storage. This function has a good performance.
func (u *userService) List() ([]models.User, error) {
	return u.store.Users().List()
}

func (u *userService) Create(c *gin.Context, passUser *models.RegisterForm) error {
	res, _ := u.store.Users().Get(passUser.StuNum)

	if res.StuNum == passUser.StuNum {
		return fmt.Errorf("用户已存在")
	}

	// 生成加密密码
	enPassword := encryptPassword([]byte(passUser.Password))
	passUser.Password = enPassword

	// 将照片解码成图片，并将路径存储
	ddd, err := base64.StdEncoding.DecodeString(passUser.Photo)
	if err != nil {
		zap.L().Error("图片转base64失败, ", zap.Error(err))
	}
	path := fmt.Sprintf("photo/%s.jpg", passUser.StuName)
	if err := ioutil.WriteFile(path, ddd, 0666); err != nil {
		zap.L().Error("照片写入失败, ", zap.Error(err))
	}
	passUser.Photo = path

	snowId, err := snow.GetID()
	if err != nil {
		// id随机生成失败
	}
	var user = &models.User{
		UserID:       snowId,
		RegisterForm: *passUser,
	}
	// 数据库插入数据
	err = u.store.Users().Create(user)
	return err
}

func (u *userService) DeleteCollection(stuNum []string) error {
	if err := u.store.Users().DeleteCollection(stuNum); err != nil {
		// 返回数据库内部错误的相应码
	}

	return nil
}

func (u *userService) Delete(stuNum string) error {
	if err := u.store.Users().Delete(stuNum); err != nil {
		// 返回数据库内部错误的相应码
	}
	return nil
}

func (u *userService) Get(username string) (*models.User, error) {
	user, err := u.store.Users().Get(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) Update(user *models.User) error {
	if err := u.store.Users().Update(user); err != nil {

	}
	return nil
}
func (u *userService) Login(c *gin.Context, loginUser *models.LoginUser) (userInfo *models.LoginedUser, err error) {
	// 在数据库中查询是否存在用户
	get, err := u.store.Users().Get(loginUser.StuNum)
	if err != nil {
		return nil, fmt.Errorf("用户不存在")
	}
	// 用户登录密码和数据库密码进行匹配
	enPassword := encryptPassword([]byte(loginUser.Password))
	if enPassword != get.Password {
		return nil, fmt.Errorf("密码错误")
	}
	// 登录成功颁发token
	accessToken, refreshToken, err := jwt.GenToken(get.UserID)
	if err != nil {
		zap.L().Error("token生成失败", zap.Error(err))
		return nil, err
	}
	// 根据数据库中存储的照片路径找到照片，并编码为base64
	path := fmt.Sprintf("photo/%s.jpg", get.StuName)
	photoFile, err := ioutil.ReadFile(path)
	if err != nil {
		zap.L().Error("读取照片文件失败, ", zap.Error(err))
	}
	ddd := base64.StdEncoding.EncodeToString([]byte(photoFile))

	loginedUser := &models.LoginedUser{
		StuNum:       get.StuNum,
		StuName:      get.StuName,
		StuGender:    get.StuGender,
		Major:        get.Major,
		Qq:           get.Qq,
		Mobile:       get.Mobile,
		Province:     get.Province,
		Introduce:    get.Introduce,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Photo:        ddd,
	}
	return loginedUser, nil
}

//
//func (u *userService) ChangePassword(ctx context.Context, user *v1.User) error {
//	// Save changed fields.
//	if err := u.store.Users().Update(ctx, user, metav1.UpdateOptions{}); err != nil {
//		return errors.WithCode(code.ErrDatabase, err.Error())
//	}
//
//	return nil
//}
