package models

type CreateReview struct {
	Description string `json:"description" form:"description" validate:"required"`
	IsAnon      int    `json:"is_anon" form:"is_anon"`
}

type RequestReview struct {
	Description string `json:"description"`
	IsAnon      int    `json:"is_anon"`
}
