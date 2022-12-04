package models

import "time"

type Transaction struct {
	ID         int64     `json:"id"`
	UserID     int       `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User       User      `json:"user"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      int       `json:"phone"`
	PostalCode int       `json:"postalcode"`
	Address    string    `json:"address"`
	Status     string    `json:"status"`
	Total      int       `json:"total"`
	Carts      []Cart    `json:"carts"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}

type TransactionResponse struct {
	ID     int64 `json:"id"`
	UserID int   `json:"user_id"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
