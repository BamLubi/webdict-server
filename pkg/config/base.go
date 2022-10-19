package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var ServerUrl string
var ServerSSLCert string
var ServerSSLKey string
var RedisUrl string
var DBUrl string
var DBDatabase string
var DBUsername string
var DBPassword string
var MODE string
var LogPath string
var LogLevel string

func init() {
	// 设置配置文件
	if len(os.Args) > 1 && os.Args[1] == "-DEV" {
		MODE = "DEVELOPMENT"
		viper.SetConfigName("dev")
	} else if len(os.Args) > 1 && os.Args[1] == "-PROD_TEST" {
		MODE = "PRODUCTION TEST"
		viper.SetConfigName("app")
	} else {
		MODE = "PRODUCTION"
		viper.SetConfigName("app")
	}
	viper.SetConfigType("properties")
	viper.AddConfigPath("./")
	// 测试是否可以打开文件
	err := viper.ReadInConfig()
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
	// 获取不同的参数
	// 日志
	LogPath = viper.GetString("logger.path")
	LogLevel = viper.GetString("logger.level")
	// GO服务
	ServerUrl = viper.GetString("server.host") + ":" + viper.GetString("server.port")
	ServerSSLCert = viper.GetString("server.ssl.cert")
	ServerSSLKey = viper.GetString("server.ssl.key")
	// Redis
	RedisUrl = viper.GetString("redis.url")
	// Database
	DBUrl = viper.GetString("db.url")
	DBDatabase = viper.GetString("db.databases")
	DBUsername = viper.GetString("db.username")
	DBPassword = viper.GetString("db.password")
}
