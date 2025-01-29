package controllers

import (
	"net/http"
	"project_api_go/database"
	"project_api_go/models"

	"github.com/gin-gonic/gin"
)

func GetProducts(ctx *gin.Context) {
	// definindo a variavel que vai receber os produtos
	var products []models.Product

	database.DB.Find(&products)
	ctx.JSON(http.StatusOK, products)
}

func CreateProduct(ctx *gin.Context) {
	// Definindo a variavel que vai receber os dados do produto
	var newProduct models.Product

	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&newProduct)
	ctx.JSON(http.StatusCreated, gin.H{"produto criado": newProduct})
}

func GetProductByID(ctx *gin.Context) {
	var product models.Product

	// pegando o id passado na url e armazenando na variavel id
	var id = ctx.Param("id")

	var err = database.DB.First(&product, id).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func UpdateProductByID(ctx *gin.Context) {
	var product models.Product
	var id = ctx.Param("id")

	var err = database.DB.First(&product, id).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "produto não encontrado"})
		return
	}

	err = ctx.ShouldBindBodyWithJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&product)

	ctx.JSON(http.StatusOK, gin.H{"produto atualizado": product})
}

func DeleteProductByID(ctx *gin.Context) {
	var product models.Product
	var id = ctx.Param("id")

	var err = database.DB.First(&product, id).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "produto não encontrado"})
	}

	database.DB.Delete(&product)

	ctx.JSON(http.StatusOK, gin.H{"produto deletado": product})
}
