package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func GetRequest() {

	fileName := "search.log"
	logFile, err := os.Create(fileName)
	defer logFile.Close()

	if err != nil {
		log.Fatalln("open file error !")
	}

	debugLog := log.New(logFile, "[Info]", log.Llongfile)

	data := []string{
		"lfc",
	}
	for _, d := range data {
		_, err := http.Get(fmt.Sprintf("http://host/app/v1/search?key=%s", d))
		if err != nil {
			debugLog.Println(err.Error())
		} else {
			debugLog.Println(fmt.Sprintf("key: %s, ok", d))
		}
	}
}

type A struct {
	Info map[string]interface{} `json:"info"`
}
