package models

type User struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"-" gorm:"type: varchar(255)"`
	Image    string `json:"image" grom:"type: varchar(255)"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Image    string `json:"image"`
}

type UsersProductsResponse struct {
	ID int `json:"id"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
