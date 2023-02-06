package main

import (
	"main/api"
	"main/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
