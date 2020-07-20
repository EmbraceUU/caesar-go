package main

import (
	"sync"
	"time"
)

// 模式, 工厂方法将goroutine和channel绑定

type receiver struct {
	sync.WaitGroup // wg是为了等待goroutine执行完毕
	data           chan int
}

func getReceiver() *receiver {
	r := receiver{
		data: make(chan int),
	}

	// 添加一个计数器
	r.Add(1)
	// 开启一个新的goroutine
	go func() {
		defer r.Done()
		// 在goroutine中, 处理chan中的数据
		for i := range r.data {
			println("receive i: ", i)
		}
	}()

	return &r
}

func _main() {
	r := getReceiver()

	r.data <- 1
	r.data <- 2

	time.Sleep(time.Second * 4)

	r.data <- 3
	r.data <- 4

	close(r.data)
	r.Wait()
}
