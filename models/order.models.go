package models

import "time"

type ReqCreateOrder struct {
	CouponId    *int             `json:"coupon_id,omitempty"`
	OrderDetail []ReqOrderDetail `json:"order_detail" validate:"required"`
}

type ReqOrderDetail struct {
	ProductId int `json:"product_id" validate:"required"`
	VariantId int `json:"variant_id"`
	Quantity  int `json:"quantity" validate:"required"`
}

type GetAllOrderParams struct {
	Search string `json:"search" form:"search"`
	Email  string `json:"email" form:"email"`
	Status *int   `json:"status" form:"status"`
	Page   int    `json:"page" form:"page"`
	Limit  int    `json:"limit" form:"limit"`
}

type GetAllOrderRes struct {
	Id        int    `json:"id"`
	InvoiceID string `json:"invoice_id"`
	Status    string `json:"status"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}

type UpdateOrderStatusByIdReq struct {
	Status int `json:"status"`
}

type GetAllHistoryOrderParams struct {
	Search string `json:"search" form:"search"`
	Page   int    `json:"page" form:"page"`
	Limit  int    `json:"limit" form:"limit"`
}

type GetAllHistoryOrderRes struct {
	Id           int       `json:"id"`
	InvoiceId    string    `json:"invoice_id"`
	ProductName  string    `json:"product_name"`
	ProductImage string    `json:"product_image"`
	Status       string    `json:"status"`
	Quantity     int       `json:"quantity"`
	Price        float64   `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
}
