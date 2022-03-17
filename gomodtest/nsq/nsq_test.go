package main

import "testing"

// keepalive 貌似不能从 os.stdin 中读取，下面会立马结束，production() 需要在 main 中运行
func TestProduction(t *testing.T) {
	production()
}

func TestConsumption1(t *testing.T) {
	consumption1()
}

func TestConsumption2(t *testing.T) {
	consumption2()
}
