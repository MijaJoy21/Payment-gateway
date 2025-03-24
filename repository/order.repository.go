package repository

import (
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (d *repository) CreateOrder(ctx *gin.Context, data *entity.Order) error {
	query := d.DB.Model(data)
	query.Create(data)
	query.First(data)

	return query.Error
}
