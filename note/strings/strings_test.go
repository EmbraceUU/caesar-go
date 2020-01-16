package strings

import (
	"fmt"
	"strings"
	"testing"
)

func TestReplace(t *testing.T) {
	fmt.Println(Replace(1))
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

func TestAppend(t *testing.T) {
	Append()
}

func TestHasPrefix(t *testing.T) {
	fmt.Println(HasPrefix("34ETH_USDT", "4"))
}

func TestKeyToExAndSymbol(t *testing.T) {
	ex, sy, err := KeyToExAndSymbol("")
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(fmt.Sprintf("%s %s", ex, sy))
	}
}
