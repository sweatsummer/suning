package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type respTemplate struct {
	Status int    `json:"status" gorm:"status"`
	Info   string `json:"info" gorm:"info"`
}

// register 注册//
var register = respTemplate{
	Status: 200,
	Info:   "registered successfully",
}

// Login 登录//
var login = respTemplate{
	Status: 200,
	Info:   "login successfully",
}

var paramError = respTemplate{
	Status: 300,
	Info:   "param error",
}

var update = respTemplate{
	Status: 200,
	Info:   "update successfully",
}

var InternalErr = respTemplate{
	Status: 500,
	Info:   "internal err",
}

var PowerValid = respTemplate{
	Status: 401,
	Info:   "权限不足",
}

var rm = respTemplate{
	Status: 200,
	Info:   "delete successfully",
}

func RespOk(c *gin.Context) {
	c.JSON(http.StatusOK, register)
}
func RespParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, paramError)
}

func ResInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, InternalErr)
}

// NormErr 自定义错误//
func NormErr(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
}

func Update(c *gin.Context) {
	c.JSON(http.StatusOK, update)
}

func PowerErr(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, PowerValid)
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, rm)
}
