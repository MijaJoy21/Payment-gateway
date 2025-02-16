package repository

import (
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (db *repository) CreateEtalase(ctx *gin.Context, data entity.Etalase) error {
	query := db.DB.Model(&data)
	query.Save(&data)

	return query.Error
}

func (db *repository) GetEtalase(ctx *gin.Context) ([]entity.Etalase, error) {
	var data []entity.Etalase

	query := db.DB.Model(&data)
	query.Find(&data)

	return data, query.Error
}

func (db *repository) GetEtalaseById(ctx *gin.Context, id int) (entity.Etalase, error) {
	var data entity.Etalase
	query := db.DB.Model(&data)
	query = query.Where("id = ?", id)
	query.First(&data)

	return data, query.Error

}

func (db *repository) PutEtalase(ctx *gin.Context, id int, updatedData entity.Etalase) error {
	var existingData entity.Etalase

	query := db.DB.Model(&existingData)
	if err := query.Where("id = ?", id).First(&existingData).Error; err != nil {
		return err
	}

	query = db.DB.Model(&existingData)
	query.Where("id = ?", id).Updates(updatedData)

	return query.Error
}
