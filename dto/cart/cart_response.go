package cartdto

type CartResponse struct {
	ID       int `json:"id"`
	QTY      int `json:"qty"`
	SubTotal int `json:"subtotal"`
}
