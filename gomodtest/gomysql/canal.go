package main

import (
	"log"

	"github.com/go-mysql-org/go-mysql/canal"
)

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	log.Printf("%s %v\n", e.Action, e.Rows)
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

func main() {
	cf := canal.NewDefaultConfig()
	cf.Password = "rootroot"
	cf.Dump.TableDB = "test"
	cf.Dump.Tables = []string{"test_t"}

	c, err := canal.NewCanal(cf)
	if err != nil {
		log.Fatal(err)
	}
	c.SetEventHandler(&MyEventHandler{})
	c.Run()
}
