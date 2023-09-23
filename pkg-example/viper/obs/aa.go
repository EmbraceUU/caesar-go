package obs

import (
	"caesar-go/pkg-example/viper/conf"
	"fmt"
)

type A struct {
	ApiFilePath   string
	NodeFilePath  string
	ChainFilePath string
}

func (a *A) Update(config *conf.Config) {
	a.ApiFilePath = config.File.ApiFilePath
	a.ChainFilePath = config.File.ChainFilePath
	a.NodeFilePath = config.File.NodeFilePath

	fmt.Println(a)
}
