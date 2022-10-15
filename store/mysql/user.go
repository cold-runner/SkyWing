package mysql

import (
	"Skywing/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type Users struct {
	db *sqlx.DB
}

func newUsers(ds *Datastore) *Users {
	return &Users{ds.Db}
}

// Create creates a new user account.
func (u *Users) Create(reg *models.User) error {
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
func (u *Users) GetCount() (int, error) {
	var count int
	sqlStr := "select count(stu_num) from user"
	if err := u.db.Get(&count, sqlStr); err != nil {
		return -1, err
	}
	return count, nil
}

// Get return a user by the user identifier.
func (u *Users) GetByUuid(uuid string) (*models.User, error) {
	tar := &models.User{}
	sqlStr := "select * from user where user_id = ?"
	err := u.db.Get(tar, sqlStr, uuid)
	return tar, err
}
func (u *Users) GetByStuNum(stuNum string) (*models.User, error) {
	tar := &models.User{}
	sqlStr := "select * from user where stu_num = ?"
	err := u.db.Get(tar, sqlStr, stuNum)
	return tar, err
}

// List return all user. No constraints, need to be optimized
func (u *Users) List() ([]models.User, error) {
	var ret []models.User
	sqlStr := "select * from user"
	err := u.db.Select(&ret, sqlStr)
	return ret, err
}

// Update updates a user account information.
func (u *Users) Update(user *models.User) error {
	sqlStr := "update user set photo = :photo, introduce = :introduce, update_time= :updateTime where user_id = :uuid"
	_, err := u.db.NamedExec(sqlStr, map[string]interface{}{
		"updateTime": time.Now(),
		"uuid":       user.UserID,
		"photo":      user.Photo,
		"introduce":  user.Introduce,
	})
	return err
}

// Delete deletes the user by the user identifier.
func (u *Users) Delete(string) error {
	// delete related policy first

	return nil
}

// DeleteCollection batch deletes the user.
func (u *Users) DeleteCollection([]string) error {
	// delete related policy first
	return nil
}
