package controllers

import (
	"net/http"
	"project_api_go/database"
	"project_api_go/models"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetProducts retorna todos os produtos cadastrados
// @Summary Lista todos os produtos
// @Description Retorna um array de produtos do banco de dados
// @Tags produtos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetProducts(ctx *gin.Context) {
	// definindo a variavel que vai receber os produtos
	var products []models.Product

	database.DB.Find(&products)
	ctx.JSON(http.StatusOK, products)
}

// CreateProduct cria um novo produto
// @Summary Cria um produto
// @Description Adiciona um novo produto ao banco de dados
// @Tags produtos
// @Accept  json
// @Produce  json
// @Param product body models.Product true "Dados do produto"
// @Success 201 {object} models.Product
// @Failure 400 {object} map[string]string "Erro ao processar JSON"
// @Router /products [post]
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

// GetProductByID retorna um produto pelo ID
// @Summary Busca um produto por ID
// @Description Retorna um único produto com base no ID fornecido
// @Tags produtos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do produto"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]string "Produto não encontrado"
// @Router /products/{id} [get]
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

// UpdateProductByID atualiza um produto pelo ID
// @Summary Atualiza um produto
// @Description Modifica os dados de um produto existente com base no ID
// @Tags produtos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do produto"
// @Param product body models.Product true "Dados atualizados do produto"
// @Success 200 {object} map[string]interface{} "Produto atualizado com sucesso"
// @Failure 400 {object} map[string]string "Erro ao processar JSON"
// @Failure 404 {object} map[string]string "Produto não encontrado"
// @Router /products/{id} [put]
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

// DeleteProductByID deleta um produto pelo ID
// @Summary Deleta um produto
// @Description Remove um produto do banco de dados com base no ID fornecido
// @Tags produtos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do produto"
// @Success 200 {object} map[string]interface{} "Produto deletado com sucesso"
// @Failure 404 {object} map[string]string "Produto não encontrado"
// @Router /products/{id} [delete]
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
