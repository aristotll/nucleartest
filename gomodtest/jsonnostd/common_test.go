package jsonnostd

import "testing"

func TestConv(t *testing.T) {
	if GlobalJsonString != string(GlobalJsonByte) {
		t.Error("string not equal byte slice")
	}
}
