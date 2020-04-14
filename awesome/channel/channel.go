package channel

import "time"

var (
	WaitRoutineChan chan struct{}
	WaitMainChan    chan struct{}
)

/*
竞争资源的保险做法, 在异步协程中, 加入信号, 等待信号出现, 再往下进行
*/
func WaitLoad() {
	WaitRoutineChan = make(chan struct{}, 1)
	WaitMainChan = make(chan struct{}, 1)

	go func() {
		<-WaitRoutineChan
		println("go routine 2")
		WaitMainChan <- struct{}{}
	}()

	go func() {
		time.Sleep(time.Second * 10)
		WaitRoutineChan <- struct{}{}
		println("go routine 1")
	}()

	<-WaitMainChan
	println("main routine")
}
