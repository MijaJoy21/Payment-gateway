package controllers

import (
	"log"
	"net/http"
	"payment-gateway/models"

	"github.com/gin-gonic/gin"
)

func (c *controllers) CreateProduct(ctx *gin.Context) {
	log.Println("<<Controllers Create Product>>")
	var res models.Response
	payload := models.CreateProduct{}

	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error Bind JSON", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.CreateProduct(ctx, payload)
	log.Println("Response Create Product", res)

	ctx.JSON(res.Code, res)
}
