package api

import (
	"database/sql"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"log"
	"main/middleware"
	"main/model"
	"main/service"
	"main/util"
	"math/rand"
	"strings"
	"time"
)

var Letters = []string{"a", "s", "d", "f", "g", "h", "j", "k", "l", "z", "x",
	"c", "v", "b", "n", "m", "q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}

// register 用户注册//
func register(c *gin.Context) {
	UserName := c.PostForm("username")
	Password := c.PostForm("password")
	if len(UserName) == 0 || len(Password) == 0 {
		util.NormErr(c, 400, "注册密码或帐号不能为空...")
		return
	}
	u, err := service.IsExistUserByName(UserName)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error: %v", err)
		util.ResInternalErr(c)
		return
	}
	if u.UserName != "" {
		util.NormErr(c, 300, "用户名已存在...")
		return
	}
	for i := 0; i <= 25; i++ {
		if strings.Contains(Password, Letters[i]) == true || strings.Contains(Password, strings.ToUpper(Letters[i])) == true {
			break
		}
		if i == 25 {
			util.NormErr(c, 400, "密码需包含大写或小写字母...")
			return
		}
	}
	if len(Password) <= 5 {
		util.NormErr(c, 400, "密码不能少于6位")
		return
	}
	for {
		//随机生成5位编号//
		id := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000)
		u, err := service.IsExistUserByID(int(id)) //查重,防止唯一编号重复//
		if err != nil && err != sql.ErrNoRows {
			log.Printf("search user error:%v", err)
			util.ResInternalErr(c)
			return
		}
		if u.UserName == "" { //若用户编号不重复,插入并跳出循环//
			err = service.CreateUser(model.User{
				UserName: UserName,
				Password: base64.StdEncoding.EncodeToString([]byte(Password)), //base64加密传入//
				UserID:   int(id),
			})
			if err != nil {
				util.ResInternalErr(c)
				return
			}
			break
		}
	}
	util.RespOk(c)
}

// login 用户登录//
func login(c *gin.Context) {
	UserName := c.PostForm("username")
	Password := c.PostForm("password")
	if len(UserName) == 0 || len(Password) == 0 {
		util.NormErr(c, 400, "登录密码与帐号不能为空或空格...")
		return
	}
	u, err := service.IsExistUserByName(UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "帐号或密码错误")
			return
		} else {
			log.Printf("search user err %v", err)
			util.ResInternalErr(c)
			return
		}
	}
	//登录base64解密，判断是否正确//
	rePassword, err := base64.StdEncoding.DecodeString(u.Password)
	if err != nil {
		log.Fatalf("err :%s", err)
		return
	}
	if string(rePassword) != Password {
		util.NormErr(c, 200, "帐号或密码错误")
		return
	}
	//发放token//
	token, err := middleware.ReleaseToken(u)
	if err != nil {
		util.ResInternalErr(c)
		log.Printf("token generate error:%v", err)
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"Info":   "login successfully",
		"data":   gin.H{"token": token},
	})
}

// GetInfo 获取个人信息//
func GetInfo(c *gin.Context) {
	username, _ := c.Get("username")
	userid, _ := c.Get("userID")
	u, err := service.InfoByID(userid.(int))
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Search err :%s", err)
		return
	}
	c.JSON(200, gin.H{
		"userID":       userid,
		"username":     username,
		"user-balance": u.Balance,
		"user-address": u.Address,
	})
}

// alter 修改密码//
func alterPassword(c *gin.Context) {
	userid, _ := c.Get("userID")
	var u model.User
	if err := c.ShouldBind(&u); err != nil { //接收新密码
		log.Printf("err:%s", err)
		return
	}
	if len(u.Password) == 0 {
		util.NormErr(c, 400, "更改密码不能为空...")
		return
	}
	for i := 0; i <= 25; i++ {
		if strings.Contains(u.Password, Letters[i]) == true || strings.Contains(u.Password, strings.ToUpper(Letters[i])) == true {
			break
		}
		if i == 25 {
			util.NormErr(c, 400, "新密码需包含大写或小写字母...")
			return
		}
	}
	if len(u.Password) <= 5 {
		util.NormErr(c, 400, "新密码不能少于6位")
		return
	}
	err := service.UpdatePassword(model.User{
		UserID:   userid.(int),
		Password: base64.StdEncoding.EncodeToString([]byte(u.Password)), //base64加密写入//
	})
	if err != nil {
		util.ResInternalErr(c)
		log.Fatalf("err: %s", err)
		return
	}
	util.Update(c)
}

// alterUsername 修改用户名//
func alterUsername(c *gin.Context) {
	userid, _ := c.Get("userID")
	var u1 model.User
	if err := c.ShouldBind(&u1); err != nil {
		log.Printf("err:%s", err)
		return
	}
	if len(u1.UserName) == 0 {
		util.NormErr(c, 200, "更改用户名不能为空...")
		return
	}
	u, err := service.IsExistUserByName(u1.UserName)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error: %v", err)
		util.ResInternalErr(c)
		return
	}
	if u.UserName != "" {
		util.NormErr(c, 300, "用户名已存在...")
		return
	}
	err = service.UpdateUsername(model.User{
		UserID:   userid.(int),
		UserName: u1.UserName,
	})
	if err != nil {
		util.ResInternalErr(c)
		log.Fatalf("err : %s", err)
		return
	}
	util.Update(c)
}

// address 添加收货地址//
func address(c *gin.Context) {
	userid, _ := c.Get("userID")
	var u model.User
	if err := c.ShouldBind(&u); err != nil {
		log.Printf("err:%s", err)
		return
	}
	if len(u.Address) == 0 {
		util.NormErr(c, 400, "地址名无效...")
		return
	}
	err := service.AddAddress(model.User{
		UserID:  userid.(int),
		Address: u.Address,
	})
	if err != nil {
		util.ResInternalErr(c)
		log.Fatalf("err: %s", err)
		return
	}
	util.Update(c)
	return
}

// leave 退出登录//
func leave(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "已退出...",
	})
}
