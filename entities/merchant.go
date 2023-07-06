package entities

import "time"

type Merchant struct {
	ID                 int64     `json:"id" gorm:"type:int;primary_key"`
	UserID             int64     `json:"user_id"`
	QrisData           string    `json:"qris_data"`
	QrisEmail          string    `json:"qris_email"`
	QrisPassword       string    `json:"qris_password"`
	LastMutationDate   string    `json:"last_mutation_date"`
	LastMutationAmount int64     `json:"last_mutation_amount"`
	LastMutationFound  int8      `json:"last_mutation_found"`
	CreatedDate        time.Time `json:"created_date"`
	CreatedBy          string    `json:"created_by"`
	UpdatedDate        time.Time `json:"updated_date"`
	UpdatedBy          string    `json:"updated_by"`
}

type MerchantInsert struct {
	UserID       int64  `json:"user_id"`
	QrisData     string `json:"qris_data"`
	QrisEmail    string `json:"qris_email"`
	QrisPassword string `json:"qris_password"`
}
