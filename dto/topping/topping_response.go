package toppingdto

type ToppingResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Image string `json:"image"`
	// Qty   int    `json:"qty"`
	// UserID int                  `json:"-"`
	// User   UsersProfileResponse `json:"user"`
}
