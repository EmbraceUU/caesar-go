package rate

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func WaitN() {
	limit := rate.NewLimiter(120, 8)
	for {
		if limit.Allow() {
			fmt.Println("log:event happen", time.Now())
		} else {
			fmt.Println("log:event not allow", time.Now())
		}
		time.Sleep(time.Millisecond * 1)
	}
}
