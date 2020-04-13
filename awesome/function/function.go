package function

import "sync"

/*
本包包含关于函数调用的笔记和测试
*/

/*
defer为延迟调用, 在return之前调用, 即使有异常抛出, defer也会被调用
多个defer被注册时, 按照FILO的顺序执行
*/
func DeferPractice(x int) {
	defer println("a")
	defer println("b")
	defer func() {
		println(100 / x)
	}()
	defer println("c")
}

/*
延迟调用参数在注册时求值或者复制, 可用指针或者闭包延迟读取
defer中, i的值为注册时x的值, y为当前的值
*/
func DeferPractice1() {
	x, y := 10, 20

	defer func(i int) {
		println("defer: ", i, y)
	}(x)

	x += 10
	y += 100
	println("x=", x, " y=", y)
}

/*
在循环中重复调用Defer, 会导致性能问题
BenchmarkDeferPractice2-4   	30000000	        51.9 ns/op		使用defer
BenchmarkDeferPractice2-4   	100000000	        15.8 ns/op		未使用defer
*/
var lock sync.Mutex

func DeferPractice2() {
	lock.Lock()
	defer lock.Unlock()
}
