package routes

import (
	"project_api_go/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	var ProductGroup = router.Group("/products")
	{
		ProductGroup.GET("/", controllers.GetProducts)
		ProductGroup.POST("/", controllers.CreateProduct)
		ProductGroup.GET("/:id", controllers.GetProductByID)
		ProductGroup.PUT("/:id", controllers.UpdateProductByID)
		ProductGroup.DELETE("/:id", controllers.DeleteProductByID)
	}
}
