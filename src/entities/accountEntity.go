package entities

import "time"

type Account struct {
	Id            int        `json:"id" gorm:"primaryKey"`
	AccountNumber string     `json:"account_number" gorm:"unique; not null"`
	Amount        int        `json:"amount" gorm:"not null"`
	CustomerId    int        `json:"customer_id" gorm:"not null"`
	CreatedAt     *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Customer      *Customer  `json:"customer" gorm:"foreignKey:CustomerId;constraint:OnDelete:CASCADE;"`
}
