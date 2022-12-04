package authdto

type RegisterRequest struct {
	Fullname string `json:"fullname" gorm:"type: varchar(255) required"`
	Email    string `json:"email" gorm:"type: varchar(255) required"`
	Password string `json:"password" gorm:"type: varchar(255) required"`
}

type LoginRequest struct {
	Email    string `json:"email" gorm:"type: varchar(255) required"`
	Password string `json:"password" gorm:"type: varchar(255) required"`
}

type LoginResponse struct {
	Fullname string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Token    string `json:"token" gorm:"type: varchar(255)"`
	Role     string `gorm:"type: varchar(50)"  json:"role"`
}

type CheckAuthResponse struct {
	Id       int    `gorm:"type: int" json:"id"`
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Role     string `gorm:"type: varchar(50)"  json:"role"`
}
