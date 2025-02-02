package entities

import "time"

type Transaction struct {
	Id              int        `json:"id" gorm:"primaryKey"`
	Amount          int        `json:"amount" gorm:"not null"`
	TransactionType string     `json:"transaction_type" gorm:"not null"`
	CustomerId      int        `json:"customer_id" gorm:"not null"`
	CreatedAt       *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Customer        Customer   `json:"customer" gorm:"onDelete:CASCADE onUpdate:CASCADE"`
}

type TransactionRequest struct {
	Amount        int    `json:"amount"`
	AccountNumber string `json:"account_number"`
}
