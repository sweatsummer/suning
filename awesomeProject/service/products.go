package service

import (
	"main/dao"
	"main/model"
)

func IsExistProducts(ID int) (p model.Products, err error) {
	p, err = dao.SearchProductsByID(ID)
	return
}

func IsExistProductsByID(id int) (s model.Cart, err error) {
	s, err = dao.SearchProductsByID2(id)
	return
}

func CreateCart(s model.Cart) error {
	err := dao.InsertInfo(s)
	return err
}

func RemoveCart(s model.Cart) error {
	err := dao.DeleteCart(s)
	return err
}

func UpdateCondition(s model.Cart) error {
	err := dao.AlterStatus(s)
	return err
}

func GetInfoFromCart(id int, i int) (s model.Cart, err error) {
	s, err = dao.SelectInfoByID(id, i)
	return
}

func GetStatus(id int) (s model.Cart, err error) {
	s, err = dao.SelectStatus(id)
	return
}
