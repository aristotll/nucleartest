package main

import (
	"context"
	"fmt"
	"github.com/golang/groupcache"
	"log"
)

var cache = map[string]string{
	"a": "1",
	"b": "2",
	"c": "3",
}

type GetData struct {}

func (g *GetData) Get(ctx context.Context, key string, dest groupcache.Sink) error {
	v, ok := cache[key]
	if !ok {
		return fmt.Errorf("not find")
	}
	if err := dest.SetString(v); err != nil {
		return err
	}
	return nil
}

func main() {
	group := groupcache.NewGroup("test1", 1024, &GetData{})
	var res string
	if err := group.Get(context.Background(), "a", groupcache.StringSink(&res)); err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(res)

}
