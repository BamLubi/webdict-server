package config

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
	"os"
)

var RD redis.Conn

func init() {
	var err error
	if len(os.Args) > 1 && os.Args[1] == "-DEV" {
		viper.SetConfigName("dev")
	} else {
		viper.SetConfigName("app")
	}
	viper.SetConfigType("properties")
	viper.AddConfigPath("./")
	url := viper.GetString("redis.url")
	RD, err = redis.Dial("tcp", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis Connect Success")
}
