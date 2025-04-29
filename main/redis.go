package main

import (
	"context"
	"github.com/aobco/log"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName: "mymaster",
		SentinelAddrs: []string{
			"redis-master.default:26379",
		},
		SentinelPassword: "Xsky@123",
		Password:         "Xsky@123",
		DB:               0,
		TLSConfig:        nil,
	})
	log.Infof("start to set value")
	if _, err := client.Set(context.Background(), "set", "set123", time.Minute).Result(); err != nil {
		log.Errorf("set failed, %v", err)
		return
	}
	log.Infof("set value success")
}
