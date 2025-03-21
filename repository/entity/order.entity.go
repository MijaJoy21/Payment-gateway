package entity

import "time"

type Order struct {
	Id           int       `gorm:"column:id" json:"id"`
	UserId       int       `gorm:"column:user_id" json:"user_id"`
	ResiId       string    `gorm:"column:resi_id" json:"resi_id"`
	ExpeditionId int       `gorm:"column:expedition_id" json:"expedition_id"`
	CouponId     *int      `gorm:"column:coupon_id" json:"coupon_id"`
	InvoiceId    string    `gorm:"column:invoice_id" json:"invoice_id"`
	StatusOrder  int       `gorm:"column:status_order" json:"status_order"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Order) TableName() string {
	return "orders"
}

func (o *Order) StatusOrderName() string {
	status := map[int]string{
		0: "Confirmed",
		1: "Paid",
		2: "Waiting For Delivery",
		3: "Delivering",
		4: "Delivered",
		5: "Finish",
	}

	return status[o.StatusOrder]
}
