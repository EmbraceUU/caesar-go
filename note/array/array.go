package array

import "fmt"

// 定义后的arr的零值是什么
func ArrDefineZero() {
	var arrStr []string
	fmt.Println(arrStr)
	// []
	fmt.Println(len(arrStr))
	// 0
}

// 定义后的arr[结构体]的零值是什么
func ArrDefineStructZero() {
	var arrStruct []ArrStr
	fmt.Println(arrStruct)
	// []
	fmt.Println(len(arrStruct))
	// 0
}
