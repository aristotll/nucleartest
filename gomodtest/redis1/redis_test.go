package redis1

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v9"
)

var redisConn *redis.Client

func init() {
	redisConn = Conn()
}

func TestSetNX(t *testing.T) {
	for i := 0; i < 10; i++ {
		go setNX(redisConn, "lock", "", fmt.Sprintf("service:%d", i))
	}
	time.Sleep(time.Second * 50)
}
