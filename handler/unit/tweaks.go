package handler

import (
	"zsxyww.com/scheduler/model"
)

// 增加一项tweak
func (uo *uoPrototype) addTweak(in *model.Tweak) error {
	return nil
}

// 删除一项tweak
func (uo *uoPrototype) deleteTweak(in *model.Tweak) error {
	return nil
}

// 查询一些tweak，通过IssueID
func (uo *uoPrototype) getTweakByIssueID(in *model.Tweak) (result []*model.Tweak, err error) {
	return nil, nil
}

// 查询一些tweak，通过一个日期
func (uo *uoPrototype) getTweakByTime(in *model.Tweak) (result []*model.Tweak, err error) {
	return nil, nil
}

// 查询一些tweak，通过一个工号
func (uo *uoPrototype) getTweakByID(in *model.Tweak) (result []*model.Tweak, err error) {
	return nil, nil
}
