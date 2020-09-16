package sync

import (
	"fmt"
	"sync"
	"time"
)

type Run struct {
	runWG *sync.WaitGroup
}

func RunWGFunc() {
	var run Run
	run.runWG = new(sync.WaitGroup)

	defer func() {
		run.runWG.Wait()
		fmt.Println("wait done. ")
	}()

	run.runWG.Add(1)
	go func() {
		time.Sleep(time.Second * 4)
		fmt.Println("done. ")
		run.runWG.Done()
	}()
}
