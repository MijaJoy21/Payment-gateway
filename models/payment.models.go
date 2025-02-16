package models

type ReqPayment struct {
	Amount          float64 `json:"amount"`
	PaymentMethodId int     `json:"payment_method_id"`
	Status          int     `json:"status"`
}
