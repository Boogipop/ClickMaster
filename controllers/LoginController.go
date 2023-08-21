package controllers

import (
	"ClickMaster/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
	"net/http"
)

func PostLogin(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")
	userinfo := model.GetUserByName(username)
	if userinfo.Password == password && username == "admin" {
		session.Set("admin", "1")
		session.Save()
		c.Redirect(302, "/admin")
	} else {
		session.Set("admin", "0")
		c.Redirect(302, "/")
	}
}
func Index(c *gin.Context) {
	// 渲染HTML模板
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
func GetRegister(c *gin.Context) {
	// 渲染HTML模板
	c.HTML(http.StatusOK, "register.html", gin.H{})
}
func PostRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := model.User{
		Username: username,
		Password: password,
	}
	success := model.AddUser(user)
	if success {
		c.Redirect(302, "/")
	} else {
		c.HTML(http.StatusOK, "register.html", gin.H{"message": "注册失败"})
	}
}
