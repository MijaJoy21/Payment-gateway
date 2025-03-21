package entity

import "time"

type Cart struct {
	Id        int        `gorm:"column:id" json:"id"`
	ProductId int        `gorm:"column:product_id" json:"product_id"`
	Product   Product    `gorm:"foreignKey:product_id" json:"product"`
	UserId    int        `gorm:"column:user_id" json:"user_id"`
	User      User       `gorm:"foreignKey:user_id" json:"user"`
	Quantity  int        `gorm:"column:quantity" json:"quantity"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}
