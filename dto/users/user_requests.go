package orderdto

type CreateUserRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Image    string `json:"image" form:"image" validate:"required"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Fullname string `json:"fullname" form:"fullname"`
	Image    string `json:"image" form:"image"`
}
