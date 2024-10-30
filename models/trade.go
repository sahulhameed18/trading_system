package models

type Trade struct {
	BuyOrderId  string  `json:"buyOrderId"`
	SellOrderId string  `json:"sellOrderId"`
	Symbol      string  `json:"symbol"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}
