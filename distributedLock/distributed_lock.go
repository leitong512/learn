package distributedLock

import "github.com/go-redis/redis"

func incr()  {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 1,
	})
	var lock = "lock.foo"
	var couterKet
}
