package controllers

import (
	"log"
	"net/http"
	"payment-gateway/models"

	"github.com/gin-gonic/gin"
)

func (c *controllers) CreatePayment(ctx *gin.Context) {
	log.Println("<<Controllers Create Payment>>")
	var res models.Response
	payload := models.ReqPayment{}

	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error Bind Json", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.CreatePayment(ctx, payload)
	log.Println("Response Create Payment", res)

	ctx.JSON(res.Code, res)

}
