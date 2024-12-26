package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/beego/beego/v2/server/web"
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var ctx = context.Background()

func Init() {
	addr := web.AppConfig.DefaultString("redis", "127.0.0.1:6379")
	password := web.AppConfig.DefaultString("redis_password", "")
	db := web.AppConfig.DefaultInt("redis_db", 0)

	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	fmt.Println("Connected to Redis")
}
