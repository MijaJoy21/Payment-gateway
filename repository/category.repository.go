package repository

import (
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (db *repository) CreateCategory(ctx *gin.Context, data entity.Category) error {
	query := db.DB.Model(&data)
	query.Save(&data)

	return query.Error
}

func (db *repository) GetCategory(ctx *gin.Context) ([]entity.Category, error) {
	var data []entity.Category

	query := db.DB.Model(&data)
	query.Find(&data)

	return data, query.Error
}

func (db *repository) GetCategoryById(ctx *gin.Context, id int) (entity.Category, error) {
	var data entity.Category
	query := db.DB.Model(&data)
	query = query.Where("id = ?", id)
	query.First(&data)

	return data, query.Error
}

func (db *repository) PutCategory(ctx *gin.Context, id int, updatedData entity.Category) error {
	var existingData entity.Category

	query := db.DB.Model(&existingData)
	if err := query.Where("id = ?", id).First(&existingData).Error; err != nil {
		return err
	}

	query = db.DB.Model(&existingData)
	query.Where("id = ?", id).Updates(updatedData)

	return query.Error
}
