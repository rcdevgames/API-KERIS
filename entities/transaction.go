package entities

type TransactionInput struct {
	Amount int `json:"amount" binding:"required"`
}
