package models

type Product struct {
	ID     int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title  string               `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price  int                  `json:"price" form:"price" gorm:"type: int"`
	Image  string               `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty    int                  `json:"-" form:"qty"`
	UserID int                  `json:"-" form:"user_id"`
	User   UsersProfileResponse `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ProductResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Image string `json:"image"`
	// Qty   int    `json:"qty"`
	// UserID int                  `json:"-"`
	// User   UsersProfileResponse `json:"user"`
}

type ProductUserResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Price  int    `json:"price"`
	Image  string `json:"image"`
	Qty    int    `json:"qty"`
	UserID int    `json:"-"`
}

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductUserResponse) TableName() string {
	return "products"
}
