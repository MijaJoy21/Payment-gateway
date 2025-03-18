package repository

import (
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (db *repository) CreateProduct(ctx *gin.Context, data entity.Product) error {
	query := db.DB.Model(&data)
	query.Save(&data)

	return query.Error
}
