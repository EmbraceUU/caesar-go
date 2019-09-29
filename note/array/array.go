package array

import "fmt"

// 定义后的arr的零值是什么
func ArrDefineZero() []string {
	var arrStr []string
	return arrStr
}

// 定义后的arr[结构体]的零值是什么
func ArrDefineStructZero() []ArrStr {
	var arrStruct []ArrStr
	return arrStruct
}

// for循环非迭代的格式
func ArrForIndex() {
	for index := 0; index < 100; index++ {
		fmt.Println(index)
	}
}
