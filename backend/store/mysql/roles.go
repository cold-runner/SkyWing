package mysql

import (
	"Skywing/models"
	"github.com/jmoiron/sqlx"
)

type Roles struct {
	db *sqlx.DB
}

func NewRoles(ds *Datastore) *Roles {
	return &Roles{ds.Db}
}

// Create创建
func (r *Roles) Create(character *models.RoleCharacter) error {
	sqlStr := "insert into role_character(uuid, role) values (?, ?)"
	_, err := r.db.Exec(sqlStr, character.Uuid, character.Role)
	if err != nil {
		return err
	}
	return nil
}
func (r *Roles) Get(uuid uint64) (*models.RoleCharacter, error) {
	sqlStr := "select * from role_character where uuid = ?"
	rc := &models.RoleCharacter{}
	err := r.db.Get(rc, sqlStr, uuid)
	if err != nil {
		return nil, err
	}
	return rc, nil
}
