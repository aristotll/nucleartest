package watch

import (
	"fmt"
)

type KV struct {
	K string
	V interface{}
}

type gedis struct {
	kv map[string]interface{}
	watch *watcher
}

func NewGedis() *gedis {
	return &gedis{
		kv: make(map[string]interface{}),
		watch: NewWatcher(),
	}
}

type watcher struct {
	watchKeys map[string]struct{}
	ch chan *KV
}

func NewWatcher() *watcher {
	return &watcher{
		watchKeys: make(map[string]struct{}),
		ch: make(chan *KV),
	}
}

func (g *gedis) Watch(key ...string) {
	for _, v := range key {
		g.watch.watchKeys[v] = struct{}{}
	}
	fmt.Printf("watching [%v]... \n", key)
	v := <-g.watch.ch
	fmt.Printf("key[%v] was changed! new val: %v \n", v.K, v.V)
}

func (g *gedis) Set(key string, val interface{}) {
	g.watch.checkoutWathcer(key, val)
	g.kv[key] = val
}

func (w *watcher) checkoutWathcer(key string, val interface{}) {
	if _, ok := w.watchKeys[key]; ok {
		w.ch <- &KV{
			K: key,
			V: val,
		}
	}
}

