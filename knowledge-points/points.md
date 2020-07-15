# 记录

### for 循环中的 select, 如果 channel关闭会怎么样, 如果只有一个case会怎么样

测试select中有一个case, 一个default, 3s后关闭channel

```go
package main

import "time"

/*
在for循环中的select, 如果channel被关闭了会怎么样, 如果只有一个case会怎么样
*/
func main() {
	c := make(chan int, 3)
	go func() {
		for i:=0; i<cap(c); i++ {
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
```

```text
default
channel c received:  0 true
channel c received:  1 true
channel c received:  2 true
channel c received:  0 false
channel c received:  0 false
channel c received:  0 false
channel c received:  0 false
...
```

与想象中的不一致, 预期是在c close以后, 会一直走default, 但是结果是一直在走case

因为channel关闭后, 接收方可以读取并且不会引发panic, 只有把channel设置为nil, 接收方才会被阻塞. 

看看下面的例子, 我们在发送方close以后直接将channel改为nil

```gotemplate
func main() {
	c := make(chan int, 3)
	go func() {
		for i:=0; i<cap(c); i++ {
			c <- i
			time.Sleep(time.Second)
		}
		// 关闭c
		close(c)
        // 设置为nil
		c = nil
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
```

```text
default
channel c received:  0 true
channel c received:  1 true
default
default
default
default
default
default
```

运行结果出现了问题, 在接收方接受完数据之前, channel被置为nil. 

所以应该在接收方读取完数据后, 由接收方将channel设置为nil.

在发送方发送完数据后, 由发送方close channel.

如果只有一个case, 会出现deadlock. 

```gotemplate
func main() {
	c := make(chan int, 3)
	go func() {
		for i:=0; i<cap(c); i++ {
			c <- i
			time.Sleep(time.Second)
		}
		// 关闭c
		close(c)

		c = nil
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
```

```text
channel c received:  0 true
channel c received:  1 true
channel c received:  2 true
fatal error: all goroutines are asleep - deadlock!
```

#### 总结

由发送方控制channel的close

由接收方控制channel是否要置为Nil, 或者是否跳出select循环.

经测试, 不管是同步还是异步, close以后均可以继续从channel中读取数据. 