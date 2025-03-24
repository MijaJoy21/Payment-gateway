package models

type CreateProduct struct {
	Name         string  `json:"name" form:"name"`
	CategoryId   int     `json:"category_id" form:"category_id"`
	Description  string  `json:"description" form:"description"`
	Price        float64 `json:"price" form:"price"`
	Status       int     `json:"status" form:"status"`
	IsPreOrder   int     `json:"is_preorder" form:"is_preorder"`
	PreOrderDate string  `json:"preorder_date" form:"preorder_date"`
	Weight       int     `json:"weight" form:"weight"`
}

type RequestProduct struct {
	Name        string `json:"name"`
	CategoryId  int    `json:"category_id"`
	Description string `json:"description"`
	// Image        string    `json:"image"`
	Price      float64 `json:"price"`
	Status     int     `json:"status"`
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
