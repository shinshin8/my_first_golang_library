package my_first_golang_library

import (
	"fmt"
	"reflect"
)

var err error

func Hello(name interface{}) (string, error) {
	if reflect.TypeOf(name).Kind() != reflect.String {
		return "", fmt.Errorf("%s is wrong input", name)
	}
	return fmt.Sprintf("hello %s\n", name), nil
}