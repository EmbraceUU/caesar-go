package obs

import (
	"caesar-go/pkg-example/viper/conf"
	"fmt"
)

type C struct {
	Trace   bool
	Proxy   string
	Timeout int64
}

func (c *C) Update(config *conf.Config) {
	c.Trace = config.HttpClient.Trace
	c.Proxy = config.HttpClient.Proxy
	c.Timeout = config.HttpClient.Timeout

	fmt.Println(c)
}
