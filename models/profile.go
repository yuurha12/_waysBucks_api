package models

type Profile struct {
	ID     int                  `json:"id" gorm:"primary_key:auto_increment"`
	Profile []UsersProfileResponse `json:"profile"`
}

// for association relation with another table (user)
type ProfileResponse struct {
	Profile int `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
