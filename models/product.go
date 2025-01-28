package models

// Estrutura model do Produto
type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int32  `json:"price"`
}
