package model

type Products struct {
	ProductsID      int    `json:"products_id" form:"products_id"`
	ProductsName    string `json:"products_name" form:"products_name"`
	ProductsPrice   int    `json:"products_price" form:"products_price"`
	ProductsComment int    `json:"products_comment" form:"products_comment"`
	ProductsSales   int    `json:"products_sales" form:"products_sales"`
	ProductsMessage string `json:"products_message" form:"products_message"`
}
