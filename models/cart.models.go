package models

type CreateCart struct {
	Quantity  int `json:"quantity"`
	UserId    int `json:"user_id"`
	ProductId int `json:"product_id"`
}

type RequestCart struct {
	Quantity int `json:"quantity"`
}

type ResponseListCart struct {
	ProductId       int     `json:"product_id"`
	ProductName     string  `json:"product_name"`
	ProductQuantity int     `json:"product_quantity"`
	ProductImage    string  `json:"product_image"`
	ProductPrice    float64 `json:"product_price"`
}
