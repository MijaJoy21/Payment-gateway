package entity

import "time"

type Expedition struct {
	Id        int        `gorm:"column:id" json:"id"`
	Name      string     `gorm:"column:name" json:"name"`
	Price     int        `gorm:"column:price" json:"price"`
	Weight    int        `gorm:"column:weight" json:"weight"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Expedition) TableName() string {
	return "expedition"
}
