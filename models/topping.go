package models

import "time"

type Topping struct {
	ID       int       `json:"id" gorm:"primary_key:auto_increment"`
	Title    string    `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price    int       `json:"price" form:"price" gorm:"type: int"`
	Image    string    `json:"-" form:"image" gorm:"type: varchar(255)"`
	Qty      int       `json:"-" form:"qty"`
	CreateAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
}

type ToppingResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Image string `json:"image"`
	Qty   int    `json:"qty"`
}

type ToppingOrder struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func (ToppingResponse) TableName() string {
	return "topping"
}

func (ToppingOrder) TableName() string {
	return "topping"
}
