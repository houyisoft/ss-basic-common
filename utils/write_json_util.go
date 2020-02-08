package utils

import (
	"encoding/json"
	"github.com/kataras/iris"
	"time"
)

type writer struct{}


//单例helper对象
var Writer *writer

func init() {
	Writer = new(writer)
}

type Time time.Time

type jsonTime time.Time

// 等于0就是成功，不等于0就是失败。默认错误编码为-1，可自定义错误编码
const (
	//定义成功的返回字符串
	Const_ok_msg string = "ok"
	//定义成功的返回码
	Const_ok_code int = 0
	//定义错误的默认返回字符串
	Const_err_msg string = "error"
	//定义错误的默认返回码
	Const_err_code int = -1

	timeFormart = "2006-01-02"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (jsonTime jsonTime) MarshalJSON() ([]byte, error) {
	//当返回时间为空时，需特殊处理
	if time.Time(jsonTime).IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + time.Time(jsonTime).Format("2006-01-02 15:04:05") + `"`), nil
}


//失败响应处理
func (this *writer) Error(c iris.Context, message string) {
	diffTime := (c).Values().Get("requestCurrentTime")
	var timeConsumed int64
	if diffTime == nil {
		timeConsumed = 0
	} else {
		//currentTime := time.Now().UnixNano() / 1e3 //微秒
		currentTime := time.Now().Unix()
		//timeConsumed := cast.IntToDateString(currentTime - cast.ToInt64(diffTime))
		timeConsumed = currentTime - diffTime.(int64)
	}
	result := iris.Map{"code": Const_err_code, "msg": message, "timeConsumed": timeConsumed}
	(c).JSON(result)
}

func (this *writer) Success(c iris.Context, data interface{}) {
	diffTime := (c).Values().Get("requestCurrentTime")
	var timeConsumed int64
	if diffTime == nil { //成功应该都有值的，还是判断一下吧。
		timeConsumed = 0
	} else {
		currentTime := time.Now().UnixNano() / 1e3 //微秒
		timeConsumed = currentTime - diffTime.(int64)
	}
	result := iris.Map{"code": Const_ok_code, "msg": Const_ok_msg, "data": data, "timeConsumed": timeConsumed}
	(c).Header("Content-Type", "text/json;charset=utf-8")
	res, _ := json.Marshal(result)
	if len(res) > 4*1024 {
		(c).WriteGzip(res)
	} else {
		(c).Write(res)
	}
}
