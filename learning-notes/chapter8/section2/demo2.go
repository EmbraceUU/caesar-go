package main

import (
	"runtime"
	"sync"
	"time"
)

// 使用channel实现信号量
// 关键步骤, 设置channel的长度, 使用抢占缓冲区阻塞goroutine
// 可以使用在结合缓冲区限流, 使用一个channel作为缓冲区, 使用另外一个channel作为同一时间并发的流量限制

func __main() {
	// 将最大协程设置为4
	runtime.GOMAXPROCS(4)
	// 设置缓冲区为2的channel
	seg := make(chan struct{}, 2)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			seg <- struct{}{} // 发送信号

			defer func() { <-seg }() // 释放信号

			time.Sleep(time.Second)
			println("seg,", id)
		}(i)
	}

	wg.Wait()
}
