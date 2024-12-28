package models

type Offer struct {
	ID           int     `json:"id"`
	ItemID       int     `json:"item_id"`
	UserID       int     `json:"user_id"`
	OfferedPrice float64 `json:"offered_price"`
	Status       string  `json:"status"`
}
