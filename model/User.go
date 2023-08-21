package model

type User struct {
	Id       int    `gorm:"column:id;primary_key;auto_increment"`
	Username string `gorm:"column:username;primary_key;size:200" form:"username"`
	Password string `gorm:"column:password;primary_key;size:200" form:"password"`
}
