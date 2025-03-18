package models

type CreateProduct struct {
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
