package controllers

import (
	"payment-gateway/usecase"

	"github.com/gin-gonic/gin"
)

type (
	controllers struct {
		Usecase usecase.Usecase
	}

	Controllers interface {
		GetHealthCheck(ctx *gin.Context)
	}
)

func InitControllers(usecase usecase.Usecase) Controllers {
	return &controllers{
		Usecase: usecase,
	}
}
