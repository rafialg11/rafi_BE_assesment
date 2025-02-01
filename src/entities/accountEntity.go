package entities

type Account struct {
	Id            int       `json:"id" gorm:"primaryKey"`
	AccountNumber string    `json:"account_number" gorm:"unique; not null"`
	Amount        int       `json:"amount" gorm:"not null"`
	CustomerId    int       `json:"customer_id" gorm:"unique; not null"`
	CreatedAt     string    `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Customer      *Customer `json:"customer" gorm:"onDelete:CASCADE onUpdate:CASCADE"`
}
