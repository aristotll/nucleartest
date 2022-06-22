package gorm

import (
	"sync"
	"testing"
)

func TestUpdateRowIsLock(t *testing.T) {
	var wg sync.WaitGroup
	var num = 100
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			if err := UpdateRowIsLock(mysqlConn); err != nil {
				return
			}
		}()
	}

	wg.Wait()
}
