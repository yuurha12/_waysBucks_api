package models

type Profile struct {
	ID       int                  `json:"id" gorm:"primary_key:auto_increment"`
	UserID   int                  `json:"user_id"`
	User     UsersProfileResponse `json:"user"`
	Fullname string               `json:"fullname"`
	Email    string               `json:"email"`
	Image    string               `json:"image"`
}

// for association relation with another table (user)
type ProfileResponse struct {
	UserID   int    `json:"-"`
	Fullname string `json:"-"`
	Email    string `json:"-"`
	Image    string `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
