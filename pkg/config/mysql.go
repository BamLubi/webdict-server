package config

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	Logger.Info("Mysql Connecting on " + DBUrl)
	dsn := DBUsername + ":" + DBPassword + "@tcp(" + DBUrl + ")/" + DBDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger.Error("Mysql Connect Fail" + err.Error())
		panic(err)
	}
	Logger.Info("Mysql Connect Success.")
}
