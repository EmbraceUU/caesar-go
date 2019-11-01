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

var (
	constMap = make(map[string][]ArrStr)
)

// 测试切片的指针
func MapArrStructPoint() {
	arrMap := make(map[string][]ArrStr)
	arrMap["k1"] = []ArrStr{
		{
			String1: "s11",
			String2: "s12",
		},
		{
			String1: "s13",
			String2: "s14",
		},
	}
	arrMap["k2"] = []ArrStr{
		{
			String1: "s21",
			String2: "s22",
		},
		{
			String1: "s23",
			String2: "s24",
		},
	}
	a1 := arrMap["k1"]
	a1 = append(a1, ArrStr{
		String1: "s15",
		String2: "s16",
	})

	constMap["k3"] = []ArrStr{
		{
			String1: "s31",
			String2: "s32",
		},
		{
			String1: "s33",
			String2: "s34",
		},
	}

	constMap = arrMap
	fmt.Println(constMap)
	fmt.Println(arrMap)
	fmt.Println(a1)
}
