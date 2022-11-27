package models

type Order struct {
	ID        int           `json:"id" gorm:"primary_key: auto_increment"`
	BuyerID   int           `json:"-"`
	Buyer     UsersResponse `json:"buyer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductID int           `json:"-"`
	Product   Product       `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Topping   []Topping     `json:"topping" gorm:"many2many:topping_order;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Price     int           `json:"price" gorm:"type: int"`
	Qty       int           `json:"qty" gorm:"type: int"`
}

type OrderResponse struct {
	ID        int `json:"id"`
	BuyerID   int `json:"buyer_id"`
	Qty       int `json:"qty"`
	Price     int `json:"price"`
	ProductID int `json:"product_id"`
	TopingID  int `json:"toping_id"`
}

func (OrderResponse) TableName() string {
	return "orders"
}
