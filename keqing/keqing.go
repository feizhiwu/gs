package keqing

import (
	"strconv"
	"strings"
	"time"
)

// Timestamp 日期转换时间戳
func Timestamp(date string) string {
	var layout string
	s := strings.Split(date, " ")
	if len(s) > 1 {
		layout = "2006-01-02 15:04:05"
	} else {
		s := strings.Split(date, "-")
		if len(s) > 2 {
			layout = "2006-01-02"
		} else {
			layout = "2006-01"
		}
	}
	location, _ := time.LoadLocation("Local")
	timestamp, _ := time.ParseInLocation(layout, date, location)
	return strconv.FormatInt(timestamp.Unix(), 10)
}

// Format 获取全格式日期
func Format(timestamp interface{}) string {
	return Date("2006-01-02 15:04:05", timestamp)
}

func Time() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// Date 获取格式日期
func Date(format string, timestamp interface{}) string {
	sec := int64(makeInt(timestamp))
	date := time.Unix(sec, 0).Format(format)
	return date
}

func Day(days int) string {
	year, month, day := time.Now().Date()
	thisMonth := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	thisMonth = thisMonth.AddDate(0, 0, days)
	return thisMonth.Format("2006-01-02")
}

func Week(weeks int) string {
	weekDay := int(time.Now().Weekday())
	if weekDay == 0 {
		weekDay = 7
	}
	timestamp := makeInt(Time()) - (weekDay-1-weeks*7)*24*3600
	return Date("2006-01-02", timestamp)
}

// Month 获取指定月第一天日期
func Month(months int) string {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	thisMonth = thisMonth.AddDate(0, months, 0)
	return thisMonth.Format("2006-01-02")
}

func makeInt(num interface{}) int {
	switch num.(type) {
	case int:
		return num.(int)
	case uint:
		return int(num.(uint))
	case float32:
		return int(num.(float32))
	case float64:
		return int(num.(float64))
	case string:
		i, _ := strconv.Atoi(num.(string))
		return i
	default:
		return 0
	}
}
