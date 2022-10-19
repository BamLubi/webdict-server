package config

import (
	"github.com/garyburd/redigo/redis"
)

var RD redis.Conn

func InitRedis() {
	var err error
	Logger.Info("Redis Connecting on " + RedisUrl)
	RD, err = redis.Dial("tcp", RedisUrl)
	if err != nil {
		Logger.Error("Redis Connect Fail" + err.Error())
		panic(err)
	}
	Logger.Info("Redis Connect Success.")
}
