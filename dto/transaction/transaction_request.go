package transactiondto

type CreateTransaction struct {
	ID     int64  `json:"id"`
	UserID int    `json:"user_id" form:"user_id"`
	Status string `json:"status"`
}

type UpdateTransaction struct {
	UserID int    `json:"user_id" form:"user_id"`
	Status string `json:"status"`
	Total  int    `json:"total"`
}
