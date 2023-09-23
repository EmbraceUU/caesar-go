package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	HttpClient   HttpClient
	Postgres     Postgres
	BridgePoolI  BridgePool
	BridgePoolII BridgePool
	APIServer    APIServer
	File         File
}

type File struct {
	ApiFilePath   string
	NodeFilePath  string
	ChainFilePath string
}

type APIServer struct {
	Port         string
	ReadTimeOut  uint8
	WriteTimeOut uint8
	Proxy        string
}

type Postgres struct {
	Host     string
	Port     string
	UserName string
	Password string
	DataBase string
}

type HttpClient struct {
	Trace   bool
	Proxy   string
	Timeout int64
}

type BridgePool struct {
	Concurrence     int
	ExpiredDuration int
}

func initConfig(updated chan struct{}) *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath(".")
	v.AddConfigPath("conf")
	v.AddConfigPath("pkg-example/viper/conf")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("init config failed, %s", err.Error())
	}

	config := new(Config)
	err = v.Unmarshal(config)
	if err != nil {
		log.Fatalf("init config failed, %s", err.Error())
	}

	go loadDynamically(v, updated, config)

	return config
}

func loadDynamically(v *viper.Viper, updated chan struct{}, config *Config) {
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config changed. ")
		err := v.Unmarshal(config)
		if err != nil {
			log.Printf("unmarshal config failed, %s. ", err.Error())
			return
		}

		updated <- struct{}{}
	})
	v.WatchConfig()
}
