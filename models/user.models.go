package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ReqRegistrationUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ClaimsJwt struct {
	jwt.RegisteredClaims
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
}
