package main

import (
	"caesar-go/caesartest/bitmex"
	"fmt"
)

func main() {
	//lasa := []string{
	//    "aaa",
	//    "bbb",
	//    "ccc",
	//    "ddd",
	//    "eee",
	//    "fff",
	//    "ggg",
	//    "hhh",
	//    "iii",
	//    "jjj",
	//    "kkk",
	//}
	//for index, l := range lasa {
	//    fmt.Println(l, " item")
	//    index = index + 2
	//}

	//for index := 0; index < len(lasa); index++{
	//    fmt.Println(lasa[index], " item")
	//    index = index + 2
	//}

	//maps := make(map[string][]int)
	//maps["aaa"] = []int{123, 234}
	//fmt.Println(maps)
	intss := make([]int, 2)
	fmt.Println(intss)

	bitmex.UpdateIN(intss, 3)
	fmt.Println(intss)
}
