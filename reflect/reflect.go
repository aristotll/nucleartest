package reflect

import "reflect"

func fn() {
	i := 123
	reflect.ValueOf(i)
}