package model

type Order struct {
	ID       int    `json:"id"`
	BuyerId  string `json:"buyer_id"`
	OrderId  string `json:"order_id"`
	Title    string `json:"title"`
	TotalFee int    `json:"total_fee"`
}
