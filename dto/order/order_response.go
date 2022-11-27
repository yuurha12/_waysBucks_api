package orderdto

type OrderResponse struct {
	ID        int `json:"id"`
	BuyerID   int `json:"buyer_id"`
	Price     int `json:"price"`
	Qty       int `json:"qty"`
	ProductID int `json:"product_id"`
	TopingID  int `json:"toping_id"`
}
