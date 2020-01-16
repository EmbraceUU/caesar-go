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

func ArrCut(arrStr []string, curIndex int) []string {
	return append(arrStr[:curIndex], arrStr[curIndex+1:]...)
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

func SubStrArr() {
	arr := []int{2, 3, 4, 5, 6, 76, 8}
	arrNew := arr[0:5]
	fmt.Println(arr)
	fmt.Println(arrNew)
	fmt.Println(fmt.Sprintf("arr: %p", arr))
	fmt.Println(fmt.Sprintf("arrNew: %p", arrNew))
}

func NewCurElement(curListOld, curListNew []string) []string {
	result := make([]string, 0, len(curListOld))
	temp := map[string]struct{}{}
	// 先将旧的curList存入 map
	for _, item := range curListOld {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
		}
	}
	// 然后遍历新的,  如果有新的有不存在的key, 加入result
	for _, item := range curListNew {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
