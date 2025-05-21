// CRUD的基础操作
package uo

import (
	"zsxyww.com/scheduler/database"
	"zsxyww.com/scheduler/model"
)

// 增加一项tweak
func AddTweak(in *model.Tweak) error {
	result := db.Main.Create(in)
	return result.Error

}

// 删除一项tweak
func DeleteTweak(in *model.Tweak) error {
	if db.Main.Error != nil {
		return db.Main.Error
	}
	return nil
}

// 查询一些tweak，通过IssueID
func GetTweakByIssueID(in *model.Tweak) (result []*model.Tweak, err error) {
	if db.Main.Error != nil {
		return nil, db.Main.Error
	}
	return nil, nil
}

// 查询一些tweak，通过一个日期
func GetTweakByTime(in *model.Tweak) (result []*model.Tweak, err error) {
	if db.Main.Error != nil {
		return nil, db.Main.Error
	}
	return nil, nil
}

// 查询一些tweak，通过一个工号
func GetTweakByID(in *model.Tweak) (result []*model.Tweak, err error) {
	if db.Main.Error != nil {
		return nil, db.Main.Error
	}
	return nil, nil
}
