package main

import (
	"fmt"
	"sync"
)

/*
Local Storage 问题
*/
func main() {
	var wg sync.WaitGroup
	// 使用了一个外界的容器, 然后在goroutine中进行访问和修改
	var gs [5]struct { // 类似与局部存储 TLS
		id     int // 编号
		result int // 返回值
	}

	for i := 0; i < len(gs); i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			gs[id].id = id
			gs[id].result = (id + 1) * 100
		}(i) // 使用参数, 避免延迟求值
	}

	wg.Wait()
	fmt.Printf("%+v", gs)
}
