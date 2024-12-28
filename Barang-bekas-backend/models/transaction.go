package models

type Transaction struct {
	ID             int     `json:"id"`
	OfferID        int     `json:"offer_id"`
	TotalPrice     float64 `json:"total_price"`
	TransactionDate string  `json:"transaction_date"`
}
