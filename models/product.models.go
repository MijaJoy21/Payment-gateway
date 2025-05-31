package models

import "time"

type CreateProduct struct {
	Name         string          `json:"name" form:"name" validate:"required"`
	CategoryId   int             `json:"category_id" form:"category_id" validate:"required"`
	Description  string          `json:"description" form:"description" validate:"required"`
	Variant      []CreateVariant `json:"variant" form:"variant"`
	Price        float64         `json:"price" form:"price" validate:"required"`
	Status       *int            `json:"status" form:"status"`
	IsPreOrder   int             `json:"is_preorder" form:"is_preorder"`
	PreOrderDate string          `json:"preorder_date" form:"preorder_date"`
	Weight       int             `json:"weight" form:"weight" validate:"required"`
}

type RequestProduct struct {
	Name        string `json:"name"`
	CategoryId  int    `json:"category_id"`
	Description string `json:"description"`
	// Image        string    `json:"image"`
	Price      float64 `json:"price"`
	Status     *int    `json:"status"`
	IsPreOrder int     `json:"is_preorder"`
	// PreOrderDate time.Time `json:"preorder_date"`
	Weight int `json:"weight"`
}

type ParamsGetProduct struct {
	Page       int    `json:"page" form:"page"`
	Limit      int    `json:"limit" form:"limit"`
	Search     string `json:"search" form:"search"`
	CategoryId int    `json:"category_id" form:"category_id"`
}
type CreateVariant struct {
	Name     string `json:"name" validate:"required"`
	Price    int    `json:"price" validate:"required"`
	Weight   int    `json:"weight" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
	Status   *int   `json:"status"`
}

type GetProductByIdResponse struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	CategoryId   int        `json:"category_id"`
	Description  string     `json:"description"`
	Image        []string   `json:"image"`
	Quantity     int        `json:"quantity"`
	Price        float64    `json:"price"`
	Status       *int       `json:"status"`
	IsPreorder   int        `json:"is_preorder"`
	PreOrderDate *time.Time `json:"pre_orderdate"`
	Weight       int        `json:"weight"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type GetProductsResponse struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Image        string     `json:"image"`
	Quantity     int        `json:"quantity"`
	Price        float64    `json:"price"`
	Status       *int       `json:"status"`
	IsPreorder   int        `json:"is_preorder"`
	PreOrderDate *time.Time `json:"pre_orderdate"`
	Weight       int        `json:"column"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
