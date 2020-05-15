package _defer

/*
测试一下defer相关的特性
*/

// 测试延时调用的特性
// defer的入参会提前准备好
func DelayCallDefer() {
	var flag bool

	defer func() {
		if flag {
			println("delay call.")
		}
	}()

	flag = true
}
