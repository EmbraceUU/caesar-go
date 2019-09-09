package time

import (
	"time"
)

const (
	LayoutStr  = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
)

// 字符串转时间戳, 加时区设置
// "Asia/Chongqing" "2019-09-05 14:40:23"
func StrToTimestampLoc(locName, dateStr string) int64 {
	loc, _ := time.LoadLocation(locName)
	tm2, _ := time.ParseInLocation(LayoutStr, dateStr, loc)
	return tm2.Unix()
}

// 获取当前日期, 加时区设置
// "Asia/Chongqing"
func GetDateNowLoc(locName string) string {
	loc, _ := time.LoadLocation(locName)
	data := time.Now().In(loc).Format(DateFormat)
	return data
}
