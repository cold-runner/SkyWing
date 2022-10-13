package service

import (
	"Skywing/store"
)

// ServiceFunc defines functions used to return resource interface.
type ServiceFunc interface {
	Users() UserSrv
	Policies() PolicySrv
	Roles() RoleSrv
}

type Service struct {
	Store store.Factory
}

// NewService returns Service interface.
func NewService(store store.Factory) ServiceFunc {
	return &Service{
		Store: store,
	}
}

func (s *Service) Users() UserSrv {
	return newUsers(s)
}
func (s *Service) Policies() PolicySrv {
	return newPolicies(s)
}
func (s *Service) Roles() RoleSrv {
	return newRoles(s)
}
