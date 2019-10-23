package time

import (
	"time"
)

const (
	LayoutStr    = "2006-01-02T15:04:05.000Z"
	LayoutStrSec = "2006-01-02 15:04:05"
	DateFormat   = "2006-01-02"
)

// 字符串转时间戳, 加时区设置
func StrToTimestampLoc(locName, dateStr string) int64 {
	loc, _ := time.LoadLocation(locName)
	tm2, _ := time.ParseInLocation(LayoutStrSec, dateStr, loc)
	return tm2.Unix()
}

// 获取当前日期, 加时区设置
func GetDateNowLoc(locName string) string {
	loc, _ := time.LoadLocation(locName)
	data := time.Now().In(loc).Format(DateFormat)
	return data
}

// UTC字符串转时间戳
func StrToTimestampUTC(dateStr string) int64 {
	t, _ := time.Parse(LayoutStr, dateStr)
	return t.Unix()
}

func BeforeTs(ts, step int64) int64 {
	if ts < step {
		return ts
	}
	return ts - step
}

func NowTs() int {
	return int(time.Now().Unix())
}
