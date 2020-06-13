package channel

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

// 学习笔记 示例1
// 消息传递和事件通知
func Action01() {
	done := make(chan struct{})
	c := make(chan string)

	go func() {
		s := <-c
		println(s)
		close(done)
	}()

	c <- "hi!"
	<-done
}

/*
学习笔记 示例2
异步模式: 缓冲区
缓冲区大小为3
当写入4个时, 会报fatal, all goroutines are asleep
因为发送方会一直阻塞在第四个消息处
*/
func Action02() {
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	println(<-c)
	println(<-c)
}

/*
学习笔记 示例3
channel变量是指针
可以判断相等和nil
*/
func Action03() {
	var a, b = make(chan int, 3), make(chan int)
	var c chan bool

	println(a == b)
	println(c == nil)
	fmt.Printf("%p %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("%p %d\n", b, unsafe.Sizeof(b))
}

/*
学习笔记 示例04
可以用len判断缓冲区数据长度
用cap判断缓冲区大小
同步模式的len, cap为0
*/
func Action04() {
	a, b := make(chan int, 3), make(chan int)

	a <- 1
	a <- 2

	println("a: ", len(a), cap(a))
	println("b: ", len(b), cap(b))
}

/*
学习笔记 示例5
使用ok-idom方式接收数据
*/
func Action05() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for {
			x, ok := <-c
			if !ok {
				return
			}
			println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	close(c)

	<-done
}

/*
学习笔记 示例6
使用range方式接收数据
一定要关闭c通道, 否则goroutine会导致死锁
*/
func Action06() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for x := range c {
			println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	close(c)

	<-done
}

/**
学习笔记 示例7
群体事件通知
一次性通知
使用close

todo sync.Cond
*/
func Action07() {
	var wg sync.WaitGroup
	ready := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(i)
		go func(id int) {
			defer wg.Done()

			println(id, ": ready.")
			<-ready
			println(id, ": running...")
		}(i)
	}

	// 如果没有sleep, ready会提前关闭
	time.Sleep(time.Second)
	println("Ready? GO")

	close(ready)

	wg.Wait()
}

/*
学习笔记 示例8
从已关闭通道读取消息, 会读取缓冲区数据或者零值
*/
func Action08() {
	c := make(chan int, 3)

	c <- 1
	c <- 2

	// 关闭通道
	close(c)

	for i := 0; i <= cap(c); i++ {
		x, ok := <-c
		println(i, ":", ok, x)
	}
}

/*
学习笔记 示例9
单向通道
*/
func Action09() {
	var wg sync.WaitGroup
	wg.Add(2)

	// 强制转换后的收发双方的指针是同一个
	c := make(chan int)
	var recv <-chan int = c
	var send chan<- int = c
	fmt.Printf("recv %p \n", recv)
	fmt.Printf("send %p \n", send)
	fmt.Printf("c %p \n", c)

	// 接收方
	go func() {
		defer wg.Done()

		for x := range recv {
			println(x)
		}
	}()

	// 发送方
	go func() {
		defer wg.Done()
		// 关闭 send和c一样
		defer close(send)

		for i := 0; i < 3; i++ {
			send <- i
		}
	}()

	wg.Wait()
}
