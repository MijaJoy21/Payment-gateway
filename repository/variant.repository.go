package repository

import (
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (d *repository) CreateVariant(ctx *gin.Context, data []entity.Variant) error {
	query := d.DB.Model(&data)
	query.Create(&data)
	return query.Error
}
