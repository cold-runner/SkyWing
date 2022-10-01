package store

import "Skywing/models"

type RoleStore interface {
	Create(*models.RoleCharacter) error
	//Update(user *models.RoleCharacter) error
	//Delete(stuNum string) error
	//DeleteCollection(stuNum []string) error
	Get(uint64) (*models.RoleCharacter, error)
	//List() ([]models.User, error)
}
