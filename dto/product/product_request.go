package productdto

type ProductRequest struct {
	ID    string `json:"id" form:"id" gorm:"type: int" validate:"required"`
	Title string `json:"title" form:"title" gorm:"type: varchar(255)" validate:"required"`
	Price int    `json:"price" form:"price" gorm:"type: int" validate:"required"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty   int    `json:"qty" form:"qty" gorm:"type: int" validate:"required"`
}
