package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("count") == nil {
		session.Set("count", 0)
		session.Save()
	}
	submit := c.DefaultQuery("submit", "false")
	if submit == "true" {
		count := session.Get("count").(int)
		count += 1
		session.Set("count", count)
		session.Save()
		c.HTML(200, "welcome.html", gin.H{"count": count})
	} else {
		c.HTML(200, "welcome.html", gin.H{"count": session.Get("count").(int)})
	}
	if session.Get("count").(int) == 100000 {
		c.HTML(200, "welcome.html", gin.H{"message": "HnuSec{Y0u_aRe_Really_Cl1ck_M4ster}"})
	}
}
