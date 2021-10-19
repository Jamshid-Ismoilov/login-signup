package redis

import (
	"app/config"
	"time"
)

func SetRedis(email, password string) {

	config.Client.Set(email, password, time.Duration(time.Second *300))
}

func GetRedis(email string) (string, error) {
	return config.Client.Get(email).Result()
}