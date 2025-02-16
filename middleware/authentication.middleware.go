package middleware

import (
	"log"
	"net/http"
	"os"
	"payment-gateway/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication(level string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := strings.Split(ctx.Request.Header.Get("Authorization"), " ")
		newData := &models.ClaimsJwt{}

		if tokenString[0] != "Bearer" {
			log.Println("invalid token", tokenString)
			res := models.Response{
				Code:    http.StatusUnauthorized,
				Message: "Invalid Token",
			}
			ctx.JSON(res.Code, res)
			ctx.Abort()
			return
		}

		_, err := jwt.ParseWithClaims(tokenString[1], newData, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Unexpected Signing Method"})
				ctx.Abort()
			}
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil {
			log.Println("Error Token :", err)
			res := models.Response{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			}
			ctx.JSON(res.Code, res)
			ctx.Abort()
			return
		}

		if level == "Admin" {
			if newData.Type != 0 {
				log.Println("Admin Only")
				res := models.Response{
					Code:    http.StatusForbidden,
					Message: "You don't have access to use this feature",
				}
				ctx.JSON(res.Code, res)
				ctx.Abort()
			}
		}

		ctx.Set("user", newData)
		ctx.Next()
	}
}
