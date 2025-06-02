package controllers

import (
	"github.com/gin-gonic/gin"
)

func (c *controllers) GetHealthCheck(ctx *gin.Context) {
	res := c.Usecase.GetHealthCheck(ctx)

	ctx.JSON(res.Code, res)
}
