package rate

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func WaitN() {
	limit := rate.NewLimiter(0.2, 1)
	for {
		if limit.Allow() {
			fmt.Println("log:event happen", time.Now())
		} else {
			fmt.Println("log:event not allow", time.Now())
		}
		time.Sleep(time.Second * 1)
	}
}
