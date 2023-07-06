package entities

import "time"

type Merchant struct {
	ID                 int       `json:"id" gorm:"type:int;primary_key"`
	UserID             int       `json:"user_id"`
	QrisData           string    `json:"qris_data"`
	QrisEmail          string    `json:"qris_email"`
	QrisPassword       string    `json:"qris_password"`
	QrisMnid           string    `json:"qris_mnid"`
	QrisName           string    `json:"qris_name"`
	LastMutationDate   string    `json:"last_mutation_date"`
	LastMutationAmount int       `json:"last_mutation_amount"`
	LastMutationFound  int       `json:"last_mutation_found"`
	CreatedDate        time.Time `json:"created_date"`
	CreatedBy          string    `json:"created_by"`
	UpdatedDate        time.Time `json:"updated_date"`
	UpdatedBy          string    `json:"updated_by"`
}

type MerchantInsert struct {
	UserID       int    `json:"user_id"`
	QrisData     string `json:"qris_data" binding:"required" gorm:"unique"`
	QrisEmail    string `json:"qris_email" binding:"required" gorm:"unique"`
	QrisPassword string `json:"qris_password" binding:"required"`
	QrisMnid     string `json:"qris_mnid"`
	QrisName     string `json:"qris_name"`
}
