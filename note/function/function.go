package function

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
