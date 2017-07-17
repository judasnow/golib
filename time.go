package golib

import (
	"time"
	"fmt"
)

// 传入一个 time.Time 返回 t 当天的零点
func GetZeroHourOfDay(t time.Time) (time.Time, error) {
	day, err := time.Parse("2006-01-02", t.Format("2006-01-02"))
	if err != nil {
		return time.Time{}, nil
	} else {
		return day, nil
	}
}

// 返回指定区间的时间列表 (会格式化到每天的零点)
func GetDaysList(begin time.Time, end time.Time) []time.Time {
	days := []time.Time{}
	crtDay := begin

	for crtDay.Before(end) {
		crtDayZero, err := GetZeroHourOfDay(crtDay)
		if err != nil {
			crtDayZero = time.Time{}
		}
		days = append(days, crtDayZero)
		crtDay = crtDay.AddDate(0, 0, 1)
	}

	endDayZero, err := GetZeroHourOfDay(crtDay)
	if err != nil {
		endDayZero = time.Time{}
	}

	days = append(days, endDayZero)
	return days
}

// 将指定秒数进格式化处理
func ConvSecondsToHoursMinusSeconds(s int) string {
	var h = s / 3600
	var remainSeconds = s % 3600
	var m = remainSeconds / 60
	var _remainSeconds = remainSeconds % 60

	var hoursStr string
	if h > 0 {
		hoursStr = fmt.Sprintf("%02d:", h)
	} else {
		hoursStr = "00:"
	}

	var mStr string
	if m > 0 {
		mStr = fmt.Sprintf("%02d:", m)
	} else {
		mStr = "00:"
	}

	return fmt.Sprintf("%s%s%02d", hoursStr, mStr, _remainSeconds)
}