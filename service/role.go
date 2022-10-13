package service

import (
	"Skywing/models"
	"Skywing/store"
)

// UserSrv defines functions used to handle user request.
type RoleSrv interface {
	Create(*models.GenCreateInfo) error
}
type RoleService struct {
	store store.Factory
}

var _ RoleSrv = (*RoleService)(nil)

func newRoles(srv *Service) *RoleService {
	return &RoleService{store: srv.Store}
}
func (r *RoleService) Create(obj *models.GenCreateInfo) error {
	role := &models.RoleCharacter{
		Uuid: obj.Uuid,
		Role: "newcomer",
	}
	if err := r.store.Roles().Create(role); err != nil {
		return err
	}
	return nil
}
