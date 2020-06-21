package goroutine

import (
	"fmt"
	"testing"
)

func TestAction(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
