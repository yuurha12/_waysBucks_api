package models

import "time"

type User struct {
	ID         int             `json:"id"`
	Fullname   string          `json:"fullname" gorm:"type: varchar(255)"`
	Email      string          `json:"email" gorm:"type: varchar(255)"`
	Password   string          `json:"password" gorm:"type: varchar(255)"`
	Role       string          `json:"role" gorm:"type: varchar(255)"`
	Phone      int             `json:"phone" gorm:"type: int(20)"`
	PostalCode int             `json:"postalcode" gorm:"type: int(6)"`
	Address    string          `json:"address" gorm:"type: varchar(255)"`
	Profile    ProfileResponse `json:"profile"`
	CreatedAt  time.Time       `json:"-"`
	UpdatedAt  time.Time       `json:"-"`
}

type UserProfile struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

func (UserProfile) TableName() string {
	return "users"
}
