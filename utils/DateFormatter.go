package utils

import (
	"errors"
	"math/rand"
	"time"
)

// 随机生成大写字母
func GetRandomString(l int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

const (
	LongDateFormat  = "2006-01-02 15:04:05"
	ShortDateFormat = "2006-01-02"
)

//获取日期格式
func GetLongDateString(date string, Hours int64) (dateString string, err error) {
	if len(date) <= 0 {
		return "", errors.New("时间为空")
	}
	inputDate, err := time.Parse(ShortDateFormat, date)
	if err == nil {
		h, _ := time.ParseDuration("1h")
		d := inputDate.Add(time.Duration(Hours) * h)
		return d.Format(LongDateFormat), err
	} else {
		return "", errors.New("时间格式不对")
	}
}

//获取相差时间
func GetMinuteDiffer(start_time, end_time string) int64 {
	var hour int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 60
		return hour
	} else {
		return hour
	}
}
