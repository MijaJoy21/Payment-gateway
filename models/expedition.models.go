package models

type RequestExpedition struct {
	Name   string `json:"name" validate:"required"`
	Price  int    `json:"price" validate:"required"`
	Weight int    `json:"weight" validate:"required"`
}
