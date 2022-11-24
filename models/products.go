package models

type Product struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title     string               `json:"title" form:"name" gorm:"type: varchar(255)"`
	Price     int                  `json:"price" form:"price" gorm:"type: int"`
	Image     string               `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty       int                  `json:"qty" form:"qty"`
	BuyerID   int                  `json:"buyer_id" form:"user_id"`
	Buyer     UsersProfileResponse `json:"buyer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ToppingID int                  `json:"topping_id" form:"topping_id"`
	Topping   ToppingResponse      `json:"topping" form:"topping"`
}

// func (ProductResponse) TableName() string {
// 	return "products"
// }

type ProductResponse struct {
	ID        int                  `json:"id"`
	Title     string               `json:"title"`
	Price     int                  `json:"price"`
	Image     string               `json:"image"`
	Qty       int                  `json:"qty"`
	BuyerID   int                  `json:"-"`
	Buyer     UsersProfileResponse `json:"buyer"`
	ToppingID int                  `json:"-"`
	Topping   ToppingResponse      `json:"topping"`
}

type ProductUserResponse struct {
	ID        int                  `json:"id"`
	Title     string               `json:"title"`
	Price     int                  `json:"price"`
	Qty       int                  `json:"qty"`
	BuyerID   int                  `json:"-"`
	Buyer     UsersProfileResponse `json:"buyer"`
	ToppingID int                  `json:"-"`
	Topping   ToppingResponse      `json:"topping"`
}

type ProductTotalResponse struct {
	ID    int `json:"id"`
	Price int `json:"price"`
}

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductUserResponse) TableName() string {
	return "products"
}

func (ProductTotalResponse) TableName() string {
	return "products"
}
