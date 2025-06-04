package repository

import (
	"payment-gateway/models"
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (d *repository) CreateCoupon(ctx *gin.Context, data entity.Coupon) error {
	query := d.DB.Model(&data)
	query.Create(&data)

	return query.Error
}

func (d *repository) GetListCoupon(ctx *gin.Context, params models.GetListCouponParams) ([]entity.Coupon, int64, error) {
	var data []entity.Coupon
	var total int64
	query := d.DB.Model(&data)
	if params.Status != nil {
		query = query.Where("status = ?", params.Status)
	}

	if params.Search != "" {
		query = query.Where("name like ?", "%"+params.Search+"%")
	}

	query.Count(&total)

	offset := (params.Page - 1) * params.Limit
	query = query.Limit(params.Limit).Offset(offset)
	query.Find(&data)

	return data, total, query.Error
}

func (d *repository) UpdateStatusCouponById(ctx *gin.Context, id int, status *int) error {
	var data entity.Coupon
	query := d.DB.Model(&data)
	query = query.Where("id = ?", id)
	query.Updates(map[string]interface{}{
		"status": status,
	})
	return query.Error
}
