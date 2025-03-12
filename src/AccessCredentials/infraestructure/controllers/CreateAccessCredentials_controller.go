package controllers

import (
	"api_accesscredentials/src/AccessCredentials/application"
	"api_accesscredentials/src/AccessCredentials/domain/entities"
	"api_accesscredentials/src/AccessCredentials/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAccessCredentialsController struct {
	useCase *application.CreateAccessCredentials
}

func NewCreateAccessCredentialsController() *CreateAccessCredentialsController {
	mysql := infraestructure.GetMySQL()
	app := application.NewCreateAccessCredentials(mysql)

	return &CreateAccessCredentialsController{useCase: app}
}

func (cac_c *CreateAccessCredentialsController) Run(ctx *gin.Context) {
	var accessCredentials entities.AccessCredentials

	if err := ctx.ShouldBindJSON(&accessCredentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Validamos que los campos no estén vacíos
	if accessCredentials.IdStudent == 0 && accessCredentials.User == "" && accessCredentials.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Los campos están vacíos o son inválidos"})
		return
	}

	err := cac_c.useCase.Run(&accessCredentials)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"Message": "Credencial de acceso asociado a dicho estudiante"})

}