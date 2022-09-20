package redisT

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func getDb() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.219.131:6379",
		Password: "root",
		DB:       0,
	})
	result, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return client
}

func getString(arg string) {
	db := getDb()
	result, err := db.Get(arg).Result()
	if err != nil {
		fmt.Println(err)
		return
	} else if err == redis.Nil {
		fmt.Println("key 不存在")
	}
	defer db.Close()
	fmt.Println(result)
}

func setString(key, v string) bool {
	db := getDb()
	err := db.Set(key, v, 0).Err()
	if err != nil {
		panic("set fail")
	}
	defer db.Close()
	return true
}

func setAdd(key, v string) int64 {
	db := getDb()
	result, err := db.SAdd(key, v).Result()
	if err != nil {
		fmt.Println("set err：", err)
		return 0
	}
	defer db.Close()
	return result
}

func getSetAll(key string) {
	db := getDb()
	defer db.Close()
	members := db.SMembers(key)
	val := members.Val()
	for _, v := range val {
		fmt.Println(v)
	}
}

func hsetGet(obj, k string) {
	db := getDb()
	defer db.Close()
	val := db.HGet(obj, k).Val()
	fmt.Println(val)
}

func hmget(obj string, k ...string) {
	db := getDb()
	defer db.Close()
	for i := 0; i < 3; i++ {
		get := db.HMGet(obj, k[i])
		fmt.Println(get)
	}
}

func zset(obj string, z []*redis.Z) {
	db := getDb()
	defer db.Close()
	result, err := db.ZAdd(obj, z...).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}