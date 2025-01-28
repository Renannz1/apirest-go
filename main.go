package main

import (
	"project_api_go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.ProductRoutes(r)

	r.Run()

	// controllers.GetProductByID(1)
	// controllers.UpdateProductByID(1, "JBL", 2300)
	// controllers.GetProductByID(1)

}
