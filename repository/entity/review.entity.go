package entity

import "time"

type Review struct {
	Id          int       `gorm:"columnn:id" json:"id"`
	UserId      int       `gorm:"column:user_id" json:"user_id"`
	User        User      `gorm:"references:user_id; foreignKey:id" json:"user"`
	ProductId   int       `gorm:"column:product_id" json:"product_id"`
	Product     Product   `gorm:"references:product_id; foreignKey:id" json:"product"`
	Description string    `gorm:"column:description" json:"description"`
	Image       string    `gorm:"column:image" json:"image"`
	IsAnon      int       `gorm:"column:is_anon" json:"is_anon"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
}

func (Review) TableName() string {
	return "review"
}
