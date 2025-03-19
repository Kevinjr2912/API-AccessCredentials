package responses

import (
	"api_accesscredentials/src/AccessCredentials/domain/entities"
	"fmt"
)

type ResponseLinkStudents struct {
	Self string `json:"self"`
}

type ResponseDataStudents struct {
	Type       string `json:"type"`
	Id         string `json:"id"`
	Attributes struct {
		Name        string `json:"name"`
		Age         string `json:"age"`
		PhoneNumber string `json:"phoneNumber"`
		Email       string `json:"email"`
	} `json:"attributes"`
}

type ResponseGetAllStudents struct {
	Links ResponseLinkStudents   `json:"links"`
	Data  []ResponseDataStudents `json:"data"`
}

func NewResponseGetAllStudents(students *[]entities.StudentAccessCredentials) *ResponseGetAllStudents {
	data := []ResponseDataStudents{}

	for i := 0; i < len(*students); i++ {
		data = append(data, ResponseDataStudents{
			Type: "Students",
			Id:   fmt.Sprintf("%d", (*students)[i].IdStudent),
			Attributes: struct {
				Name        string `json:"name"`
				Age         string `json:"age"`
				PhoneNumber string `json:"phoneNumber"`
				Email       string `json:"email"`
			}{
				Name:        fmt.Sprintf((*students)[i].Name),
				Age:         fmt.Sprintf("%d", (*students)[i].Age),
				PhoneNumber: fmt.Sprintf("%d", (*students)[i].PhoneNumber),
				Email:       fmt.Sprintf((*students)[i].Email),
			},
		})
	}

	return &ResponseGetAllStudents{
		Links: ResponseLinkStudents{
			Self: "http://localhost:8081/accessCredentials",
		},
		Data: data,
	}
}
