package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
	"time"
)

var sum int32

type A struct {
	Num int32
	Ts  int64
}

func myFunc(i interface{}) A {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
	var a A
	a.Num = n
	a.Ts = time.Now().Unix()
	return a
}

func main() {
	runTimes := 10000
	var wg sync.WaitGroup
	numSet := make(map[int32]A)
	mutex := sync.RWMutex{}
	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		a := myFunc(i)
		mutex.Lock()
		numSet[i.(int32)] = a
		mutex.Unlock()
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
		fmt.Println("add task", i)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}
