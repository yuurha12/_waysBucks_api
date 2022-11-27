package models

type Product struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Title string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty   int    `json:"-" form:"qty"`
}

type ProductResponse struct {
	Title string `json:"title"`
	Price int    `json:"price"`
	Image string `json:"image"`
	Qty   int    `json:"qty"`
}

// type ProductUserResponse struct {
// 	ID     int    `json:"id"`
// 	Title  string `json:"title"`
// 	Price  int    `json:"price"`
// 	Image  string `json:"image"`
// 	Qty    int    `json:"qty"`
// 	// UserID int    `json:"-"`
// }

func (ProductResponse) TableName() string {
	return "products"
}
