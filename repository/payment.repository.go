package repository

import (
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (db *repository) CreatePayment(ctx *gin.Context, data entity.Payment) error {
	query := db.DB.Model(&data)
	query.Save(&data)

	return query.Error
}
