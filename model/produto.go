package model
type Produto struct {
	ID   int64    `json:"id"`
	Nome string `json:"nome"`
	Preco float64 `json:"preco"`
}