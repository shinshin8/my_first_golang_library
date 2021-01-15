package my_first_golang_library

import (
	"fmt"
	"reflect"
)

func Hello(name interface{}) string {
	if reflect.TypeOf(name).Kind() != reflect.String {
		return ""
	}
	return fmt.Sprintf("hello %s\n", name)
}