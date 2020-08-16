package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Println("hello golang! ")

	// 增加信号监听, 这里是监听了所有信号, 只要有信号就退出
	quit := make(chan os.Signal)
	signal.Notify(quit)
	<-quit
	println("quit...")
	time.Sleep(time.Second * 10)

}
