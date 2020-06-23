package geek_learning

import "unsafe"

func UnAddressable() {
	const num = 123
	_ = num
	//println(&num) 常量不可寻址
	//println(&(123))	基本类型的字面量不可寻址

	var str = "123"
	println(&str)
	//println(&(str[0]))	字符串的索引结果值
	//println(&(str[0:2])) 	字符串的切片结果值

	str2 := str[0]
	println(&str2)

	//println(&(123+234))	算术操作的结果值

	num2 := 456
	_ = num2
	//println(&(num + num2))	算术操作的结果值

	//_ = &([3]int{1, 2, 3}[0]) 	数组字面量的索引结果值
	//_ = &([3]int{1, 2, 3}[0:2])		数字字面量的切片结果值
	_ = &([]int{1, 2, 3}[0]) // 切片字面量的索引结果值
	//_ = &([]int{1, 2, 3}[0:2]) // 切片字面量的切片结果值
	//_ = &(map[int]string{1: "a"}[0]) 字典字面量的索引结果值

	var map1 = map[int]string{1: "a", 2: "b", 3: "c"}
	_ = map1
	//_ = &(map1[1])	对字典变量的索引结果值

	//_ = &(func(x, y int) {return x + y})	字面量代表的函数
	//_ = &(fmt.Sprintf)  标识符代表的函数
	//_ = &(fmt.Sprintf("asb"))  对函数的调用结果值

	dog := Dog{"little pig"}
	_ = dog
	//_ = &(dog.Name)	标识符代表的函数
	//_ = &(dog.Name())	对方法的调用结果值

	//_ = &(Dog{"little pig"}.name)	结构体字面量的字段
	//_ = &(interface{}(dog))	类型转换表达式的结果值
	dogI := interface{}(dog)
	_ = dogI
	//_ = &(dogI.(Named))		类型断言表达式的结果值
	named := dogI.(Named)
	_ = named
	//_ = &(named.(Dog))	类型断言表达式的结果值

	var chan1 = make(chan int, 1)
	chan1 <- 1
	//_ = &(<-chan1)	接收表达式结果值
}

/*
这种方式可以在不知道dog类型以及拿不到dogP的情况下, 直接读取nameP以及操作nameP
这种方式虽然在使用中没有好处, 但是namePtr如果被人拿到了, 会发生灾难性的危险
*/
func UnSafePointer() {
	// 声明一个dog变量
	dog := Dog{"little pig"}
	// 通过&符号获取dog的指针值
	dogP := &dog
	// 将dogP转换成unsafe.Pointer类型的值, 然后转换成uintptr, 将值赋给dogPtr
	dogPtr := uintptr(unsafe.Pointer(dogP))
	println(dogPtr)

	// unsafe.Offsetof()传入选择表达式, 可以返回结构体值和字段值在内存中的存储地址之间的偏移量
	namePtr := dogPtr + unsafe.Offsetof(dogP.name)
	// 将uintptr类型的值转换成unsafe.Pointer类型, 再转成指针值nameP
	nameP := (*string)(unsafe.Pointer(namePtr))
	// 通过*获取到nameP指向的值
	println(*nameP)
}
