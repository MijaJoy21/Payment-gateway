package entity

type OrderDetail struct {
	Id        int     `gorm:"column:id" json:"id"`
	OrderId   int     `gorm:"column:order_id" json:"order_id"`
	ProductId int     `gorm:"column:product_id" json:"product_id"`
	Product   Product `gorm:"references:product_id; foreignKey:id" json:"product"`
	Quantity  int     `gorm:"column:quantity" json:"quantity"`
}

func (OrderDetail) TableName() string {
	return "orders_detail"
}
