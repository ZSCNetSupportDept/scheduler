package signals

import (
	"sync"
	//"zsxyww.com/scheduler/database"
)

var IssueID int
var IssueIDMutex sync.Mutex

func init() {
	//启动程序时从数据库获取最后的IssueID用来初始化程序的IssueID变量～
}

// IssueID++
func IssueIDPP() {

	IssueIDMutex.Lock()
	IssueID++
	IssueIDMutex.Unlock()

}

func GetIssueID() int {
	IssueIDMutex.Lock()
	defer IssueIDMutex.Unlock()
	return IssueID
}
