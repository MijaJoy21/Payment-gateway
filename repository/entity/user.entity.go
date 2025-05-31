package entity

import "time"

type User struct {
	Id        int        `gorm:"column:id" json:"id"`
	Name      string     `gorm:"column:name" json:"name"`
	Email     string     `gorm:"column:email" json:"email"`
	Address   string     `gorm:"column:address" json:"address"`
	Password  string     `gorm:"column:password" json:"password"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Type      int        `gorm:"column:type" json:"type"`
}

func (User) TableName() string {
	return "users"
}
