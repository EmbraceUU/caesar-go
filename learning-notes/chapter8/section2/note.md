# 并发 - 通道

## 整理

### channel 是什么

* 底层是一个队列

* 以通讯代替内存共享, 达到并发安全

* 同步模式: 配对成功 -> 复制数据; 不成功 -> 等待

* 异步模式: 有缓冲区, 

    * 发送方, 有空间 -> 写, 没空间 -> 等待
    
    * 接收方, 有数据 -> 读, 没数据 -> 等待
    
* 可以用 <- 的方式接受, 也可以用 ok-idom 的方式, 或者 range 的方式
    
### 应用场景

* 传递数据, 通讯的方式达到并发安全

* 一次性事件通知用close, 连续多样性的通知, 使用sync.Cond

* ID generator

* Pool

* 信号量 semaphore

* time的 timeout和tick功能

* 捕获INT/TERM信号, 实现atexit函数

* 缓冲区, 防止一次性操作大量I/O或者计算

### 注意事项

* channel 变量本身是指针, 可以判等和判nil

* 只能在发送方close

* 建议在接收方置nil

* close后, 发送方继续发送会引发panic

* close后, 接收方接续接受会读取缓冲区或者零值

* 置nil后, 收发双方都会阻塞

* select 处理多个channel

* select中, close channel会一直被选中

* select中, 置nil channel不会被选中

* select中, 添加default避免阻塞

* range中, close channel会退出遍历

