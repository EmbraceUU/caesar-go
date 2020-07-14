package geek_learning

import (
	"fmt"
	"reflect"
)

/*
learning reflect
*/

type AStruct struct {
	A string
	B bool
	C map[string]string
	D []int64
	E int64
}

func Action() {
	// 定义一个变量
	var aStruct AStruct
	aStruct.A = "aaa"
	aStruct.E = 9
	aStruct.C = make(map[string]string)

	// 拿到type
	t := reflect.TypeOf(aStruct)
	println(t.String())
	println(t.Name())
	println(t.Field(0).Name, t.Field(1).Index, t.Field(2).PkgPath)
	println(t.NumField())
	// 拿到value
	fmt.Println(reflect.ValueOf(aStruct))
}

func (a AStruct) String() string {
	return a.A
}
