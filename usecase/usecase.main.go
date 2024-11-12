package usecase

import (
	"payment-gateway/models"
	"payment-gateway/repository"

	"github.com/gin-gonic/gin"
)

type (
	usecase struct {
		Repository repository.Repository
	}

	Usecase interface {
		GetHealthCheck(ctx *gin.Context) models.Response
	}
)

func InitUsecase(repository repository.Repository) Usecase {
	return &usecase{
		Repository: repository,
	}

}
