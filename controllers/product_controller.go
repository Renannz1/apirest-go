package controllers

import (
	"net/http"
	"project_api_go/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Criando uma lista do meu model 'Product' e guardando na variável 'products'
// Dados simulados (em vez de um banco de dados real)
var products = []models.Product{
	{ID: 1, Name: "Televisão", Price: 2300},
	{ID: 2, Name: "Celular", Price: 1200},
}

// Listing all Products
func GetProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, products)
}

// Creating new Product
func CreateProduct(ctx *gin.Context) {
	var newProduct models.Product

	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct.ID = uint(len(products) + 1)
	products = append(products, newProduct)
	ctx.JSON(http.StatusCreated, newProduct)
}

// Get a unique Product
func GetProductByID(ctx *gin.Context) {
	var id, err = strconv.Atoi(ctx.Param("id"))

	// Erro na conversão do Parametro da URL
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mostrando na resposta JSON o produto referente ao id
	for _, product := range products {
		if id == int(product.ID) {
			ctx.JSON(http.StatusAccepted, product)
			return
		}
	}

	// Caso o produto não seja encontrado
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
}

// Updating a unique Product
func UpdateProductByID(ctx *gin.Context) {

	// Pegando o id do parametro da url e convertendo para int
	var id, err = strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Pegando os dados passados no JSON e atribuindo para a variavel 'productEdited'
	var productEdited models.Product

	err = ctx.ShouldBindJSON(&productEdited)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Procurando o produto com o id passado pelo parametro
	for i, product := range products {
		if id == int(product.ID) {
			products[i] = productEdited
			products[i].ID = uint(id)
			ctx.JSON(http.StatusOK, gin.H{"message": "product edited.", "product": products[i]})
			return
		}
	}

	// Erro para o caso o produto com o id do parametro nao seja encontrado
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado."})
}

// Deleting product by Id
func DeleteProductByID(ctx *gin.Context) {
	var id, err = strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, product := range products {
		if id == int(product.ID) {
			products = append(products[:i], products[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"product deleted": product})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado."})
}
