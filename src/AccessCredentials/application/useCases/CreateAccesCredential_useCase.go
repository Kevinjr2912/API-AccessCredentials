package usecases

import (
	"api_accesscredentials/src/AccessCredentials/domain/entities"
	"api_accesscredentials/src/AccessCredentials/domain/repositories"
)

type CreateAccessCredentials struct {
	db repositories.IAccessCredentials
}

func NewCreateAccessCredentials(db repositories.IAccessCredentials) *CreateAccessCredentials {
	return &CreateAccessCredentials{db: db}
}

func (cac *CreateAccessCredentials) Run(accessCredentials *entities.AccessCredentials) (err error) {
	return cac.db.AssociateAccessCredentialsToStudent(accessCredentials)
}