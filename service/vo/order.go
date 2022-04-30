package vo

type CreateOrderVo struct {
	BuyerId  string `json:"buyer_id" query:"buyer_id"`
	OrderId  string `json:"order_id" query:"order_id"`
	Title    string `json:"title" query:"title"`
	TotalFee int    `json:"total_fee" query:"total_fee"`
}
