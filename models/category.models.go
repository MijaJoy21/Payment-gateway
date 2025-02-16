package models

type RequestCategory struct {
	Name string `json:"name" validate:"required"`
}
