package controllers

import (
	"api_accesscredentials/src/AccessCredentials/application/services"
	usecases "api_accesscredentials/src/AccessCredentials/application/useCases"
	"api_accesscredentials/src/AccessCredentials/domain/entities"
	"api_accesscredentials/src/AccessCredentials/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAccessCredentialsController struct {
	useCase *usecases.CreateAccessCredentials
	event *services.Event
}

func NewCreateAccessCredentialsController() *CreateAccessCredentialsController {
	// MySQL
	mysql := infraestructure.GetMySQL()
	app := usecases.NewCreateAccessCredentials(mysql)

	// Rabbit
	rabbit := infraestructure.GetRabbit()
	event := services.NewEvent(rabbit)

	return &CreateAccessCredentialsController{useCase: app, event: event}
}

func (cac_c *CreateAccessCredentialsController) Run(ctx *gin.Context) {
	var studentAC entities.AccessCredentials

	if err := ctx.ShouldBindJSON(&studentAC); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Validamos que los campos no estén vacíos
	if studentAC.IdStudent == 0 && studentAC.Email == "" && studentAC.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Los campos están vacíos o son inválidos"})
		return
	}

	var accessCredentials entities.AccessCredentials

	accessCredentials.IdStudent   = studentAC.IdStudent
	accessCredentials.Email       = studentAC.Email
	accessCredentials.Password    = studentAC.Password

	err := cac_c.useCase.Run(&accessCredentials)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	// Enviamos el email de dicho estudiante nuevo al exchange
	cac_c.event.Run(studentAC.Email)

	ctx.JSON(http.StatusCreated, gin.H{"Message": "Credencial de acceso asociado a dicho estudiante"})

}