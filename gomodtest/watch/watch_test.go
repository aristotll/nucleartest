package watch

import (
	//"log"
	"log"
	"testing"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

var g = NewGedis()

func TestSet(t *testing.T) {
	log.Println(g.watch.watchKeys)
	g.Set("name", "zhang3")
}

func TestWatch(t *testing.T) {

	g.Watch("name")
	log.Println(g.watch.watchKeys)
}
func Test1(t *testing.T) {
	go func() {
		g.Watch("name")
	}()

	time.Sleep(time.Second * 5)

	go func() {
		g.Set("name", "1")
	}()

	time.Sleep(time.Second * 10)
}
