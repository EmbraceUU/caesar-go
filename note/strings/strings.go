package strings

import "strings"

func Replace() string {
	symbol := "XBTUSD"
	root := "XBT"
	return strings.Replace(symbol, root, "", -1)
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
