package time

import (
	"fmt"
	"time"
)

func aaa() {
	loc, _ := time.LoadLocation("Asia/Chongqing")
	tm2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-09-05 14:40:23", loc)
	fmt.Println(tm2.Unix())
}

func bbb() {
	loc, _ := time.LoadLocation("Asia/Chongqing")
	data := time.Now().In(loc).Format("2006-01-02")
	fmt.Println(data)
}
