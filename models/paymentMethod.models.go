package models

type ReqPaymentMethod struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}
