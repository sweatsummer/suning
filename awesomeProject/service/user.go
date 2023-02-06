package service

import (
	"main/dao"
	"main/model"
)

func InfoByID(ID int) (u model.User, err error) {
	u, err = dao.SearchInfoByID(ID)
	return
}

func IsExistUserByID(ID int) (u model.User, err error) {
	u, err = dao.SearchUserByID(ID)
	return
}

func IsExistUserByName(name string) (u model.User, err error) {
	u, err = dao.SearchUserByUserName(name)
	return
}

func CreateUser(u model.User) error {
	err := dao.InsertUser(u)
	return err
}

func UpdatePassword(u model.User) error {
	err := dao.AlterPassword(u)
	return err
}

func UpdateUsername(u model.User) error {
	err := dao.AlterUsername(u)
	return err
}

func AddAddress(u model.User) error {
	err := dao.AddAddress(u)
	return err
}

func UpdateUserBalance(u model.User) error {
	err := dao.AlterUserBalance(u)
	return err
}
