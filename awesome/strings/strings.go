package strings

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
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

func KeyToExAndSymbol(key string) (string, string, error) {
	if key == "" {
		return key, key, errors.New("key to exchange and symbol is nil")
	}
	items := strings.Split(key, "_")
	if len(items) < 3 {
		return "", "", errors.New(fmt.Sprintf("key to exchange and symbol is err: %s", key))
	}
	exchange := items[0]
	symbol := strings.Replace(key, fmt.Sprintf("%s_", exchange), "", 1)
	return exchange, symbol, nil
}

func IsHan(key string) bool {
	for _, r := range key {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

func HasDigit(key string) bool {
	matchDigit, _ := regexp.Match(`[\d]`, []byte(key))
	return matchDigit
}
