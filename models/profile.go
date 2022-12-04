package models

import "time"

type Profile struct {
	ID         int         `json:"id"`
	Phone      int         `json:"phone" gorm:"type: int(20)"`
	PostalCode int         `json:"postalcode" gorm:"type: int(5)"`
	Address    string      `json:"address" gorm:"type: varchar(255)"`
	Image      string      `json:"image" gorm:"type: varchar(255)"`
	UserID     int         `json:"user_id"`
	User       UserProfile `json:"user"`
	CreatedAt  time.Time   `json:"-"`
	UpdatedAt  time.Time   `json:"-"`
}

type ProfileResponse struct {
	ID         int    `json:"id"`
	Phone      int    `json:"phone"`
	PostalCode int    `json:"postalcode"`
	Address    string `json:"address"`
	Image      string `json:"image"`
	UserID     int    `json:"user_id"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
