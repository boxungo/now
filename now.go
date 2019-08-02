package now

import "time"

// Timestamp 时间戳,秒
func Timestamp() int64 {
	return time.Now().Local().Unix()
}

// Haosecond 豪秒
func Haosecond() int64 {
	return Nanosecond() / 1000000
}

// Microsecond 微秒
func Microsecond() int64 {
	return Nanosecond() / 1000
}

// Nanosecond 纳秒
func Nanosecond() int64 {
	return time.Now().Local().UnixNano()
}

// Time time.Time
func Time() time.Time {
	return time.Now().Local()
}

// Date 日期, 年-月-日
func Date() string {
	return time.Now().Local().Format("2006-01-02")
}

// Datetime 日期时间, 年-月-日 时:分:秒
func Datetime() string {
	return time.Now().Local().Format("2006-01-02 15:04:05")
}

// String 日期时间.微秒
func String() string {
	return time.Now().Local().Format("2006-01-02 15:04:05.000000")
}
