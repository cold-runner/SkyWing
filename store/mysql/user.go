package mysql

import (
	"Skywing/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type users struct {
	db *sqlx.DB
}

func newUsers(ds *datastore) *users {
	return &users{ds.db}
}

// Create creates a new user account.
func (u *users) Create(reg *models.User) error {
	sqlStr := "insert into user(user_id, stu_num, stu_name, stu_gender,password, major, qq, mobile, province, photo, introduce, create_time) values (:userId,:stuNum,:stuName,:stuGender,:password,:major,:qq,:mobile,:province,:photo,:introduce,:createTime)"
	_, err := u.db.NamedExec(sqlStr, map[string]interface{}{
		"userId":     reg.UserID,
		"stuNum":     reg.StuNum,
		"stuName":    reg.StuName,
		"stuGender":  reg.StuGender,
		"password":   reg.Password,
		"major":      reg.Major,
		"qq":         reg.Qq,
		"mobile":     reg.Mobile,
		"province":   reg.Province,
		"photo":      reg.Photo,
		"introduce":  reg.Introduce,
		"createTime": time.Now(),
	})
	return err
}

// Update updates a user account information.
func (u *users) Update(reg *models.User) error {
	sqlStr := "update user set stu_num = :stuNum, stu_name = :stuName, stu_gender = :stuGender, major = :major, qq = :qq, mobile = :mobile, province = :province, photo = :photo, introduce = :introduce"
	_, err := u.db.NamedExec(sqlStr, map[string]interface{}{
		"stuNum":    reg.StuNum,
		"stuName":   reg.StuName,
		"stuGender": reg.StuGender,
		"major":     reg.Major,
		"qq":        reg.Qq,
		"mobile":    reg.Mobile,
		"province":  reg.Province,
		"photo":     reg.Photo,
		"introduce": reg.Introduce,
	})
	return err

}

// Delete deletes the user by the user identifier.
func (u *users) Delete(stuNum string) error {
	// delete related policy first

	return nil
}

// DeleteCollection batch deletes the user.
func (u *users) DeleteCollection(stuNum []string) error {
	// delete related policy first
	return nil
}

// Get return a user by the user identifier.
func (u *users) Get(StuNum string) (*models.User, error) {
	tar := &models.User{}
	sqlStr := "select * from user where stu_num = ?"
	err := u.db.Get(tar, sqlStr, StuNum)
	return tar, err
}

// List return all user. No constraints, need to be optimized
func (u *users) List() ([]models.User, error) {
	var ret []models.User
	sqlStr := "select * from user"
	err := u.db.Select(&ret, sqlStr)
	return ret, err
}
