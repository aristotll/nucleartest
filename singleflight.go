package main

import (
	"log"
	"sync"
	"time"
)

var db = map[string]int64{
	"TOM":  18,
	"JACK": 20,
	"WANG": 22,
}

type singleflight struct {
	cache map[string]any
	// 大坑：因为 map 的 value 类型定义成了 sync.WaitGroup 而导致死锁问题
	// 找了半天死活发现不了死锁的原因，浪费了几个小时时间
	// 其实使用 go vet 可以检测出问题：
	// go vet "/Users/zenghao/pj/justtest/singleflight.go"
	// # command-line-arguments
	// pj/justtest/singleflight.go:29:15: assignment copies lock value to wg: (sync.WaitGroup, bool) contains sync.WaitGroup contains sync.noCopy
	// pj/justtest/singleflight.go:41:21: assignment copies lock value to s.flightKey[key]: sync.WaitGroup contains sync.noCopy
	flightKey map[string]*sync.WaitGroup // 正在请求 db 的 key
	sync.Mutex
}

var ids int64

func (s *singleflight) get(id int, key string) any {
	s.Lock()
	if v, ok := s.cache[key]; ok {
		s.Unlock()
		return v
	}
	if wg, ok := s.flightKey[key]; ok {
		s.Unlock()
		log.Printf("id=%v wait\n", id)
		wg.Wait()
		return s.cache[key]
	}
	// 第一个请求暂停一会，方便测试效果，看后面的请求会不会等待
	// 直到第一个请求返回
	log.Printf("id=%v first request, sleep 5s\n", id)
	wg := sync.WaitGroup{}
	wg.Add(1)
	defer wg.Done()
	s.flightKey[key] = &wg
	s.Unlock()
	v := s.getFromDB(id, key)
	if id != 999 {
		time.Sleep(time.Second * 5)
	}

	s.Lock()
	delete(s.flightKey, key)
	s.cache[key] = v
	s.Unlock()
	return v
}

func (s *singleflight) getFromDB(id int, key string) any {
	log.Printf("id=%v get from DB\n", id)
	return db[key]
}

var sf = &singleflight{
	cache:     make(map[string]any),
	flightKey: make(map[string]*sync.WaitGroup),
}

func main() {
	count := 50
	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i_ int) {
			defer wg.Done()
			v := sf.get(i_, "TOM")
			log.Printf("id=%v, value: %v\n", i_, v)
		}(i)
	}
	v := sf.get(999, "JACK")
	log.Printf("key=JACK, v=%v\n", v)
	wg.Wait()
}
