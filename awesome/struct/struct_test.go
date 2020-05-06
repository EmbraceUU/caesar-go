package _struct

import (
	"fmt"
	"testing"
)

func TestGetBoolStruct(t *testing.T) {
	sc := GetBoolStruct()
	fmt.Println(sc)
}

func TestGetPairHttpData(t *testing.T) {
	GetPairHttpData()
}

func TestNilValue(t *testing.T) {
	NilValue()
}
