package usecase

import (
	"log"
	"net/http"
	"payment-gateway/models"
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (uc *usecase) CreateProduct(ctx *gin.Context, payload models.CreateProduct) models.Response {
	res := models.Response{}
	data := entity.Product{
		Name:        payload.Name,
		Categoryid:  payload.CategoryId,
		Description: payload.Description,
		Price:       payload.Price,
		Status:      payload.Status,
		IsPreorder:  payload.IsPreOrder,
		Weight:      payload.Weight,
	}

	if err := uc.Repository.CreateProduct(ctx, data); err != nil {
		log.Println("Error Create Product", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unprocessable Entity"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"

	return res
}
