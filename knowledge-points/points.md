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

### 什么是sync.Cond, 为什么可以使用channel作为事件通知, 还需要Cond?

Cond是一个条件锁, 又名条件变量, 或者叫定期唤醒锁

Cond基于互斥锁, 也就是必须有Lock

不是用来保护共享资源的, 而是用来协调协程的, 也就是协调哪个协程可以运行或者共享资源的

代替轮询的去获取锁或者权限, 而是在条件发生改变时, 被动的接受通知, 也就代替了for{}

#### 有三个接口, Wait/Signal/Broadcast

Wait是用来阻塞当前协程

Signal是用来发送至少一个通知

Broadcast是用来广播的

#### 为什么需要加锁 ? 

#### 为什么连续或者多个通知时, channel不合适 ?
当只需要通知单个协程时, channel比较合适, 但是当需要连续多个通知时, 需要根据不同条件创建多个channel, 因对多种通知情况.
这样显得杂乱而且不便于管理. 当然如果一次性通知所有范围内的协程, channel也可以做到.

#### cond在多处通知中的应用

即使是Cond可以满足连续多次通知阻塞的协程, 有时候也需要根据协程本身的条件进行判断, 

所以需要有for{}或者goto来不断的循环判断, 当收到通知时, 可以根据条件判断是否符合执行的要求

> 借鉴ants中的retrieveWorker函数中的应用