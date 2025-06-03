package entity

import "time"

type Coupon struct {
	ID              int        `gorm:"column:id" json:"id"`
	Name            string     `gorm:"column:name" json:"name"`
	Quantity        int        `gorm:"column:quantity" json:"quantity"`
	Code            string     `gorm:"column:code" json:"code"`
	MinimumPurchase float64    `gorm:"column:minimum_purchase" json:"minimum_purchase"`
	MaximumDiscount float64    `gorm:"column:maximum_discount" json:"maximum_discount"`
	Discount        *int       `gorm:"column:discount" json:"discount"`
	Type            int        `gorm:"column:type" json:"type"`
	Status          int        `gorm:"column:status" json:"status"`
	CreatedAt       *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Coupon) TableName() string {
	return "coupons"
}
