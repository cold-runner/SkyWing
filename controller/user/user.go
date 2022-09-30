package user

import (
	"Skywing/service"
	"Skywing/store"
)

// UserController create a user handler used to handle request for user resource.
type UserController struct {
	Srv service.Service
}

// NewUserController creates a user handler.
func NewUserController(store store.Factory) *UserController {
	return &UserController{
		Srv: service.NewService(store),
	}
}
