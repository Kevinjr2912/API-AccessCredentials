package main

import (
	"api_accesscredentials/src/AccessCredentials/infraestructure"
	"api_accesscredentials/src/AccessCredentials/infraestructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	infraestructure.InitDependencies()

	// Creamos el router
	r := gin.Default()

	// CORS
	r.Use(cors.Default())

	// Registramos las rutas
	routes.RegisterRoutes(r)

	// Levantamos el servidor
	r.Run(":8081")
}