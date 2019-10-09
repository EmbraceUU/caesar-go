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
