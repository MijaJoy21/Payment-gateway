package usecase

import (
	"log"
	"net/http"
	"os"
	"payment-gateway/models"
	"payment-gateway/repository/entity"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func (uc *usecase) GetAllUsers(ctx *gin.Context) models.Response {
	res := models.Response{}

	data, err := uc.Repository.GetUsers(ctx)

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = "Server Error"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = data

	return res

}

func (uc *usecase) RegistrationUser(ctx *gin.Context, payload models.ReqRegistrationUser) models.Response {
	res := models.Response{}

	_, err := uc.Repository.GetUserByEmail(ctx, payload.Email)
	if err == nil {
		log.Println("Email already Registered")
		res.Code = http.StatusBadRequest
		res.Message = "Email already registred"

		return res
	}

	newPassword, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	data := entity.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(newPassword),
	}

	if err := uc.Repository.CreateUser(ctx, data); err != nil {
		log.Println("Error Create User", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unprocessable Entity"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success Create User"

	return res
}

func (uc *usecase) GetUserById(ctx *gin.Context) models.Response {
	res := models.Response{}
	userData := &models.ClaimsJwt{}

	if ctx.Value("user") != nil {
		userData = ctx.Value("user").(*models.ClaimsJwt)
	}

	data, err := uc.Repository.GetUserById(ctx, userData.Id)

	if err != nil {
		res.Code = http.StatusNotFound
		res.Message = "Data not found"

		return res
	}

	var user models.ResUser
	copier.Copy(&user, data)

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = user

	return res
}

func (uc *usecase) LoginUser(ctx *gin.Context, payload models.RegLogin) models.Response {
	res := models.Response{}
	data, err := uc.Repository.GetUserByEmail(ctx, payload.Email)

	if err != nil {
		log.Println("Error Get User by email:", err)
		res.Code = http.StatusNotFound
		res.Message = "Email Not Found"

		return res
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(payload.Password)); err != nil {
		log.Println("Wrong Password:", err)
		res.Code = http.StatusUnauthorized
		res.Message = "Wrong Password"

		return res
	}

	expiredTime := time.Now().Add(time.Hour * 1)
	claims := models.ClaimsJwt{
		Id:   data.Id,
		Name: data.Name,
		Type: data.Type,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		log.Println("Error Signed JWT", err)
		res.Code = http.StatusInternalServerError
		res.Message = "Server Error"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{
		"token": tokenString,
	}

	return res
}

func (uc *usecase) RegistrationAdmin(ctx *gin.Context, payload models.ReqRegistrationUser) models.Response {
	res := models.Response{}

	_, err := uc.Repository.GetUserByEmail(ctx, payload.Email)
	if err == nil {
		log.Println("Email already Registered")
		res.Code = http.StatusBadRequest
		res.Message = "Email Already Registered"

		return res
	}

	newPassword, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	data := entity.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(newPassword),
		Type:     0,
	}

	if err := uc.Repository.CreateUser(ctx, data); err != nil {
		log.Println("Error Create User", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unprocessable Entity"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success Create Admin"

	return res
}
