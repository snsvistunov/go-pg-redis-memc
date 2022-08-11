package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var RC *redis.Client

const redisTTl int = 30

func ConnectRedis() *redis.Client {
	client := redis.NewClient(rConf)
	if err := client.Ping().Err(); err != nil {
		log.Fatal(err)
	}
	RC = client
	return client
}

func SetRedisValue(c *redis.Client, key string, value interface{}) {
	p, err := json.Marshal(value)
	if err != nil {
		log.Fatal(err)
	}
	if _, err = c.Set(key, p, time.Second*time.Duration(redisTTl)).Result(); err != nil {
		log.Fatal(err)
	}
}

func GetRedisValue(c *redis.Client, key string, dest interface{}) error {
	p, _ := c.Get(key).Result()
	return json.Unmarshal([]byte(p), dest)
}
