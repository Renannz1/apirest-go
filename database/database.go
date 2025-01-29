package database

import (
	"fmt"
	"project_api_go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "user_apirest_go:GodABe@tcp(127.0.0.1:3306)/apirest_go?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Migrar os modelos para o banco
	database.AutoMigrate(&models.Product{})

	DB = database
}
