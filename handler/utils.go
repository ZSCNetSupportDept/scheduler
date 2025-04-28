package handler

import (
	"github.com/golang-module/carbon/v2"
	"zsxyww.com/scheduler/config"
	"zsxyww.com/scheduler/model"
)

// 输入一个时间，返回时间是第几周的第几天
func getWorkDay(in carbon.Carbon) (weekOffset int, dayOffset int) {
	time := carbon.Parse(config.StartTime)
	_weekOffset := time.DiffInWeeks(in)
	_dayOffset := in.DayOfWeek()
	return int(_weekOffset), _dayOffset
}

// 输入成员工号，返回这个成员的信息
func getMemberByID(in int) *model.Member {
	return nil
}
