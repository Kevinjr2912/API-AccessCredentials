package routes

import (
	"api_accesscredentials/src/AccessCredentials/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	routes := router.Group("/accessCredentials") // Definimos el nombre del recurso
	{
		// Definimos las rutas
		routes.POST("", controllers.NewCreateAccessCredentialsController().Run)
	}

}