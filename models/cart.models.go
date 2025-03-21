package models

type CreateCart struct {
	Quantity  int `json:"quantity"`
	UserId    int `json:"user_id"`
	ProductId int `json:"product_id"`
}

type RequestCart struct {
	Quantity int `json:"quantity"`
}
