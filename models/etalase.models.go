package models

type RequestEtalase struct {
	Name string `json:"name" validate:"required"`
}
