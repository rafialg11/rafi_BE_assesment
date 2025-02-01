package entities

type Transaction struct {
	Id              int      `json:"id" gorm:"primaryKey"`
	Amount          int      `json:"amount" gorm:"not null"`
	TransactionType string   `json:"transaction_type" gorm:"not null"`
	CustomerId      int      `json:"customer_id" gorm:"not null"`
	CreatedAt       string   `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Customer        Customer `json:"customer" gorm:"onDelete:CASCADE onUpdate:CASCADE"`
}
