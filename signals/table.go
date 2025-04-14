package signals

import (
	"github.com/golang-module/carbon/v2"
	"sync"
)

// 有关值班表相关的状态
type TablePrototype struct {
	lastUpdated carbon.Carbon //值班表最后更新的时间,不要直接读取这个变量
	LUMutex     sync.RWMutex

	needUpdate bool //值班表是否需要更新，不要直接读取这个变量
	NUMutex    sync.RWMutex
}

var Table TablePrototype

// 下面是操作这些状态的专用函数，用且仅用它们来操作上面的全局变量

// 表格最后的更新时间
func (t *TablePrototype) GetLastUpdated() carbon.Carbon {
	t.LUMutex.RLock()
	defer t.LUMutex.RUnlock()
	return t.lastUpdated

}

// 标记最后更新的时间
func (t *TablePrototype) MarkUpdateTime(newtime carbon.Carbon) error {
	t.LUMutex.Lock()
	defer t.LUMutex.Unlock()
	t.lastUpdated = newtime
	return nil
}

// 表格需要更新吗？
func (t *TablePrototype) IsNeedUpdate() bool {
	t.NUMutex.RLock()
	defer t.NUMutex.RUnlock()
	return t.needUpdate

}

// 标记表格要求更新吗
func (t *TablePrototype) MarkUpdateStatus(i bool) error {
	t.NUMutex.Lock()
	defer t.NUMutex.Unlock()
	t.needUpdate = i
	return nil
}

// 标记表格更新，是这两个函数的wrapper,一般用这个
func (t *TablePrototype) SetUpdated(newtime carbon.Carbon) error {
	err := t.MarkUpdateTime(newtime)
	if err != nil {
		return err
	}
	err = t.MarkUpdateStatus(false)
	if err != nil {
		return err
	}
	return nil

}
