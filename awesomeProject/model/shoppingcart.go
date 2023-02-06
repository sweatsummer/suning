package model

type Cart struct {
	UserID     int    `json:"user_id" form:"user_id"`
	ProductsID int    `json:"products_id" form:"products_id"`
	Status     string `json:"status" form:"status"`
}
