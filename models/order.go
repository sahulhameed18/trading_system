package models

type Order struct {
	Id       string  `json:"id"`
	UserId   string  `json:"userId"`
	Symbol   string  `json:"symbol"`
	Type     string  `json:"type"`
	Side     string  `json:"side"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type OrderResponse struct {
	Id           string `json:"id"`
	ResponseCode string `json:"responseCode"`
}
