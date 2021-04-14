package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {

	c := cron.New(cron.WithParser(cron.NewParser(cron.Second|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.DowOptional|cron.Descriptor)), cron.WithChain())

	id, err := c.AddFunc("@hourly", func() {
		fmt.Println(time.Now())
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(id)
	}

	// 启动
	c.Start()

	// Inspect the cron job entries' next and previous run times.
	fmt.Println(c.Entries())

	defer c.Stop()

	// 阻塞主进程
	select {}
}
