package error

import "testing"

func TestErrorCf(t *testing.T) {
	err := New(90001, "aaaaaa")
	if err != nil {
		println(err.Error())
	}
}
