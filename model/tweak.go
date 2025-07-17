package model

import (
	"gorm.io/gorm"
	"time"
	"zsxyww.com/scheduler/database"
)

// 这个结构体是供数据库使用的表结构，换班补班蹭班的记录都会以这种方式储存
type Tweak struct {
	gorm.Model
	IssueID   int       //请求的编号,例如一个换班请求会产生两个记录，IssueID相同
	IssueTime time.Time //需要进行操作的时间
	Type      int       //操作类型,具体看下面的常量声明
	SubjectID int       //工号
	Name      string    //名字
}

const (
	OP_SWITCH_ADD = 0 //换班加入值班表
	OP_SWITCH_SUB = 1 //换班移出值班表
	OP_VOLUNTEER  = 2 //蹭班
	OP_REPAY      = 3 //补班
	OP_ADMIN_ADD  = 4
	OP_ADMIN_SUB  = 5
)

// 增加一项tweak
func (t *Tweak) addTweak() error {
	result := db.Main.Create(t)
	return result.Error

}

// 删除一项tweak
func (t *Tweak) deleteTweak() error {
	result := db.Main.Delete(t)
	return result.Error
}

// 查询一些tweak，通过IssueID
func (t *Tweak) getTweakByIssueID() (result []*Tweak, err error) {
	if db.Main.Error != nil {
		return nil, db.Main.Error
	}
	return nil, nil
}

// 查询一些tweak，通过一个日期
func (t *Tweak) getTweakByTime() (result []*Tweak, err error) {
	if db.Main.Error != nil {
		return nil, db.Main.Error
	}
	return nil, nil
}

// 查询一些tweak，通过一个工号
func (t *Tweak) getTweakByID() (result []*Tweak, err error) {
	if db.Main.Error != nil {
		return nil, db.Main.Error
	}
	return nil, nil
}
