package model

type Product struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Price  uint   `json:"price"`
	Amount uint   `json:"amount"`
}