package usecase

import (
	"log"
	"net/http"
	"payment-gateway/models"
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (uc *usecase) CreatePayment(ctx *gin.Context, payload models.ReqPayment) models.Response {
	res := models.Response{}
	data := entity.Payment{
		Amount:          payload.Amount,
		PaymentMethodId: payload.PaymentMethodId,
		Status:          payload.Status,
	}

	if err := uc.Repository.CreatePayment(ctx, data); err != nil {
		log.Println("Error Create User", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unnprocessable Entity"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"

	return res
}
