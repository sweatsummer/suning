package dao

import (
	"main/model"
)

func SearchInfoByID(ID int) (u model.User, err error) {
	row := DB.QueryRow("select balance,address from user where ID=?", ID)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Balance, &u.Address)
	return
}

// SearchUserByUserName 查询数据库//
func SearchUserByUserName(name string) (u model.User, err error) {
	row := DB.QueryRow("select id,username,password from user where username=?", name)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.UserID, &u.UserName, &u.Password)
	return
}

// SearchUserByID 依据编号ID查询数据库//
func SearchUserByID(ID int) (u model.User, err error) {
	row := DB.QueryRow("select username,password from user where id=?", ID)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.UserName, &u.Password)
	return
}

// InsertUser 插入数据库注册数据//
func InsertUser(u model.User) (err error) {
	_, err = DB.Exec("insert into user(id,username,password,balance) values (?,?,?,?)", u.UserID, u.UserName, u.Password, 0)
	if err != nil {
		return
	}
	return err
}

// AlterPassword 修改密码//
func AlterPassword(u model.User) (err error) {
	_, err = DB.Exec("update user set password=? where id=?", u.Password, u.UserID)
	if err != nil {
		return
	}
	return err
}

// AlterUsername 修改用户名//
func AlterUsername(u model.User) (err error) {
	_, err = DB.Exec("update user set username=? where id=?", u.UserName, u.UserID)
	if err != nil {
		return
	}
	return err
}

func AddAddress(u model.User) (err error) {
	_, err = DB.Exec("update user set address=? where id=? ", u.Address, u.UserID)
	if err != nil {
		return
	}
	return err
}

func AlterUserBalance(u model.User) (err error) {
	_, err = DB.Exec("update user set balance=? where id=?", u.Balance, u.UserID)
	if err != nil {
		return
	}
	return
}
