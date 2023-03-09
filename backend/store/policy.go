package store

import "Skywing/models"

type PolicyStore interface {
	Create(*models.Policy) error
	CreateCollection([]*models.Policy) error
	// Get Update(user *models.RoleCharacter) error
	//Delete(stuNum string) error
	//DeleteCollection(stuNum []string) error
	//Get(uint64) (*models.RoleCharacter, error)
	//List() ([]models.User, error)
}
