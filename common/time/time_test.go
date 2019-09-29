package time

import (
	"fmt"
	"testing"
)

func TestGetDateNowLoc(t *testing.T) {
	loc := "Asia/Chongqing"
	fmt.Println(GetDateNowLoc(loc))
}

func TestStrToTimestampLoc(t *testing.T) {
	loc := "Asia/Chongqing"
	timeStr := "2019-09-05 14:40:23"
	fmt.Println(StrToTimestampLoc(loc, timeStr))
}

func TestStrToTimestampUTC(t *testing.T) {
	timeStr := "2019-09-05T14:40:23.000Z"
	fmt.Println(StrToTimestampUTC(timeStr))
}
