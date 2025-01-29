package main

import (
	"project_api_go/database"
	"project_api_go/routes"

	"github.com/gin-gonic/gin"

	_ "project_api_go/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	routes.ProductRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
