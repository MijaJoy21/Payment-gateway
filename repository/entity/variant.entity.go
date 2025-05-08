package entity

import "time"

type Variant struct {
	Id        int        `gorm:"column:id" json:"id"`
	ProductId int        `gorm:"column:product_id" json:"product_id"`
	Name      string     `gorm:"column:name" json:"name"`
	Price     int        `gorm:"column:price" json:"price"`
	Weight    int        `gorm:"column:weight" json:"weight"`
	Quantity  int        `gorm:"column:quantity" json:"quantity"`
	Status    int        `gorm:"column:status" json:"status"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Variant) TableName() string {
	return "variant"
}
