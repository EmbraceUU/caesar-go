package _map

import (
	"fmt"
	"testing"
)

func TestWriteMap(t *testing.T) {
	fmt.Println(WriteMap("m", "a", 12.1111))
}

func TestReadMap(t *testing.T) {
	InitMap()
	TMap["t"] = map[string]string{"tk1": "tv1", "tk2": "tv2"}
	rTMap := ReadMap("t")
	TMap["t"]["tk3"] = "tv3"
	fmt.Println(rTMap)
	fmt.Println(TMap)
}
