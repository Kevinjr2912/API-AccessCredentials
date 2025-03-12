package repositories

import "api_accesscredentials/src/AccessCredentials/domain/entities"

type IAccessCredentials interface {
	AssociateAccessCredentialsToStudent(accessCredentials *entities.AccessCredentials) (err error)
}
