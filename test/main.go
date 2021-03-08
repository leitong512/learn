package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func main() {
	key := "test"
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB: 0,
	})
	client.ZAdd(context.TODO(),key,&redis.Z{
		Member: "google.com",
		Score: float64(9),
	})
	client.ZAdd(context.TODO(),key,&redis.Z{
		Member: "baidu.com",
		Score: float64(12),
	})
	client.ZAdd(context.TODO(),key,&redis.Z{
		Member: "bing.com",
		Score: float64(16),
	})

	//idZRange := &redis.ZRangeBy{
	//	Max: "9",
	//	Min: "-inf",
	//}
	//lists, _ := client.ZRangeByScore(context.TODO(),key,idZRange).Result()
	//fmt.Println(lists)
	//for _, list := range lists {
	//	client.ZRem(context.TODO(),key,list)
	//	fmt.Println("删除：" + list)
	//}
}
