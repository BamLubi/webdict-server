package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func init() {
	var err error
	if len(os.Args) > 1 && os.Args[1] == "-DEV" {
		viper.SetConfigName("dev")
	} else {
		viper.SetConfigName("app")
	}
	viper.SetConfigType("properties")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No file ...")
		} else {
			fmt.Println("Find file but have err ...")
		}
	}
	PORT = viper.GetString("server.port")
	HOST = viper.GetString("server.host")
	url := viper.GetString("db.url")
	db := viper.GetString("db.databases")
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	dsn := username + ":" + password + "@tcp(" + url + ")/" + db + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Mysql Connect Success")
}
