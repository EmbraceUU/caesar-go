package main

import (
	"sync"
	"time"
)

var cond *sync.Cond
var wg sync.WaitGroup

func init() {
	cond = sync.NewCond(&sync.Mutex{})
}

func test(x int) {
	// 需要先加锁
	cond.L.Lock()
	// 阻塞当前协程
	cond.Wait()
	// 执行
	println(x)
	time.Sleep(time.Second)
	wg.Done()
	// 解锁
	cond.L.Unlock()
}

func main() {
	for i := 0; i < 40; i++ {
		wg.Add(1)
		go test(i)
	}

	println("start .")
	time.Sleep(time.Second * 3)
	cond.Signal()
	time.Sleep(time.Second * 3)
	cond.Signal()
	time.Sleep(time.Second * 3)
	println("broad cast .")
	// 发出广播以后, 全部goroutine不再阻塞, 进而主协程执行完毕.
	cond.Broadcast()

	// 等待所有协程都运行完
	wg.Wait()
}
