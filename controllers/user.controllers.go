package controllers

import (
	"log"
	"net/http"
	"payment-gateway/helpers"
	"payment-gateway/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *controllers) GetAllUsers(ctx *gin.Context) {
	res := c.Usecase.GetAllUsers(ctx)

	ctx.JSON(res.Code, res)
}

func (c *controllers) RegistrationUser(ctx *gin.Context) {
	log.Println("<<Controllers Registration User>>")
	var res models.Response
	payload := models.ReqRegistrationUser{}

	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error Bind JSON", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.RegistrationUser(ctx, payload)
	log.Println("Response Registration User", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) GetUserById(ctx *gin.Context) {
	log.Println("<<<Controllers Get User By Id>>>>")
	res := c.Usecase.GetUserById(ctx)
	log.Println("Response Registration User", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) LoginUser(ctx *gin.Context) {
	log.Println("<<<Login User>>>")
	var res models.Response
	payload := models.RegLogin{}

	if err := ctx.BindJSON(&payload); err != nil {
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.LoginUser(ctx, payload)
	log.Println("Response Login User", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) RegistrationAdmin(ctx *gin.Context) {
	log.Println("<<< Controllers Registration Admin >>>")
	var res models.Response
	payload := models.ReqRegistrationUser{}

	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error Bind JSON", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.RegistrationAdmin(ctx, payload)
	log.Println("Response Registration Admin", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) UpdateUser(ctx *gin.Context) {
	log.Println("<<< Controllers update user >>>")
	var res models.Response
	payload := models.UpdateUser{}
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error Bind JSON", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	if err := helpers.Validator(payload); err != nil {
		log.Println("Error validate", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.UpdateUser(ctx, id, payload)

	ctx.JSON(res.Code, res)
}
