package pool

import (
	"fmt"
	"time"
)

func PoolFunc() {
	for {
		poolArr := []int{1, 2, 3, 4}
		// 轮询刷新币对
		for key := range poolArr {
			fmt.Println(fmt.Sprintf("pool func key: %d", key))
			time.Sleep(time.Second * 5)
		}
	}
}
