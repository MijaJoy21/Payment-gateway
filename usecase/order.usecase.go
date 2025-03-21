package usecase

import (
	"net/http"
	"payment-gateway/models"

	"github.com/gin-gonic/gin"
)

func (u *usecase) CreateOrder(ctx *gin.Context, payload models.ReqCreateOrder) models.Response {
	res := models.Response{}

	res.Code = http.StatusOK
	res.Message = "Success"
	return res
}
