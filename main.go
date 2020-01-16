package main

import (
	"fmt"
	"strings"
)

func main() {
	/*key := "btcaaa"
	match := "bta"
	match1 := "btc"
	match2 := "b"
	if checkFuzzy(match, key) || checkFuzzy(match1, key) || checkFuzzy(match2, key) {
		fmt.Println(true)
	}else{
		fmt.Println(false)
	}*/

	/*test1 := make([]string, 0, 100)
	resData := make([]string, 0, 200)
	resData = append(resData, test1...)
	fmt.Println(resData, len(resData))*/

	//key := "中国 BTC"
	//fmt.Println(strings.ToLower(key))
	/*channels := make(chan string, 100)
	go handlerChan(channels)

	for {
		time.Sleep(time.Second * 2)
		channels <- "ni hao"
	}*/

	key := "中国 btc"
	data := "中国 "
	fmt.Println(strings.Contains(key, data))
}

func handlerChan(channel chan string) {
	for {
		data := <-channel
		fmt.Println(data)
	}
}

func checkFuzzy(key, data string) bool {
	if strings.HasPrefix(data, key) {
		fmt.Println(true)
		return true
	}
	fmt.Println(false)
	return false
}
