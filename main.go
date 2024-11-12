package main

import (
	"log"
	"payment-gateway/controllers"
	"payment-gateway/repository"
	"payment-gateway/routes"
	"payment-gateway/usecase"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Load env file")
		panic(500)
	}

	repo := repository.InitRepository()
	uc := usecase.InitUsecase(repo)

	controllers.InitControllers(uc)
	ctrl := controllers.InitControllers(uc)
	router := routes.InitRoutes(ctrl)

	router.StartGinServer()
}
