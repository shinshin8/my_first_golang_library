package my_first_golang_library

import (
	"testing"
)

func TestHello(t *testing.T) {
	val:= Hello("hoge")
	if val != "" {
		t.Error("A returned value must be a string")
	}
}
