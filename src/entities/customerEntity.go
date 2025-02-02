package entities

type Customer struct {
	Id          int           `json:"id" gorm:"primaryKey"`
	Name        string        `json:"name" gorm:"not null"`
	Phone       string        `json:"phone" gorm:"unique; not null"`
	NIK         string        `json:"NIK" gorm:"unique; not null"`
	Transaction []Transaction `json:"transaction" gorm:"foreignKey:CustomerId"`
	Account     Account       `json:"account" gorm:"foreignKey:CustomerId;constraint:OnDelete:CASCADE;"`
}
