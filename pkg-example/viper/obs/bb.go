package obs

import (
	"caesar-go/pkg-example/viper/conf"
	"fmt"
)

type B struct {
	Port         string
	ReadTimeOut  uint8
	WriteTimeOut uint8
	Proxy        string
}

func (b *B) Update(config *conf.Config) {
	b.Port = config.APIServer.Port
	b.ReadTimeOut = config.APIServer.ReadTimeOut
	b.WriteTimeOut = config.APIServer.WriteTimeOut
	b.Proxy = config.APIServer.Proxy

	fmt.Println(b)
}
