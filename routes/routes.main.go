package routes

import (
	"fmt"
	"payment-gateway/controllers"
	"payment-gateway/helpers"

	"github.com/gin-gonic/gin"
)

type (
	Router struct {
		controllers controllers.Controllers
		gin         *gin.Engine
	}

	RouterInterface interface {
		StartGinServer() error
	}
)

func InitRoutes(ctrl controllers.Controllers) RouterInterface {
	return &Router{
		controllers: ctrl,
		gin:         gin.New(),
	}
}

func (r *Router) StartGinServer() error {
	fmt.Println("Start Server")
	//prefix api
	api := r.gin.Group("/api")
	api.GET("/hc", r.controllers.GetHealthCheck)

	if err := helpers.StartGinServer(r.gin); err != nil {
		fmt.Println("Error Start Server", err)
	}

	return nil
}
