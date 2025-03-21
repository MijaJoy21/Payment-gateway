package models

type ReqCreateOrder struct {
	CouponId  *int             `json:"omitempty"`
	ProductId []ReqOrderDetail `json:"order_detail"`
}

type ReqOrderDetail struct {
	ProductId int `json:"product_id" validate:"required"`
	VariantId int `json:"variant_id"`
	Quantity  int `json:"quantity" validate:"required"`
}
