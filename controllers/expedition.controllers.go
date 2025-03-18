package controllers

import (
	"fmt"
	"log"
	"net/http"
	"payment-gateway/helpers"
	"payment-gateway/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *controllers) CreateExpedition(ctx *gin.Context) {
	log.Println("Controllers Registration Expedition")
	var res models.Response
	payload := models.RequestExpedition{}

	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error Bind JSON", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	err := helpers.Validator(payload)
	if err != nil {
		log.Println("Error", err)
		res.Code = http.StatusBadRequest
		res.Message = "Please FIlled The Required Field"
		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.CreateExpedition(ctx, payload)
	log.Println("Response Registration User", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) GetAllExpedition(ctx *gin.Context) {
	var res models.Response

	params := models.ParamsGetExpeditions{}
	if err := ctx.BindQuery(&params); err != nil {
		fmt.Println("Error get query params ", err)
		res.Code = http.StatusInternalServerError
		res.Message = "Server Error "

		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res = c.Usecase.GetAllExpedition(ctx, params)
	log.Println("Response Get All Expendition")

	ctx.JSON(res.Code, res)
}

func (c *controllers) GetExpeditionById(ctx *gin.Context) {
	log.Println("Conrollers Get Category By ID")
	var res models.Response

	id, _ := strconv.Atoi(ctx.Param("id"))

	res = c.Usecase.GetExpeditionById(ctx, id)
	log.Println("Response Get Detail Category By ID", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) PutExpediton(ctx *gin.Context) {
	log.Println("Controllers Update Expediiton")
	var res models.Response

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println("Invalid ID Parameter", err)
		res.Code = http.StatusBadRequest
		res.Message = "Invalid ID Parameter"
		ctx.JSON(res.Code, res)
		return
	}

	payload := models.RequestExpedition{}
	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("error Binding JSON", err)
		res.Code = http.StatusBadRequest
		res.Message = "Invalid Request Body"
		ctx.JSON(res.Code, res)
		return
	}

	err = helpers.Validator(payload)
	if err != nil {
		log.Println("Error", err)
		res.Code = http.StatusBadRequest
		res.Message = "Please Filled The required field"
		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.PutExpedition(ctx, id, payload)
	log.Println("Response Update Expedition", res)

	ctx.JSON(res.Code, res)
}
