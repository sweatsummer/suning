package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"main/model"
	"main/service"
	"main/util"
	"net/http"
	"strings"
)

// Info 所有商品信息//
func Info(c *gin.Context) {
	//循环读取商品数据//
	for i := 1; i >= 0; i++ {
		p, err := service.IsExistProducts(i)
		if err != nil {
			if err == sql.ErrNoRows {
				break
			} else {
				util.ResInternalErr(c)
				log.Fatalf("Search err:%s", err)
				return
			}
		}
		//保存输出商品数据//
		c.JSON(http.StatusOK, p)
	}
	return
}

// products 商品信息//
func products(c *gin.Context) {
	//获取关键字//
	type keyword struct {
		Keyword string `json:"keyword" form:"keyword"`
	}
	var k keyword
	if err := c.Bind(&k); err != nil {
		log.Printf("err:%s", err)
		return
	}
	//判断关键字是否合法//
	if len(k.Keyword) == 0 {
		util.NormErr(c, 400, "关键字不能为空...")
		return
	}
	for i := 1; i >= 0; i++ {
		p, err := service.IsExistProducts(i)
		if err != nil {
			if err == sql.ErrNoRows {
				break
			} else {
				util.ResInternalErr(c)
				log.Fatalf("Search err:%s", err)
				return
			}
		}
		if strings.Contains(p.ProductsName, k.Keyword) == true {
			c.JSON(200, p)
		}
	}
	return
}

// ShowCart 展示购物车信息//
func ShowCart(c *gin.Context) {
	userid, _ := c.Get("userID")
	//循环选取信息//
	for i := 1; i >= 0; i++ {
		s, err := service.GetInfoFromCart(userid.(int), i)
		if err != nil {
			if err == sql.ErrNoRows { //读完跳出
				break
			} else {
				util.ResInternalErr(c)
				log.Fatalf("Search err: %s", err)
				return
			}
		}
		c.JSON(http.StatusOK, s)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "finish...",
	})
	return
}

// AddProductsIntoCart 添加购物车//
func AddProductsIntoCart(c *gin.Context) {
	userid, _ := c.Get("userID")
	var products model.Products
	//获取添加商品编号//
	if err := c.ShouldBind(&products); err != nil {
		log.Printf("err:%s", err)
		return
	}
	//根据商品id查重//
	s, err := service.IsExistProductsByID(products.ProductsID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search err:%s", err)
		util.ResInternalErr(c)
		return
	}
	if s.UserID == userid.(int) {
		util.NormErr(c, 400, "该商品已存在...")
		return
	}
	//插入用户id与商品id//
	err = service.CreateCart(model.Cart{
		UserID:     userid.(int),
		ProductsID: products.ProductsID,
		Status:     "tentative",
	})
	if err != nil {
		util.ResInternalErr(c)
		return
	}
	//成功插入//
	util.Update(c)
}

// RmCart 删除购物车商品//
func RmCart(c *gin.Context) {
	userid, _ := c.Get("userID")
	var products model.Products
	if err := c.ShouldBind(&products); err != nil {
		log.Printf("err : %s", err)
		return
	}
	//移除购物车商品//
	err := service.RemoveCart(model.Cart{
		UserID:     userid.(int),
		ProductsID: products.ProductsID,
	})
	if err != nil {
		log.Printf("err:%s", err)
		util.ResInternalErr(c)
		return
	}
	//移除成功//
	util.Delete(c)
	return
}

// ChangeCondition 购买商品改变购物车状态//
func ChangeCondition(c *gin.Context) {
	userid, _ := c.Get("userID")
	var products model.Products
	if err := c.ShouldBind(&products); err != nil { //接受前端确认购买信息
		log.Printf("err:%s", err)
		return
	}
	//购买查重//
	s, err := service.GetStatus(userid.(int))
	if err != nil && err != sql.ErrNoRows {
		util.ResInternalErr(c)
		return
	}
	if s.Status == "down" {
		util.NormErr(c, 400, "已购买...")
		return
	}
	u, err := service.InfoByID(userid.(int)) //查询用户余额
	if err != nil && err != sql.ErrNoRows {
		util.ResInternalErr(c)
		return
	}
	//查询商品所需价格
	p, err := service.IsExistProducts(products.ProductsID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Search err: %s", err)
		util.ResInternalErr(c)
		return
	}
	//比较返回购买信息
	if p.ProductsPrice > u.Balance {
		util.NormErr(c, 400, "余额不足...")
		return
	}
	//更新用户余额
	err = service.UpdateUserBalance(model.User{
		UserID:  userid.(int),
		Balance: u.Balance - p.ProductsPrice,
	})
	if err != nil {
		util.ResInternalErr(c)
		return
	}
	//更新购物车状态
	err = service.UpdateCondition(model.Cart{
		UserID:     userid.(int),
		ProductsID: products.ProductsID,
		Status:     "down",
	})
	if err != nil {
		util.ResInternalErr(c)
		return
	}
	util.Update(c)
	return
}
