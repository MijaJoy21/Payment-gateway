package usecase

import (
	"net/http"
	"payment-gateway/models"

	"github.com/gin-gonic/gin"
)

func (uc *usecase) GetHealthCheck(ctx *gin.Context) models.Response {
	res := models.Response{}

	data, err := uc.Repository.GetUsers(ctx)

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
