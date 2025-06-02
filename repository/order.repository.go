package repository

import (
	"payment-gateway/models"
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (d *repository) CreateOrder(ctx *gin.Context, data *entity.Order) error {
	query := d.DB.Model(data)
	query.Create(data)
	query.First(data)

	return query.Error
}

func (d *repository) GetAllOrder(ctx *gin.Context, params models.GetAllOrderParams) ([]entity.Order, int64, error) {
	var data []entity.Order
	var total int64

	query := d.DB.Model(&data)

	if params.Status != nil {
		query = query.Where("status_order = ?", params.Status)
	}

	if params.Search != "" {
		query = query.Where("invoice_id = ?", params.Search)
	}

	if params.Email != "" {
		query = query.Where("User.email like ?", "%"+params.Email+"%")
	}

	query = query.Preload("OrderDetail")
	query = query.Joins("User")
	query.Count(&total)

	offset := (params.Page - 1) * params.Limit
	query = query.Limit(params.Limit).Offset(offset)
	query.Find(&data)

	return data, total, query.Error
}

func (d *repository) UpdateOrderStatusById(ctx *gin.Context, status int, id int) error {
	var data entity.Order
	query := d.DB.Model(&data)
	query = query.Where("id = ?", id)
	query.Updates(map[string]interface{}{
		"status_order": status,
	})

	return query.Error
}
