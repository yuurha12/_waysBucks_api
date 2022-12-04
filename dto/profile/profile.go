package profiledto

import "waysbucks/models"

type CreateProfile struct {
	Phone      int                `json:"phone" form:"phone" validate:"required"`
	PostalCode int                `json:"postalcode" form:"postal_code" validate:"required"`
	Address    string             `json:"address" form:"address" validate:"required"`
	Image      string             `json:"image" form:"image" validate:"required"`
	UserID     int                `json:"user_id"`
	User       models.UserProfile `json:"user"`
}

type UpdateProfile struct {
	Phone      int    `json:"phone" form:"phone"`
	PostalCode int    `json:"postalcode" form:"postalcode"`
	Address    string `json:"address" form:"address"`
	Image      string `json:"image" form:"image"`
	UserID     int    `json:"user_id"`
}

type ProfileResponse struct {
	Phone      int                `json:"phone" form:"phone"`
	PostalCode int                `json:"postalcode" form:"postalcode"`
	Address    string             `json:"address" form:"address"`
	Image      string             `json:"image" form:"image"`
	UserID     int                `json:"user_id"`
	User       models.UserProfile `json:"user"`
}
