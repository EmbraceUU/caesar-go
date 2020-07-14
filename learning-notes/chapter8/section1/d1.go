package main

import "time"

var c int

func counter() int {
	c++
	return c
}

/*
goroutine 延时执行
*/
func main() {
	a := 100

	// 虽然不会立即执行函数, 但是参数会被立即执行并复制
	go func(x, y int) {
		time.Sleep(time.Second)
		println("go:", x, y)
	}(a, counter())

	a += 100
	println("main:", a, counter())

	time.Sleep(time.Second * 3)
}
