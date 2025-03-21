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

func (uc *usecase) GetAllProduct(ctx *gin.Context) models.Response {
	res := models.Response{}

	data, err := uc.Repository.GetProduct(ctx)

	if err != nil {
		log.Println("Error Get All Product", err)
		res.Code = http.StatusInternalServerError
		res.Message = "Server Error"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = data

	return res
}

func (uc *usecase) GetProductById(ctx *gin.Context, id int) models.Response {
	res := models.Response{}

	data, err := uc.Repository.GetProductById(ctx, id)

	if err != nil {
		log.Println("Error Data Not Found", err)
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = data

	return res
}

func (uc *usecase) PutProduct(ctx *gin.Context, id int, payload models.RequestProduct) models.Response {
	res := models.Response{}

	_, err := uc.Repository.GetProductById(ctx, id)
	if err != nil {
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"
		return res
	}

	updatedData := entity.Product{
		Name:        payload.Name,
		Categoryid:  payload.CategoryId,
		Description: payload.Description,
		Price:       payload.Price,
		Status:      payload.Status,
		IsPreorder:  payload.IsPreOrder,
		Weight:      payload.Weight,
	}

	if err := uc.Repository.PutProduct(ctx, id, updatedData); err != nil {
		log.Println("Error updating Category :", err)
		res.Code = http.StatusInternalServerError
		res.Message = "Failed to update Category"
		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	return res
}
