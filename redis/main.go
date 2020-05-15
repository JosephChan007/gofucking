package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var client *redis.Client

// 连接Redis
func InitClient() (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     "hdfs-host4:6379",
		Password: "123456",
		DB:       0,
	})

	_, err = client.Ping().Result()
	if err != nil {
		fmt.Printf("redis client inti err: %v\n", err)
		return err
	}
	return nil
}

func GetVal(key string) (val interface{}, err error) {
	val, err = client.Get(key).Result()
	if err != nil {
		fmt.Printf("get value from redis failed, err:%v\n", err)
		return
	}
	return
}

func SetVal(key string, val interface{}) (err error) {
	return SetTimeVal(key, val, 0)
}

func SetTimeVal(key string, val interface{}, timeout int64) (err error) {
	err = client.Set(key, val, time.Duration(timeout)*time.Millisecond).Err()
	if err != nil {
		fmt.Printf("set value to redis failed, err:%v\n", err)
		return
	}
	return
}

func main() {

	InitClient()

	SetVal("good", "jean")
	SetVal("price", 100.23)

	good, err := GetVal("good")
	if err != nil {
		fmt.Printf("get value from redis failed, err: %v\n", err)
	}
	goodVal := good.(string)
	fmt.Printf("get value from redis success, good: %v\n", goodVal)

	price, err := GetVal("price")
	if err != nil {
		fmt.Printf("get value from redis failed, err: %v\n", err)
	}
	priceVal := price.(string)
	fmt.Printf("get value from redis success, price: %v\n", priceVal)

}
