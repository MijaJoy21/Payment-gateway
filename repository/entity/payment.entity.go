package entity

import "time"

type Payment struct {
	Id              int        `gorm:"column:id" json:"id"`
	Amount          float64    `gorm:"column:amount" json:"amount"`
	PaymentMethodId int        `gorm:"column:payment_method_id" json:"payment_method_id"`
	Status          int        `gorm:"column:status" json:"status"`
	CreatedAt       time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Payment) TableName() string {
	return "payments"
}
