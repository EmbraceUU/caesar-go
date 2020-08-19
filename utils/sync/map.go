package sync

import (
	"caesar-go/utils/time"
	"fmt"
	"sync"
)

type userInfo struct {
	Name string
	Age  int
}

var m sync.Map

var mm map[string]int

func MapOpCom() {
	if mm == nil {
		mm = make(map[string]int)
	}

	mm["1"] = time.NowTs()
	mm["2"] = time.NowTs()

	fmt.Println(mm["1"])
	fmt.Println(mm["2"])

	mm["1"] = time.NowTs()
	mm["2"] = time.NowTs()
}

func MapOp() {
	vv, ok := m.LoadOrStore("1", "one")
	fmt.Println(vv, ok)

	vv, ok = m.Load("1")
	fmt.Println(vv, ok)

	vv, ok = m.LoadOrStore("1", "oneOne")
	fmt.Println(vv, ok)

	m.Store("1", "oneOne")
	vv, ok = m.Load("1")
	fmt.Println(vv, ok)

	m.Store("2", "two")
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	m.Delete("1")
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	map1 := make(map[string]userInfo)
	var user1 userInfo
	user1.Name = "amy"
	user1.Age = time.NowTs()
	map1["user1"] = user1

	var user2 userInfo
	user2.Name = "Tom"
	user2.Age = time.NowTs()

	m.Store("map_test", &user1)

	mapValue, _ := m.Load("map_test")
	mapValue.(*userInfo).Name = "eeeeee"

	value, ok := m.Load("map_test")
	fmt.Println(value, ok)

	//for k, v := range mapValue.(interface{}).(map[string]userInfo) {
	//    fmt.Println(k, v)
	//    fmt.Println("name:", v.Name)
	//}
}
