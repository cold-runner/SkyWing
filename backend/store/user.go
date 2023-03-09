package store

import (
	"Skywing/models"
)

// UserStore defines the user storage interface.
type UserStore interface {
	Create(*models.User) error
	Update(*models.User) error
	Delete(string) error
	DeleteCollection([]string) error
	GetByStuNum(string) (*models.User, error)
	GetByUuid(string) (*models.User, error)
	GetCount() (int, error)
	List() ([]models.User, error)
}
