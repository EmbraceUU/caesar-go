package main

import (
	conf2 "caesar-go/pkg-example/viper/conf"
	"caesar-go/pkg-example/viper/obs"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf2.InitEngine()
	conf := conf2.GetConfig()

	aa := new(obs.A)
	aa.Update(conf)

	bb := new(obs.B)
	bb.Update(conf)

	cc := new(obs.C)
	cc.Update(conf)

	// add objects to observer
	// This objects would be notified when config file changed, and will update automatic.
	conf2.PushObserver(aa)
	conf2.PushObserver(bb)
	conf2.PushObserver(cc)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
