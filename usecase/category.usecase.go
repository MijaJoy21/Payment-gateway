package usecase

import (
	"log"
	"net/http"
	"payment-gateway/models"
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (uc *usecase) CreateCategory(ctx *gin.Context, payload models.RequestCategory) models.Response {
	res := models.Response{}
	data := entity.Category{
		Name: payload.Name,
	}

	if err := uc.Repository.CreateCategory(ctx, data); err != nil {
		log.Println("Error Create Category", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unprocessable Entity"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"

	return res
}

func (uc *usecase) GetAllCategory(ctx *gin.Context) models.Response {
	res := models.Response{}

	data, err := uc.Repository.GetCategory(ctx)

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = "Server Error"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = data

	return res
}

func (uc *usecase) GetCategoryById(ctx *gin.Context, id int) models.Response {
	res := models.Response{}

	data, err := uc.Repository.GetCategoryById(ctx, id)

	if err != nil {
		res.Code = http.StatusNotFound
		res.Message = "Data not found"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = data

	return res
}

func (uc *usecase) PutCategory(ctx *gin.Context, id int, payload models.RequestCategory) models.Response {
	res := models.Response{}

	existingData, err := uc.Repository.GetCategoryById(ctx, id)
	if err != nil {
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"
		return res
	}

	updatedData := entity.Category{
		Id:   existingData.Id,
		Name: payload.Name,
	}

	if err := uc.Repository.PutCategory(ctx, id, updatedData); err != nil {
		log.Println("Error updating Category :", err)
		res.Code = http.StatusInternalServerError
		res.Message = "Failed to update Category"
		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	return res
}
