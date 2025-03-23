package cache

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"simplelist/config"
	"strconv"
)

var RedisClient *redis.Client

func Redis() {
	db, _ := strconv.ParseUint(config.RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		logrus.Info(err)
		panic(err)
	}
	RedisClient = client
}
