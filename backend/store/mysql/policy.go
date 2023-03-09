package mysql

import (
	"Skywing/models"
	"github.com/jmoiron/sqlx"
)

type Policies struct {
	db *sqlx.DB
}

func NewPolicies(ds *Datastore) *Policies {
	return &Policies{ds.Db}
}

func (p *Policies) Create(policy *models.Policy) error {
	sqlStr := "insert into casbin_rule (p_type, v0, v1, v2) values (?,?,?,?)"
	_, err := p.db.Exec(sqlStr, policy.PType, policy.V0, policy.V1, policy.V2)
	if err != nil {
		return err
	}
	return nil
}
func (p *Policies) CreateCollection(policies []*models.Policy) error {
	for i := 0; i < len(policies); i++ {
		if err := p.Create(policies[i]); err != nil {
			return err
		}
	}
	return nil
}
