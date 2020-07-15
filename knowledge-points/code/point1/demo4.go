package main

import "time"

/*
如果是同步类型的channel, 在close以后, 会发生什么
*/
func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
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
			if !ok {
				c = nil
			}
		}
		time.Sleep(time.Second)
	}
}
