package models

type RequestExpedition struct {
	Name   string `json:"name" validate:"required"`
	Price  int    `json:"price" validate:"required"`
	Weight int    `json:"weight" validate:"required"`
}

type ParamsGetExpeditions struct {
	Search string `json:"search" form:"search"`
	Limit  int    `json:"limit" form:"limit"`
	Page   int    `json:"page" form:"page"`
}
