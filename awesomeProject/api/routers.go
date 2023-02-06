package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"main/middleware"
	"net/http"
	"strings"
)

func InitRouter() {
	r := gin.Default()
	r.Use(Cors())
	u := r.Group("/user") //用户
	{
		u.POST("/register", register)                                  //注册
		u.POST("/login", login)                                        //登录
		u.POST("/address", middleware.AuthMiddleware(), address)       //添加(修改)收货地址
		u.PUT("/password", middleware.AuthMiddleware(), alterPassword) //修改密码
		u.PUT("/username", middleware.AuthMiddleware(), alterUsername) //修改用户名(昵称)
		u.GET("/info", middleware.AuthMiddleware(), GetInfo)           //个人信息
		u.PUT("/view")                                                 //添加个人头像...
		u.PUT("/quit", middleware.AuthMiddleware(), leave)             //退出登录
	}
	p := r.Group("/products") //产品
	{
		p.GET("/all", middleware.AuthMiddleware(), Info)       //返回全部商品信息//
		p.POST("/sort", middleware.AuthMiddleware(), products) //按关键字搜索部分商品//
		p.POST("/write")                                       //评价商品...
		p.GET("/message")                                      //查看商品评价...
	}
	s := r.Group("/cart") //购物车
	{
		s.GET("/all", middleware.AuthMiddleware(), ShowCart)             //返回购物车信息
		s.POST("/add", middleware.AuthMiddleware(), AddProductsIntoCart) //添加商品购物车
		s.PUT("/delete", middleware.AuthMiddleware(), RmCart)            //删除购物车指定商品
		s.POST("/buy", middleware.AuthMiddleware(), ChangeCondition)     //购买购物车商品
	}
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("err:%s", err)
		return
	}
}

func Cors() func(c *gin.Context) {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		// 允许所有header
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有的域,也可以指定某几个特定的域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			// header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma") //允许跨域设置可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")                                                                                                           // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                                                                                                                                     // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                                                                                                                            // 跨域请求是否需要带cookie信息 默认设置为true
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
