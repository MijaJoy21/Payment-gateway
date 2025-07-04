package entity

import "time"

type Product struct {
	Id           int        `gorm:"column:id" json:"id"`
	Categoryid   int        `gorm:"column:category_id" json:"category_id"`
	Category     Category   `gorm:"references:category_id; foreignKey:id"  json:"category"`
	Name         string     `gorm:"column:name" json:"name"`
	Description  string     `gorm:"column:description" json:"description"`
	Image        string     `gorm:"column:image" json:"image"`
	Price        float64    `gorm:"column:price" json:"price"`
	Quantity     int        `gorm:"column:quantity" json:"quantity"`
	Status       *int       `gorm:"column:status" json:"status"`
	IsPreorder   int        `gorm:"column:is_preorder" json:"is_preorder"`
	PreOrderDate *time.Time `gorm:"column:pre_orderdate" json:"pre_orderdate"`
	Weight       int        `gorm:"column:weight" json:"weight"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Product) TableName() string {
	return "product"
}
