package strings

import (
	"fmt"
	"strings"
)

func Replace(n int) string {
	symbol := "ETH_USDT"
	root := "."
	return strings.Replace(symbol, root, "_", n)
}

func Contains(s, subs string) bool {
	return strings.Contains(s, subs)
}

func SubStr(items []string) string {
	var subStr string
	for index, item := range items {
		if index == 0 {
			subStr = item
			continue
		}
		subStr = subStr + "_" + item
	}
	return subStr
}

func Append() {
	StrArr := []string{"qwe", "wer", "ert"}
	fmt.Println(StrArr)
	str1 := "qwe"
	str2 := "wer"
	str3 := "ert"
	StrArrP := []string{}
	StrArrP = append(StrArrP, str1)
	StrArrP = append(StrArrP, str2)
	StrArrP = append(StrArrP, str3)
	fmt.Println(StrArrP)

}

func HasPrefix(s, pre string) bool {
	return strings.HasPrefix(s, pre)
}
