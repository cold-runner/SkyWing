package store

import (
	"Skywing/models"
)

// UserStore defines the user storage interface.
type UserStore interface {
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(stuNum string) error
	DeleteCollection(stuNum []string) error
	Get(stuNum string) (*models.User, error)
	List() ([]models.User, error)
}
