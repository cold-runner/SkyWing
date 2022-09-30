package models

type User struct {
	UserID     uint64 `json:"user_id" db:"user_id"`
	CreateTime string `json:"createTime" db:"create_time"`
	UpdateTime string `json:"updateTime" db:"update_time"`
	RegisterForm
}

type LoginUser struct {
	StuNum   string `json:"stuNum" validate:"valStuNum" db:"stu_num"`
	Password string `json:"password" validate:"valStuPassword" db:"password"`
}

type RegisterForm struct {
	StuNum          string `json:"stuNum" db:"stu_num" validate:"valStuNum"`
	StuName         string `json:"stuName" db:"stu_name" validate:"valStuName"`
	StuGender       string `json:"stuGender" db:"stu_gender" validate:"valStuGender"`
	Major           string `json:"major" db:"major" validate:"valStuMajor"`
	Qq              string `json:"qq" db:"qq" validate:"valStuQq"`
	Mobile          string `json:"mobile" db:"mobile" validate:"valStuMobile"`
	Province        string `json:"province" db:"province" validate:"valStuProvince"`
	Photo           string `json:"photo" db:"photo" validate:"base64"`
	Introduce       string `json:"introduce" db:"introduce" validate:"valStuIntroduce"`
	Password        string `json:"password" db:"password" validate:"valStuPassword"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`
}
type LoginedUser struct {
	StuNum       string `json:"stuNum"`
	StuName      string `json:"stuName"`
	StuGender    string `json:"stuGender"`
	Major        string `json:"major"`
	Qq           string `json:"qq"`
	Mobile       string `json:"mobile"`
	Province     string `json:"province"`
	Introduce    string `json:"introduce"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Photo        string `json:"photo"`
}
