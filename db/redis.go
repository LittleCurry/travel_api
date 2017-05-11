package db

import (
	"gopkg.in/redis.v5"
)

const REDIS_KEY_PREFIX = "AccessToken"
const REDIS_CODE_PREFIX = "VerifyCode"
const REDIS_REQUEST_PER_DAY_OF_DEVICE = "Device"

var client *redis.Client

func InitRedis(addr string) {
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
}

func Redis() *redis.Client {
	return client
}
