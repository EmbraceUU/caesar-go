package time

import (
	"log"
	"math"
	"time"
)

const (
	LayoutStr    = "2006-01-02T15:04:05.000Z"
	LayoutStrSec = "2006-01-02 15:04:05"
	DateFormat   = "2006-01-02"
)

type Segment struct {
	Start int
	End   int
}

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
	d := dateStr + "T08:00:00.000Z"
	t, err := time.Parse(LayoutStr, d)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return t.Unix()
}

func SplitTimelineTimestamp(startTs, endTs, step, nowTs, splitSize int) (head, body *Segment) {
	sgStart := CeilNthTimestamp(startTs, step)
	sgEnd := CeilNthTimestamp(endTs, step)
	mid := NthTimestamp(nowTs, step) - splitSize

	start := NthTimestamp(sgStart, step)
	end := NthTimestamp(sgEnd, step)
	sgMid := UnnthTimestamp(mid, step)

	head, body = splitHeadBody(start, end, mid, sgStart, sgEnd, sgMid)
	return
}

func splitHeadBody(start, end, mid, sgStart, sgEnd, sgMid int) (head, body *Segment) {
	if mid <= start {
		head = &Segment{sgStart, sgEnd}
		return
	}

	if mid >= end {
		body = &Segment{sgStart, sgEnd}
		return
	}

	head = &Segment{sgMid, sgEnd}
	body = &Segment{sgStart, sgMid}
	return
}

func NthTimestamp(ts, step int) int {
	return int(math.Floor(float64(ts) / float64(step)))
}

func NthWeek(ts int) int {
	// 1970年1月1日星期四
	dayOffset := 3 * 24 * 60 * 60
	weekStep := 7 * 24 * 60 * 60

	return (ts-dayOffset)/weekStep + 1
}

func NthMonth(ts int) int {
	year, month, _ := time.Unix(int64(ts), 0).UTC().Date()
	return (year-1970)*12 + int(month) - 1
}

func NthYear(ts int) int {
	tm := time.Unix(int64(ts), 0).UTC()
	return tm.Year() - 1970
}

func UnnthTimestamp(nth, step int) int {
	return nth * step
}

func UnnthWeek(nth int) int {
	// 1970年1月1日星期四
	dayOffset := 3 * 24 * 60 * 60
	weekStep := 7 * 24 * 60 * 60

	return (nth-1)*weekStep + dayOffset
}

func UnnthMonth(nth int) int {
	year := nth/12 + 1970
	month := nth%12 + 1
	return int(time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC).Unix())
}

func UnnthYear(nth int) int {
	year := nth + 1970
	return int(time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).Unix())
}

func FloorNthTimestamp(ts, step int) int {
	return UnnthTimestamp(NthTimestamp(ts, step), step)
}

func FloorNthWeek(ts int) int {
	return UnnthWeek(NthWeek(ts))
}

func FloorNthMonth(ts int) int {
	return UnnthMonth(NthMonth(ts))
}

func FloorNthYear(ts int) int {
	return UnnthYear(NthYear(ts))
}

func CeilNthTimestamp(ts, step int) int {
	nth := NthTimestamp(ts, step)
	res := UnnthTimestamp(nth, step)
	if res != ts {
		res += step
	}
	return res
}

func CeilNthWeek(ts int) int {
	nth := NthWeek(ts)
	res := UnnthWeek(nth)
	if res != ts {
		res = UnnthWeek(nth + 1)
	}
	return res
}

func CeilNthMonth(ts int) int {
	nth := NthMonth(ts)
	res := UnnthMonth(nth)
	if res != ts {
		res = UnnthMonth(nth + 1)
	}
	return res
}

func CeilNthYear(ts int) int {
	nth := NthYear(ts)
	res := UnnthYear(nth)
	if res != ts {
		res = UnnthYear(nth + 1)
	}
	return res
}

func ValidYearPoint(ts int) bool {
	tm := time.Unix(int64(ts), 0).UTC()
	year := tm.Year()
	return time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).Unix() == int64(ts)
}

func ValidMonthPoint(ts int) bool {
	tm := time.Unix(int64(ts), 0).UTC()
	year, month, _ := tm.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Unix() == int64(ts)
}

func ValidWeekPoint(ts int) bool {
	// 1970年1月1日星期四
	dayOffset := 3 * 24 * 60 * 60
	weekStep := 7 * 24 * 60 * 60

	pt := ts - dayOffset
	return pt%weekStep == 0
}

func ValidTimestampPoint(ts, step int) bool {
	return ts%step == 0
}

func NowTs() int {
	return int(time.Now().Unix())
}

func BeforeTs(ts, step int) int {
	if ts < step {
		return ts
	}
	return ts - step
}

func UTCStringToTime(ts string) time.Time {
	timestamp, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		log.Println("parse Timestamp err", ts, err)
		timestamp = time.Now()
	} else {
		timestamp.In(time.Local)
	}
	return timestamp
}
