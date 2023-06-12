package main

import (
	"fmt"
	"strconv"
	"sync"
)

type GetValFunc func(k string) (val any, err error)

type call struct {
	val any
	err error
	wg  *sync.WaitGroup
}

type singleflight struct {
	sync.Mutex
	m map[string]*call
}

func NewSingleflight() *singleflight {
	return &singleflight{
		m: make(map[string]*call),
	}
}

func (s *singleflight) Do(reqID string, key string, fn GetValFunc) (val any, err error) {
	s.Lock()

	if c, ok := s.m[key]; ok {
		s.Unlock()
		c.wg.Wait()
		fmt.Printf("reqID [%v] wait value\n", reqID)
		return c.val, c.err
	}

	s.m[key] = &call{wg: new(sync.WaitGroup)}
	s.Unlock()

	s.m[key].wg.Add(1)
	v, err := fn(key)
	if err != nil {
		s.m[key].err = err
		return
	}
    s.m[key].wg.Done()
	fmt.Printf("reqID [%v] is set value to cache\n", reqID)

	s.m[key].val = v
	s.m[key].err = err

    s.Lock()
    delete(s.m, key)
    s.Unlock()
	return v, nil
}

func main() {
	s := NewSingleflight()
	wg := sync.WaitGroup{}
	getFunc := GetValFunc(func(k string) (val any, err error) {
		return "val", nil
	},
	)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Do(strconv.Itoa(i), "key", getFunc)
		}()
	}
	wg.Wait()
}
