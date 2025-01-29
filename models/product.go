package models

// Product representa um produto no sistema
// swagger:model Product
type Product struct {
	ID        uint    `json:"id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
}
