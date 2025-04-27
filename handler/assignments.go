package handler

import (
	//"fmt"
	"errors"
	"github.com/gocarina/gocsv"
	"github.com/golang-module/carbon/v2"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"sync"
	"zsxyww.com/scheduler/config"
	"zsxyww.com/scheduler/model"
	"zsxyww.com/scheduler/signals"
)

var data *[7][]*model.Member
var mutex sync.RWMutex //lock for data
var err error

// /api/getAssignment GET 获取当日值班表，返回html
// 接受参数date,是需要生成值班表的日期
func GetAssignment(i echo.Context) error {

	//如果指定了参数，则生成参数指定的
	if date := i.QueryParam("date"); date != "" {
		mutex.Lock()
		data, err = generateTable(carbon.Parse(date))
		mutex.Unlock()

		if err != nil {
			i.String(http.StatusInternalServerError, err.Error())
			return echo.ErrInternalServerError
		}
		goto render
	}

	//如果没有参数，则生成当前时间
	if (carbon.Now().ToDateString() != signals.Table.GetLastUpdated().ToDateString()) || signals.Table.IsNeedUpdate() == true {

		mutex.Lock()
		data, err = generateTable(carbon.Now())
		mutex.Unlock()

		if err != nil {
			i.String(http.StatusInternalServerError, err.Error())
			return echo.ErrInternalServerError
		}

		//signals.Table.SetUpdated(carbon.Now())
		//测试时注释掉上面的状态更新方便调试
	}
render:
	mutex.RLock()
	i.Render(http.StatusOK, "table.html", data)
	mutex.RUnlock()

	return nil
}

// 根据指定的时间来生成对应的值班表
func generateTable(time carbon.Carbon) (*[7][]*model.Member, error) {
	table := [7][]*model.Member{} //结果放入这里
	members := []*model.Member{}  //包含所有成员信息的切片
	today := []*model.Member{}    //今天值班的人
	female := []*model.Member{}   //今天的女生
	male := []*model.Member{}     //今天的男生
	week, dayOfWeek := getWorkDay(time)
	//检查传入时间有没有问题
	//TODO:这里好像有bug（对日期是否在值班时间内的判断部分）,不过不怎么影响使用
	if (week < 0) || (week > config.Week) {
		return nil, errors.New("Invalid date,the date must lie in our duty period(startTime~startTime+week*7)in config file")
	}

	// 为了实现更换值班的片区，写的一个闭包切片访问器
	iter := func(array []*model.Member, i int) *model.Member {
		return array[(i+week)%len(array)]
	}

	//读取csv文件
	err := readTableData(&members)
	if err != nil {
		return nil, err
	}
	//添加标题
	table[0] = append(table[0], &model.Member{Name: "凤翔"})
	table[1] = append(table[1], &model.Member{Name: "朝晖"})
	table[2] = append(table[2], &model.Member{Name: "香晖AB"})
	table[3] = append(table[3], &model.Member{Name: "香晖CD"})
	table[4] = append(table[4], &model.Member{Name: "东门"})
	table[5] = append(table[5], &model.Member{Name: "北门"})
	table[6] = append(table[6], &model.Member{Name: "歧头"})

	//初始化数据
	for _, i := range members {
		if i.FreeDay == dayOfWeek {
			today = append(today, i)
		}
	}

	for _, i := range today {
		if i.Sex == false {
			female = append(female, i)
		} else if i.Sex == true {
			male = append(male, i)
		}
	}

	//为女生分配负责人
	for i := 0; i < len(female); i++ {
		if a := iter(female, i); a.Access < model.FRESH { //是正式成员
			table[i%4] = append(table[i%4], a) //轮流分配到女生片区
			a.Arranged = true
		}
	}

	//为剩下的片区分配负责人
	for i := 0; i < len(male); i++ {
		if a := iter(male, i); a.Access < model.FRESH { //是正式成员
			table[fewest(table)] = append(table[fewest(table)], a)
			a.Arranged = true
		}
	}

	//分配剩下的所有女生到女生片区
	for i := 0; i < len(female); i++ {
		if a := iter(female, i); a.Arranged != true { //还没有安排
			table[fewestF(table)] = append(table[fewestF(table)], a)
			a.Arranged = true
		}
	}

	//分配剩下的所有男生
	for i := 0; i < len(male); i++ {
		if a := iter(male, i); a.Arranged == false { //还没有安排
			table[fewest(table)] = append(table[fewest(table)], a)
			a.Arranged = true
		}
	}

	return &table, nil
}

// 读取csv文件
func readTableData(m *[]*model.Member) error {
	data, err := os.OpenFile(config.File, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer data.Close()

	err = gocsv.UnmarshalFile(data, m)
	if err != nil {
		return err
	}
	//for index, member := range *m {
	//	fmt.Printf("%v:%v\n", index, member) // for debug concerns
	//}
	return nil
}

// 找出人数最少的片区
func fewest(a [7][]*model.Member) int {
	b := min(len(a[0]), len(a[1]), len(a[2]), len(a[3]), len(a[4]), len(a[5]), len(a[6]))
	for i := range len(a) {
		if b == len(a[i]) {
			return i
		}
	}
	return -1 //error
}

// 找出人数最少的女生片区
func fewestF(a [7][]*model.Member) int {
	b := min(len(a[0]), len(a[1]), len(a[2]), len(a[3]))
	for i := range len(a) - 3 {
		if b == len(a[i]) {
			return i
		}
	}
	return -1 //error
}
