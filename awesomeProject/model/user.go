package model

type User struct {
	UserID   int    `json:"id" form:"id"`
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Balance  int    `json:"balance" form:"balance"`
	Address  string `json:"address" form:"address"`
}
