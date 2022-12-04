package profiledto

import "waysbucks/models"

type CreateProfile struct {
	Address    string             `json:"address" form:"address" validate:"required"`
	UserID     int                `json:"user_id"`
	User       models.UserProfile `json:"user"`
}

type UpdateProfile struct {
	Image      string `json:"image" form:"image"`
	UserID     int    `json:"user_id"`
}

type ProfileResponse struct {
	Image      string             `json:"image" form:"image"`
	UserID     int                `json:"user_id"`
	User       models.UserProfile `json:"user"`
}
