package controllers

import (
	"payment-gateway/models"

	"github.com/gin-gonic/gin"
)

func (c *controllers) GetHealthCheck(ctx *gin.Context) {
	var res models.Response

	res = c.Usecase.GetHealthCheck(ctx)

	ctx.JSON(res.Code, res)
}
