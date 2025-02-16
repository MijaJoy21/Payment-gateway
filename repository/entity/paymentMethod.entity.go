package entity

import "time"

type PaymentMethod struct {
	Id        int        `gorm:"column:id" json:"id"`
	Name      string     `gorm:"column:name" json:"name"`
	Status    int        `gorm:"column:status" json:"status"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (PaymentMethod) TableName() string {
	return "payments_methods"
}
