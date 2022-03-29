package services

import "example.com/sarang-apis/models"

type PersonService interface {
	CreatePerson(*models.Person) error
	GetPerson(*int) (*models.Person, error)
	GetAll() ([]*models.Person, error)
	UpdatePerson(*models.Person) error
	DeletePerson(*int) error
}
