package golib

import (
	"time"
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