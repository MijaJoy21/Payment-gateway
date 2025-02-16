package usecase

import (
	"log"
	"net/http"
	"payment-gateway/models"
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (uc *usecase) CreatePaymentMethod(ctx *gin.Context, payload models.ReqPaymentMethod) models.Response {
	res := models.Response{}
	data := entity.PaymentMethod{
		Name:   payload.Name,
		Status: payload.Status,
	}

	if err := uc.Repository.CreatePaymentMethod(ctx, data); err != nil {
		log.Println("Error Create User", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unprocessable Entity"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"

	return res
}

func (uc *usecase) GetAllPaymentMethods(ctx *gin.Context) models.Response {
	res := models.Response{}

	data, err := uc.Repository.GetPaymentMethods(ctx)

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

func (uc *usecase) GetPaymentMethodById(ctx *gin.Context, id int) models.Response {
	res := models.Response{}

	data, err := uc.Repository.GetPaymentMethodById(ctx, id)

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

func (uc *usecase) PutPaymentMethod(ctx *gin.Context, id int, payload models.ReqPaymentMethod) models.Response {
	res := models.Response{}

	existingData, err := uc.Repository.GetPaymentMethodById(ctx, id)
	if err != nil {
		res.Code = http.StatusNotFound
		res.Message = "Data not Found"
		return res
	}

	updatedData := entity.PaymentMethod{
		Id:     existingData.Id,
		Name:   payload.Name,
		Status: payload.Status,
	}

	if err := uc.Repository.PutPaymentMethod(ctx, id, updatedData); err != nil {
		log.Println("Error updating Payment Method:", err)
		res.Code = http.StatusInternalServerError
		res.Message = "Failed to update Payment Method"
		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	return res

}
