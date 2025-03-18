package usecase

import (
	"log"
	"net/http"
	"payment-gateway/models"
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (uc *usecase) CreateExpedition(ctx *gin.Context, payload models.RequestExpedition) models.Response {
	res := models.Response{}
	data := entity.Expedition{
		Name:   payload.Name,
		Price:  payload.Price,
		Weight: payload.Weight,
	}

	if err := uc.Repository.CreateExpedition(ctx, data); err != nil {
		log.Println("Error Create Category", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unprocessable Entity"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"

	return res
}

func (uc *usecase) GetAllExpedition(ctx *gin.Context, params models.ParamsGetExpeditions) models.Response {
	res := models.Response{}

	data, err := uc.Repository.GetExpedition(ctx, params)

	if err != nil {
		log.Println("Error Get all expedition", err)
		res.Code = http.StatusInternalServerError
		res.Message = "Server Error"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = data

	return res
}

func (uc *usecase) GetExpeditionById(ctx *gin.Context, id int) models.Response {
	res := models.Response{}

	data, err := uc.Repository.GetExpeditionById(ctx, id)

	if err != nil {
		log.Println("Error Get Expedition By ID ", err)
		res.Code = http.StatusNotFound
		res.Message = "Data not found"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = data

	return res

}

func (uc *usecase) PutExpedition(ctx *gin.Context, id int, payload models.RequestExpedition) models.Response {
	res := models.Response{}

	existingData, err := uc.Repository.GetExpeditionById(ctx, id)
	if err != nil {
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"
		return res
	}

	updatedData := entity.Expedition{
		Id:     existingData.Id,
		Name:   payload.Name,
		Price:  payload.Price,
		Weight: payload.Weight,
	}

	if err := uc.Repository.PutExpedition(ctx, id, updatedData); err != nil {
		log.Println("Error updating Category :", err)
		res.Code = http.StatusInternalServerError
		res.Message = "Failed to Update Expedition"
		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	return res
}
