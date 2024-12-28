package models

type Review struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	ItemID  int    `json:"item_id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
