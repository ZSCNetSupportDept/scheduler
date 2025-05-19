package uo

import (
	"zsxyww.com/scheduler/model"
)

// 增加一项tweak
func (uo *uoPrototype) addTweak(in *model.Tweak) error {
	_ = uo.c.Create(in)
	if uo.c.Error != nil {
		return uo.c.Error
	}
	return nil
}

// 删除一项tweak
func (uo *uoPrototype) deleteTweak(in *model.Tweak) error {
	if uo.c.Error != nil {
		return uo.c.Error
	}
	return nil
}

// 查询一些tweak，通过IssueID
func (uo *uoPrototype) getTweakByIssueID(in *model.Tweak) (result []*model.Tweak, err error) {
	if uo.c.Error != nil {
		return nil, uo.c.Error
	}
	return nil, nil
}

// 查询一些tweak，通过一个日期
func (uo *uoPrototype) getTweakByTime(in *model.Tweak) (result []*model.Tweak, err error) {
	if uo.c.Error != nil {
		return nil, uo.c.Error
	}
	return nil, nil
}

// 查询一些tweak，通过一个工号
func (uo *uoPrototype) getTweakByID(in *model.Tweak) (result []*model.Tweak, err error) {
	if uo.c.Error != nil {
		return nil, uo.c.Error
	}
	return nil, nil
}
