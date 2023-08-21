package main

import (
	"ClickMaster/controllers"
	"ClickMaster/middleware"
	_ "ClickMaster/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 创建基于cookie的存储引擎，shuiche 参数是用于加密的密钥，可以随便填写
	store := cookie.NewStore([]byte("Boogiepop"))

	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎
	r.Use(sessions.Sessions("mysession", store))
	// 设置HTML模板文件目录
	r.LoadHTMLGlob("view/templates/*")
	r.Static("/css", "./view/css")
	r.Static("/img", "./view/img")
	r.GET("/", controllers.Index)
	r.GET("/register", controllers.GetRegister)
	r.POST("/register", controllers.PostRegister)
	r.POST("/login", controllers.PostLogin)
	r.GET("/admin", middleware.AuthMiddle, controllers.Admin)
	r.Run(":11451")
}
