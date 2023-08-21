package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var dsn = "root:jiayou357753@tcp(127.0.0.1:3306)/ctf?charset=utf8mb4&parseTime=True&loc=Local"
var DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
	Logger: logger.Default.LogMode(logger.Info), // 设置日志级别为 Info
})

func init() {
	if DB.Migrator().HasTable(&User{}) {
		red := "\033[31m"
		reset := "\033[0m"
		log.Println("%s%s%s\n", red, "Table Exists", reset)
	} else {
		DB.AutoMigrate(&User{})
		DB.Create(&User{Id: 1, Username: "admin", Password: "SSXmMs123!"})
	}
}

func AddUser(user User) bool {
	if err := DB.Save(&user).Error; err != nil {
		log.Println("Fail to Add User!")
		return false
	}
	return true
}

func GetUserByName(name string) User {
	var result User
	if err := DB.Where("username = ?", name).First(&result).Error; err != nil {
		log.Println("Fail to search User!")
	}
	return result
}
