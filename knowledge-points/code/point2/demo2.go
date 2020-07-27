package main

import (
	"sync"
	"time"
)

func main() {
	ch := make(chan struct{})
	var wg sync.WaitGroup

	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func(n int) {
			// 阻塞
			<-ch
			println(n)
			wg.Done()
		}(i)
	}

	println("start .")
	time.Sleep(time.Second * 5)
	// 通知所有协程
	close(ch)
	// 等待所有携程运行完毕
	wg.Wait()
	println("done .")
}
