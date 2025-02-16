package usecase

import (
	"log"
	"net/http"
	"payment-gateway/models"
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (uc *usecase) CreateEtalase(ctx *gin.Context, payload models.RequestEtalase) models.Response {
	res := models.Response{}
	data := entity.Etalase{
		Name: payload.Name,
	}

	if err := uc.Repository.CreateEtalase(ctx, data); err != nil {
		log.Println("Error Create Etalase", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unprocessable Entity"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"

	return res
}

func (uc *usecase) GetAllEtalase(ctx *gin.Context) models.Response {
	res := models.Response{}

	data, err := uc.Repository.GetEtalase(ctx)

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

func (uc *usecase) GetEtalaseById(ctx *gin.Context, id int) models.Response {
	res := models.Response{}

	data, err := uc.Repository.GetEtalaseById(ctx, id)

	if err != nil {
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"
		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = data

	return res
}

func (uc *usecase) PutEtalase(ctx *gin.Context, id int, payload models.RequestEtalase) models.Response {
	res := models.Response{}

	existingData, err := uc.Repository.GetEtalaseById(ctx, id)
	if err != nil {
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"
		return res
	}

	updateData := entity.Etalase{
		Id:   existingData.Id,
		Name: payload.Name,
	}

	if err := uc.Repository.PutEtalase(ctx, id, updateData); err != nil {
		log.Println("Error updating Category :", err)
		res.Code = http.StatusInternalServerError
		res.Message = "Failed to update Category"
		return res
	}

	res.Code = http.StatusOK
	res.Message = "Succcess"
	return res
}
