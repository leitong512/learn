package distributedLock

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func Incr() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
	})
	ctx := context.Background()
	err := client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key :", val)
}
