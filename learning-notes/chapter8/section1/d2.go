package main

import (
	"sync"
	"time"
)

/*
wait, 多处使用, 都可以收到信号
*/
func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Wait() // 等待归零, 接触阻塞
		println("wait exit.")
	}()

	go func() {
		time.Sleep(time.Second)
		println("done.")
		wg.Done() // 递减计数
	}()

	wg.Wait() // 等待归零, 解除阻塞
	println("main exit.")
}
