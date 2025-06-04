package models

type CreateCoupon struct {
	Name            string  `json:"name" validate:"required"`
	Quantity        int     `json:"quantity" validate:"required"`
	Code            string  `json:"code" validate:"required"`
	MinimumPurchase float64 `json:"minimum_purchase" validate:"required"`
	MaximumDiscount float64 `json:"maximum_discount"`
	Discount        int     `json:"discount"`
	Type            int     `json:"type" validate:"required"`
}

type GetListCouponParams struct {
	Search string `json:"search" form:"search"`
	Limit  int    `json:"limit" form:"limit"`
	Page   int    `json:"page" form:"page"`
	Status *int   `json:"status" form:"status"`
}

type UpdateCouponStatusReq struct {
	Status *int `json:"status"`
}
