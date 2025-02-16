package repository

import (
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (db *repository) CreatePaymentMethod(ctx *gin.Context, data entity.PaymentMethod) error {
	query := db.DB.Model(&data)
	query.Save(&data)

	return query.Error
}

func (db *repository) GetPaymentMethods(ctx *gin.Context) ([]entity.PaymentMethod, error) {
	var data []entity.PaymentMethod

	query := db.DB.Model(&data)
	query.Find(&data)

	return data, query.Error

}

func (db *repository) GetPaymentMethodById(ctx *gin.Context, id int) (entity.PaymentMethod, error) {
	var data entity.PaymentMethod
	query := db.DB.Model(&data)
	query = query.Where("id = ?", id)
	query.First(&data)

	return data, query.Error
}

func (db *repository) PutPaymentMethod(ctx *gin.Context, id int, updatedData entity.PaymentMethod) error {
	var existingData entity.PaymentMethod

	query := db.DB.Model(&existingData)
	if err := query.Where("id = ?", id).First(&existingData).Error; err != nil {
		return err
	}

	query = db.DB.Model(&existingData)
	query.Where("id = ?", id).Updates(updatedData)

	return query.Error
}
