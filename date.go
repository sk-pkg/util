package util

import (
	"strings"
	"time"
)

const (
	DefaultTimeFormat = "2006-01-02 15:04:05"
)

func FormatTime(t time.Time, f ...string) string {
	format := DefaultTimeFormat
	if len(f) > 0 {
		format = f[0]
	}

	return t.Format(format)
}

// GetDateStartAndEnd 获取指定时间段的起止时间（string）
// 可选：
// today 今天
// yesterday 昨天
// lately7 最近7天
// lately30 最近30天
// month 本月
// year 本年
func GetDateStartAndEnd(dateString string) [2]string {
	var date = [2]string{"", ""}

	Year, Month, Day := time.Now().Date()
	todayStart := time.Date(Year, Month, Day, 0, 0, 0, 0, time.Local)
	todayEnd := time.Date(Year, Month, Day, 23, 59, 59, 0, time.Local)

	switch dateString {
	case "today":
		date[0] = todayStart.Format("2006-01-02 15:04:05")
		date[1] = todayEnd.Format("2006-01-02 15:04:05")
	case "yesterday":
		date[0] = todayStart.AddDate(0, 0, -1).Format("2006-01-02 15:04:05")
		date[1] = todayEnd.AddDate(0, 0, -1).Format("2006-01-02 15:04:05")
	case "lately7":
		date[0] = todayStart.AddDate(0, 0, -7).Format("2006-01-02 15:04:05")
		date[1] = time.Now().Format("2006-01-02 15:04:05")
	case "lately30":
		date[0] = todayStart.AddDate(0, 0, -30).Format("2006-01-02 15:04:05")
		date[1] = time.Now().Format("2006-01-02 15:04:05")
	case "month":
		date[0] = GetFirstDateOfMonth(time.Now()).Format("2006-01-02 15:04:05")
		date[1] = GetLastDateOfMonth(time.Now()).Format("2006-01-02 15:04:05")
	case "year":
		date[0] = GetFirstDateOfYear(time.Now()).Format("2006-01-02 15:04:05")
		date[1] = GetLastDateOfYear(time.Now()).Format("2006-01-02 15:04:05")
	default:
		dateSlice := strings.Split(dateString, "-")
		loc, _ := time.LoadLocation("Local")

		start, _ := time.ParseInLocation("2006/01/02", dateSlice[0], loc)
		end, _ := time.ParseInLocation("2006/01/02", dateSlice[1], loc)

		date[0] = time.Date(start.Year(), start.Month(), start.Day(), 00, 00, 00, 0, end.Location()).Format("2006-01-02 15:04:05")
		date[1] = time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 0, end.Location()).Format("2006-01-02 15:04:05")
	}

	return date
}

// GetFirstDateOfYear 获取传入的时间所在年份的第一天0点时间。
func GetFirstDateOfYear(d time.Time) time.Time {
	return time.Date(d.Year(), 1, 1, 0, 0, 0, 0, d.Location())
}

// GetLastDateOfYear 获取传入的时间所在年份的最后一天最晚点时间。
func GetLastDateOfYear(d time.Time) time.Time {
	return time.Date(d.Year(), 12, 31, 23, 59, 59, 0, d.Location())
}

// GetFirstDateOfMonth 获取传入的时间所在月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// GetLastDateOfMonth 获取传入的时间所在月份的最后一天最晚点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	d = GetFirstDateOfMonth(d).AddDate(0, 1, -1)
	return GetLastTime(d)
}

// GetZeroTime 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// GetLastTime 获取某一天的最晚点时间。
func GetLastTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
}

// DealTimeFormat 格式化时间
func DealTimeFormat(date, format string) string {
	loc, _ := time.LoadLocation("Local") // 设置时区
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc)
	timeUnix := tt.Unix()
	timeFormat := time.Unix(timeUnix, 0).Format(format)

	return timeFormat
}

// DealDateUnix 日期转化为时间戳
func DealDateUnix(date string) int64 {

	timeLayout := "2006-01-02 15:04:05"  // 转化所需模板
	loc, _ := time.LoadLocation("Local") // 获取时区
	tmp, _ := time.ParseInLocation(timeLayout, date, loc)
	timestamp := tmp.Unix()

	return timestamp
}

// GetNowToDayEndSecond 获取当前时间到当天23:59:59的剩余时间（秒）
func GetNowToDayEndSecond() int {
	now := time.Now()
	dayEnd := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())

	return int(dayEnd.Sub(now).Seconds())
}
