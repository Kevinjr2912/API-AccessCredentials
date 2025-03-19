package controllers

import (
	usecases "api_accesscredentials/src/AccessCredentials/application/useCases"
	"api_accesscredentials/src/AccessCredentials/infraestructure"
	"api_accesscredentials/src/AccessCredentials/infraestructure/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetStudentsAccessCredentialsController struct {
	useCase *usecases.GetStudentsWithAccessCredentials
}

func NewGetStudentsAccessCredentialsController() *GetStudentsAccessCredentialsController {
	mysql := infraestructure.GetMySQL()
	app := usecases.NewGetStudentsWithAccessCredentials(mysql)

	return &GetStudentsAccessCredentialsController{useCase: app}
}

func (gsac_c *GetStudentsAccessCredentialsController) Run(ctx *gin.Context) {
	studentsWithAccessCredential, err := gsac_c.useCase.Run()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	reponse := responses.NewResponseGetAllStudents(studentsWithAccessCredential)

	ctx.JSON(http.StatusOK, reponse)
}