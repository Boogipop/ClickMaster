package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddle(c *gin.Context) {
	session := sessions.Default(c)
	role := session.Get("admin")
	if role == "1" {
		c.Next()
	} else {
		c.JSON(401, gin.H{"status": 401, "message": "fail to auth"})
	}
}
