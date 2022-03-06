package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

var redisConn = redis.NewClient(&redis.Options{
	// host 必须指定为 redis，不能是 localhost
	Addr:     "redis:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func main() {
	var msg string
	res, err := redisConn.Ping(context.TODO()).Result()
	if err != nil {
		log.Println("err: ", err)
		msg = fmt.Sprintf("ping redis error: %v", err)
	} else {
		msg = fmt.Sprintf("ping redis success: %v", res)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(msg))
		if err != nil {
			log.Println("write error: ", err)
		}
	})
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hello"))
		if err != nil {
			log.Println("write error: ", err)
		}
	})

	if err := http.ListenAndServe(":6666", nil); err != nil {
		panic(err)
	}
}
