package usecases

import (
	"api_accesscredentials/src/AccessCredentials/domain/entities"
	"api_accesscredentials/src/AccessCredentials/domain/repositories"
)

type GetStudentsWithAccessCredentials struct {
	db repositories.IAccessCredentials
}

func NewGetStudentsWithAccessCredentials(db repositories.IAccessCredentials) *GetStudentsWithAccessCredentials {
	return &GetStudentsWithAccessCredentials{db: db}
}

func (gsac *GetStudentsWithAccessCredentials) Run() (studentsAC *[]entities.StudentAccessCredentials, err error) {
	return gsac.db.GetAccessCredentialsAllStudents()
}
