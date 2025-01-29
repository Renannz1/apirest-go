package main

import (
	"project_api_go/database"
	"project_api_go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	routes.ProductRoutes(r)

	r.Run()
}
