package models

import "gorm.io/gorm"

// Estrutura model do Produto
type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
