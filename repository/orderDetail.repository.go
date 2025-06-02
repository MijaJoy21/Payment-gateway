package repository

import (
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (d *repository) CreateOrderDetail(ctx *gin.Context, data []entity.OrderDetail) error {
	query := d.DB.Model(&data)
	query.Create(data)

	return query.Error
}
