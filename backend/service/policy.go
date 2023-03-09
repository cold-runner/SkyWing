package service

import (
	"Skywing/models"
	"Skywing/store"
	"fmt"
	"strconv"
)

type PolicySrv interface {
	Create(*models.GenCreateInfo) error
}
type PolicyService struct {
	store store.Factory
}

var _ PolicySrv = (*PolicyService)(nil)

func newPolicies(srv *Service) *PolicyService {
	return &PolicyService{store: srv.Store}
}

func (p *PolicyService) Create(obj *models.GenCreateInfo) error {
	uuid := strconv.FormatUint(obj.Uuid, 10)
	policies := []*models.Policy{
		{
			PType: "p",
			V0:    uuid,
			V1:    fmt.Sprintf("/api/v1/update/%s", uuid),
			V2:    "PUT",
		},
		{
			PType: "p",
			V0:    uuid,
			V1:    fmt.Sprintf("/api/v1/info/%s", uuid),
			V2:    "GET",
		},
	}
	if err := p.store.Policies().CreateCollection(policies); err != nil {
		return err
	}
	return nil
}
