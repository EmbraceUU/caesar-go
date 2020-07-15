package main

import "time"

/*
在for循环中的select, 如果channel被关闭了会怎么样, 如果只有一个case会怎么样
*/
func main() {
	c := make(chan int, 3)
	go func() {
		for i := 0; i < cap(c); i++ {
			c <- i
			time.Sleep(time.Second)
		}
		// 关闭c
		close(c)
	}()

	for {
		select {
		case n, ok := <-c:
			println("channel c received: ", n, ok)
		default:
			println("default")
		}
		time.Sleep(time.Second)
	}
}
