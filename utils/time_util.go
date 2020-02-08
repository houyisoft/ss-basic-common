package utils

import (
	"fmt"
	"regexp"
	"ss-basic-common/utils/cast"
	"strings"
	"time"
)

type JsonTime time.Time
type JsonDate time.Time

//实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02"))
	return []byte(stamp), nil
}

func (this JsonDate) MarshalDateJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02"))
	return []byte(stamp), nil
}

func GetNowMonth() string {

	currentTime := time.Now()

	fmt.Println("Current Time in String: ", currentTime.String())

	fmt.Println("MM-DD-YYYY : ", currentTime.Format("02-2006"))
	return currentTime.Format("02-2006")
	//
	//currentTime := time.Now()
	//
	//fmt.Println("Current Time in String: ", currentTime.String())
	//
	//fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))
	//fmt.Println("MM-DD-YYYY : ", currentTime.Format("20060102"))
	//
	//fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))
	//
	//fmt.Println("YYYY.MM.DD : ", currentTime.Format("2006.01.02 15:04:05"))
	//return cast.ToInt(currentTime.Format("20060102"))
}

// 根据日期字符串获取当天第一秒的时间戳
func DateToUnix(date string) int64 {
	startDate := date + " 00:00:00"
	//日期转化为时间戳
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, startDate, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64

	return timestamp
}

func StringToTime(date string) time.Time {
	timeLayout := "2006-01-02"           //转化所需模板
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, date, loc)
	return tmp
}

// 根据日期字符串获取当天最后一秒的时间戳
func DateToLastSecUnix(date string) int64 {
	endDate := date + " 23:59:59"
	//日期转化为时间戳
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, endDate, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64

	return timestamp
}

//日期转换成时间搓
func DateToStartEndUnix(date string) (int64, int64) {
	startDate := date + " 00:00:00"
	endDate := date + " 23:59:59"
	//日期转化为时间戳
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, startDate, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64

	tmp1, _ := time.ParseInLocation(timeLayout, endDate, loc)
	timestamp1 := tmp1.Unix() //转化为时间戳 类型是int64
	return timestamp, timestamp1
}

func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	date := GetFirstDateOfMonth(d).AddDate(0, 1, -1)
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 59, d.Location())
}

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func DateToStartEndMonthUnix(date string) (int64, int64) {
	startDate := date + "-01 00:00:00"
	local, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", startDate, local)
	firstMonth := GetFirstDateOfMonth(t)
	endMonth := GetLastDateOfMonth(t)
	return firstMonth.Unix(), endMonth.Unix()
}

func Yesterday() string {
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, -1)
	return yesTime.Format("2006-01-02")
}

func Today() string {
	nTime := time.Now()
	return nTime.Format("2006-01-02")
}

func LastMonth() string {
	nTime := time.Now()
	yesTime := nTime.AddDate(0, -1, 0)
	return yesTime.Format("2006-01")
}

func HistoryDay(day int) string {
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, day)
	return yesTime.Format("2006-01-02")
}

func GetJsonTimeWithString(date string) JsonTime {
	loc, _ := time.LoadLocation("Local")
	the_time, _ := time.ParseInLocation("2006-01-02", date, loc)
	fmt.Println(the_time.Format("2006-01-02"))
	return JsonTime(the_time)
}

// 根据时间字符串获取时间戳
func DatetimeToUnix(datetime string) int64 {
	//日期转化为时间戳
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64

	return timestamp
}


//获取当前时间转换成int
func GetNowTime() int {
	return int(time.Now().Unix())
}

//获取格式化时间yyyyMMddHHmmssSSS
func GetFmtTime() string {
	return strings.Replace(time.Now().Format("20060102150405.000"), ".", "", 1)
}

//由时间字符串生成时间戳
func GetInt64FromTime(value string) int64 {
	var currentTime int64 = 0
	loc, _ := time.LoadLocation("Asia/Shanghai")
	if v, err := time.ParseInLocation("2006-01-02 15:04:05", value, loc); err == nil {
		currentTime = v.Unix()
	}
	return currentTime
}

func GetIntFromTime(value string) int {
	var currentTime int = 0
	loc, _ := time.LoadLocation("Asia/Shanghai")
	if v, err := time.ParseInLocation("2006-01-02 15:04:05", value, loc); err == nil {
		currentTime = cast.ToInt(v.Unix())
	}
	return currentTime
}

//由时间字符串生成时间戳
func GetInt64FromDate(value string) int64 {
	var currentTime int64 = 0
	loc, _ := time.LoadLocation("Asia/Shanghai")
	if v, err := time.ParseInLocation("2006-01-02", value, loc); err == nil {
		currentTime = v.Unix()
	}
	return currentTime
}

func GetDateTimeFromUnix(unix int64) string {
	timeLayout := "2006-01-02 15:04:05" //待转化为时间戳的字符串
	datetime := time.Unix(unix, 0).Format(timeLayout)
	return datetime
}

// 是否是日期格式
func IsDate(value string) bool {
	matched, err := regexp.MatchString("^\\d{4}\\-\\d{2}\\-\\d{2}$", value)
	return err == nil && matched
}

//是否是日期时间格式
func IsDatetime(value string) bool {
	matched, err := regexp.MatchString("^\\d{4}\\-\\d{2}\\-\\d{2}\\s+\\d{1,2}:\\d{1,2}:\\d{1,2}$", value)
	return err == nil && matched
}

// 计算日期时间范围，例：GetDatetimeRange(0, 1)，返回当天00:00:00至第二天00:00:00的时间戳
func GetDatetimeRange(start int64, length int64) (int64, int64) {
	var timeLoc, _ = time.LoadLocation("Asia/Shanghai")
	nowTime := time.Now().In(timeLoc)
	sFromTime := nowTime.Format("2006-01-02")
	fromTime, _ := time.ParseInLocation("2006-01-02", sFromTime, timeLoc)
	var iDaySec int64 = 3600 * 24
	iFromTime := fromTime.Unix() + iDaySec*start
	iToTime := iFromTime + iDaySec*length
	return iFromTime, iToTime
}
