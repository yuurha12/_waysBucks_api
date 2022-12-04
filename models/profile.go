package models

import "time"

type Profile struct {
	ID         int         `json:"id"`
	Image      string      `json:"image" gorm:"type: varchar(255)"`
	UserID     int         `json:"user_id"`
	User       UserProfile `json:"user"`
	CreatedAt  time.Time   `json:"-"`
	UpdatedAt  time.Time   `json:"-"`
}

type ProfileResponse struct {
	ID         int    `json:"id"`
	Image      string `json:"image" gorm:"type: varchar(255)"`
	UserID     int    `json:"user_id"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
