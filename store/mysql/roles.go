package mysql

import (
	"Skywing/models"
	"github.com/jmoiron/sqlx"
)

type roles struct {
	db *sqlx.DB
}

func newRoles(ds *datastore) *roles {
	return &roles{ds.db}
}

func (r *roles) Create(character *models.RoleCharacter) error {
	sqlStr := "insert into role_character(uuid, role) values (?, ?)"
	_, err := r.db.Exec(sqlStr, character.Uuid, character.Role)
	if err != nil {
		return err
	}
	return nil
}
func (r *roles) Get(uuid uint64) (*models.RoleCharacter, error) {
	sqlStr := "select * from role_character where uuid = ?"
	rc := &models.RoleCharacter{}
	err := r.db.Get(rc, sqlStr, uuid)
	if err != nil {
		return nil, err
	}
	return rc, nil
}
