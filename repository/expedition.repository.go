package repository

import (
	"payment-gateway/models"
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (db *repository) CreateExpedition(ctx *gin.Context, data entity.Expedition) error {
	query := db.DB.Model(&data)
	query.Save(&data)

	return query.Error
}

func (db *repository) GetExpedition(ctx *gin.Context, params models.ParamsGetExpeditions) ([]entity.Expedition, error) {
	var data []entity.Expedition
	var total int64
	query := db.DB.Model(&data)

	if params.Search != "" {
		query = query.Where("name like ?", "%"+params.Search+"%")
	}

	query.Count(&total)

	offset := (params.Page - 1) * params.Limit
	query = query.Limit(params.Limit).Offset(offset)
	query.Find(&data)

	return data, query.Error
}

func (db *repository) GetExpeditionById(ctx *gin.Context, id int) (entity.Expedition, error) {
	var data entity.Expedition
	query := db.DB.Model(&data)
	query = query.Where("id = ?", id)
	query.First(&data)

	return data, query.Error
}

func (db *repository) PutExpedition(ctx *gin.Context, id int, updatedData entity.Expedition) error {
	var existingData entity.Expedition

	query := db.DB.Model(&existingData)
	if err := query.Where("id = ?", id).First(&existingData).Error; err != nil {
		return err
	}

	query = db.DB.Model(&existingData)
	query.Where("id = ?", id).Updates(updatedData)

	return query.Error
}
