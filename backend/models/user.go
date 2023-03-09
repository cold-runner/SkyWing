package models

import (
	"Skywing/models/request"
	"Skywing/settings"
	"crypto/md5"
	"encoding/hex"
	"mime/multipart"
	"time"
)

type User struct {
	UserID     uint64    `json:"user_id" db:"user_id" validate:"-"`
	CreateTime time.Time `json:"createTime" db:"create_time" validate:"-"`
	UpdateTime time.Time `json:"updateTime" db:"update_time" validate:"-"`
	StuNum     string    `json:"stuNum" db:"stu_num" validate:"valStuNum"`
	StuName    string    `json:"stuName" db:"stu_name" validate:"valStuName"`
	StuGender  string    `json:"stuGender" db:"stu_gender" validate:"valStuGender"`
	Major      string    `json:"major" db:"major" validate:"valStuMajor"`
	Qq         string    `json:"qq" db:"qq" validate:"valStuQq"`
	Mobile     string    `json:"mobile" db:"mobile" validate:"valStuMobile"`
	Province   string    `json:"province" db:"province" validate:"valStuProvince"`
	Photo      string    `json:"photo" db:"photo"`
	Introduce  string    `json:"introduce" db:"introduce" validate:"valStuIntroduce"`
	Password   string    `json:"password" db:"password" validate:"valStuPassword"`
}

type LoginUser struct {
	StuNum   string `json:"stuNum" validate:"valStuNum" db:"stu_num"`
	Password string `json:"password" validate:"valStuPassword" db:"password"`
	*request.Captcha
}

type RegisterForm struct {
	StuNum          string `json:"stuNum" validate:"valStuNum"`
	StuName         string `json:"stuName" validate:"valStuName"`
	StuGender       string `json:"stuGender" validate:"valStuGender"`
	Major           string `json:"major" validate:"valStuMajor"`
	Qq              string `json:"qq" validate:"valStuQq"`
	Mobile          string `json:"mobile" validate:"valStuMobile"`
	Province        string `json:"province"  validate:"valStuProvince"`
	Introduce       string `json:"introduce"  validate:"valStuIntroduce"`
	Password        string `json:"password"  validate:"valStuPassword"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`
	//SmsCode         string `json:"smsCode" validate:"len=4,alphanum"`
}
type UpdateForm struct {
	//StuNum    string                `json:"stuNum" validate:"valStuNum"`
	//StuName   string                `json:"stuName" validate:"valStuName"`
	//StuGender string                `json:"stuGender" validate:"valStuGender"`
	//Major     string                `json:"major" validate:"valStuMajor"`
	//Qq        string                `json:"qq" validate:"valStuQq"`
	//Mobile    string                `json:"mobile" validate:"valStuMobile"`
	//Province  string                `json:"province" validate:"valStuProvince"`
	Introduce string                `json:"introduce" validate:"valStuIntroduce"`
	Photo     *multipart.FileHeader `json:"photo"`
	*request.Captcha
}

type LoginedUser struct {
	StuNum    string `json:"stuNum"`
	StuName   string `json:"stuName"`
	StuGender string `json:"stuGender"`
	Major     string `json:"major"`
	Qq        string `json:"qq"`
	Mobile    string `json:"mobile"`
	Province  string `json:"province"`
	Introduce string `json:"introduce"`
	Atoken    string `json:"accessToken"`
	Photo     string `json:"photo"`
}
type GenCreateInfo struct {
	Uuid   uint64
	StuNum string
}

func EyPasswd(originPasswd string) (result string) {
	h := md5.New()
	h.Write([]byte(settings.Conf.Secret))
	return hex.EncodeToString(h.Sum([]byte(originPasswd)))
}
