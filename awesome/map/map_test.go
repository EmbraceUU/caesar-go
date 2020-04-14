package _map

import (
	"fmt"
	"testing"
)

func TestWriteMap(t *testing.T) {
	//fmt.Println(WriteMap("m", "a", 12.1111))
}

func TestReadMap(t *testing.T) {
	InitMap()
	var aaa AAA
	aaa.D = make(map[string]string)
	aaa.D["t1"] = "v1"
	WriteMap("out1", &aaa)
	fmt.Println(fmt.Sprintf("%v", TMap))
	rTMap := ReadMap("out1")
	fmt.Println(fmt.Sprintf("%v", rTMap))
	rTMap.D["t2"] = "t2"
	fmt.Println(fmt.Sprintf("%v", TMap))
	fmt.Println(fmt.Sprintf("%v", rTMap))
}
