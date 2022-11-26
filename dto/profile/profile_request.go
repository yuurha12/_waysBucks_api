package profiledto

import "waysbucks_BE/models"

type ProfileResponse struct {
  ID      int                         `json:"id" gorm:"primary_key:auto_increment"`
  User    models.UsersProfileResponse `json:"user"`
}