package models

type Profile struct {
	ID     int                  `json:"id" gorm:"primary_key:auto_increment"`
	UserID int                  `json:"user_id"`
	User   UsersProfileResponse `json:"user"`
}

// for association relation with another table (user)
type ProfileResponse struct {
	UserID int `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
