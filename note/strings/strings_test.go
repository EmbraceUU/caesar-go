package strings

import (
	"fmt"
	"strings"
	"testing"
)

func TestReplace(t *testing.T) {
	fmt.Println(Replace())
}

func TestContains(t *testing.T) {
	str := "write tcp 172.31.70.183:42074->104.18.217.39:443: write: broken pipe"
	subStr := "broken pipe"
	fmt.Println(Contains(str, subStr))
}

func TestSubStr(t *testing.T) {
	str := "okex_btc_usd_swap"
	strItem := strings.Split(str, "_")
	fmt.Println(SubStr(strItem[1:]))
}
