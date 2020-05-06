package time

import (
	"fmt"
	"testing"
	"time"
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
	timeStr := "20191030"
	fmt.Println(StrToTimestampUTC(timeStr))
}

func TestNowTs(t *testing.T) {
	//fmt.Println(NowTs())
	_, month, day := time.Unix(int64(1577448000), 0).UTC().Date()
	fmt.Println(int(month), day)
}

func TestSplitTimelineTimestamp(t *testing.T) {
	SplitTimelineTimestamp(1572065533, 1572077520, 60, 1572078667, 2)
}

func TestUTCStringToTime(t *testing.T) {
	data := UTCStringToTime("2020-03-23T04:55:43.109Z").Unix() * 1000
	println(data)
}

func TestCron(t *testing.T) {
	Cron()
}

func TestAlias(t *testing.T) {
	a := Alias("month", 1577870833)
	println(a)
}
