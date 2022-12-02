package models

import "time"

type Product struct {
	ID       int           `json:"id" gorm:"primary_key:auto_increment"`
	Title    string        `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price    int           `json:"price" form:"price" gorm:"type: int"`
	Image    string        `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty			 int					 `json:"qty" form:"qty" gorm:"type: varchar(255)"`
	CreateAt time.Time		 `json:"-"`
	UpdateAt time.Time		 `json:"-"`
}

type ProductResponse struct {
	ID       int           			`json:"id"`
	Title    string        			`json:"title"`
	Price    int           			`json:"price"`
	Image    string        			`json:"-"`
	Qty			 int					 			`json:"-"`
	// UserID   int           			`json:"-"`
	// User     UserResponse 			`json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// // Topping  []ToppingResponse 	`json:"topping"`
	// CreateAt time.Time		 			`json:"-"`
	// UpdateAt time.Time		 			`json:"-"`
}

type ProductUserResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Price  int    `json:"price"`
	Image  string `json:"image"`
	Qty    int    `json:"qty"`
	// UserID int    `json:"-"`
}

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductUserResponse) TableName() string {
	return "products"
}