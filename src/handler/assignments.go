package handler

import (
	"errors"
	"net/http"

	"github.com/golang-module/carbon/v2"
	"github.com/labstack/echo/v4"
	"zsxyww.com/scheduler/config"
	"zsxyww.com/scheduler/model"
)

// /api/getAssignment GET 获取当日值班表
// 接受参数date,是需要生成值班表的日期
func GetAssignment(i echo.Context) error {

	//如果没有参数，则生成当前时间
	arg := carbon.Now()
	//如果指定了参数，则生成参数指定的
	if date := i.QueryParam("date"); date != "" {
		arg = carbon.Parse(date)
	}

	data, err := generateTable(arg)

	if err != nil {
		i.String(http.StatusInternalServerError, err.Error())
		return echo.ErrInternalServerError
	}

	i.JSON(200, data)
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
	if (week < 0) || (week > config.Default.Business.Week) {
		return nil, errors.New("日期错误，日期需要在本学期的值班日期内并且格式正确")
	}

	// 切片访问函数，用来实现自动更换值班片区的功能
	iter := func(array []*model.Member, i int) *model.Member {
		return array[(i+week)%len(array)]
	}

	members = model.MemberList

	//添加标题
	table[0] = append(table[0], &model.Member{Name: "凤翔", Access: 7})
	table[1] = append(table[1], &model.Member{Name: "朝晖", Access: 7})
	table[2] = append(table[2], &model.Member{Name: "香晖AB", Access: 7})
	table[3] = append(table[3], &model.Member{Name: "香晖CD", Access: 7})
	table[4] = append(table[4], &model.Member{Name: "东门", Access: 7})
	table[5] = append(table[5], &model.Member{Name: "北门", Access: 7})
	table[6] = append(table[6], &model.Member{Name: "歧头", Access: 7})

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

	//将所有正式女生分配到女生片区
	for i := range female {
		if a := iter(female, i); a.Access < model.FRESH { //是正式成员
			table[i%4] = append(table[i%4], a) //轮流分配到女生片区
			a.Arranged = true
		}
	}

	//将所有正式男生分配到所有片区(优先分配人少的片区)
	for i := range male {
		if a := iter(male, i); a.Access < model.FRESH { //是正式成员
			table[fewest(table)] = append(table[fewest(table)], a)
			a.Arranged = true
		}
	}

	//分配剩下的所有女生到女生片区(优先分配人少的片区)
	for i := range female {
		if a := iter(female, i); a.Arranged != true { //还没有安排
			table[fewestF(table)] = append(table[fewestF(table)], a)
			a.Arranged = true
		}
	}

	//分配剩下的所有男生(优先分配人少的片区)
	for i := range male {
		if a := iter(male, i); a.Arranged == false { //还没有安排
			table[fewest(table)] = append(table[fewest(table)], a)
			a.Arranged = true
		}
	}

	return &table, nil
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
