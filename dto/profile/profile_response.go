package profiledto

import "ways-bucks-api/models"

type ProfileResponse struct {
	Address    string             `json:"address" form:"address"`
	Phone      string             `json:"phone" form:"phone"`
	Image      string             `json:"image" form:"image"`
	City       string             `json:"city" form:"city"`
	PostalCode int                `json:"postal_code" form:"postal_code"`
	UserID     int                `json:"user_id"`
	User       models.UserProfile `json:"user"`
}
