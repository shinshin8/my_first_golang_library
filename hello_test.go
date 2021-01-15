package my_first_golang_library

import (
	"testing"
)

func TestHello(t *testing.T) {
	_, err := Hello("hoge")
	if err != nil {
		t.Error("A returned value must be a string")
	}
}
