package handler

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/golang-module/carbon/v2"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"zsxyww.com/scheduler/config"
	"zsxyww.com/scheduler/model"
	"zsxyww.com/scheduler/signals"
)

var data *[7][]string
var err error

func GetAssignment(i echo.Context) error {
	if (carbon.Now().ToDateString() != signals.Table.GetLastUpdated().ToDateString()) || signals.Table.IsNeedUpdate() == true {
		fmt.Printf("At %v:start regenerate table", carbon.Now())
		data, err = generateTable()
		if err != nil {
			i.String(http.StatusInternalServerError, err.Error())
			return echo.ErrInternalServerError
		}
		//signals.Table.SetUpdated(carbon.Now())
		//测试时注释掉上面的状态更新方便调试
	}
	i.Render(http.StatusOK, "table.html", data)
	return nil
}
func generateTable() (*[7][]string, error) {

	table := [7][]string{}       //结果放入这里
	members := []*model.Member{} //包含所有成员信息的切片

	err := readTableData(&members)
	if err != nil {
		return nil, err
	}

	dayOfWeek := carbon.Now().DayOfWeek() //今天星期几
	today := []*model.Member{}            //今天值班的人
	female := []*model.Member{}           //今天的女生
	male := []*model.Member{}             //今天的男生

	table[0] = append(table[0], "凤翔")
	table[1] = append(table[1], "朝晖")
	table[2] = append(table[2], "香晖AB")
	table[3] = append(table[3], "香晖CD")
	table[4] = append(table[4], "东门")
	table[5] = append(table[5], "北门")
	table[6] = append(table[6], "歧头")

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
	for c, i := range female {
		if i.Access < model.FRESH { //是正式成员
			table[c%4] = append(table[c%4], i.Name) //轮流分配到女生片区
			i.Arranged = true
		}

	}

	//为剩下的片区分配负责人
	for _, i := range male {
		if i.Access < model.FRESH {
			table[fewest(table)] = append(table[fewest(table)], i.Name)
			i.Arranged = true
		}
	}

	//分配剩下的所有女生到女生片区
	for _, i := range female {
		if i.Arranged != true {
			table[fewestF(table)] = append(table[fewestF(table)], i.Name)
			i.Arranged = true
		}
	}

	//分配剩下的所有男生
	for _, i := range male {
		if i.Arranged == false {
			table[fewest(table)] = append(table[fewest(table)], i.Name)
		}
	}
	fmt.Printf("today:%v\n", today)
	fmt.Printf("table:%v\n", table)

	//测试的时候先注释掉这里
	//signals.Table.LastUpdated = carbon.Now()
	return &table, nil
}
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
	for index, member := range *m {
		fmt.Printf("%v:%v\n", index, member) // for debug concerns
	}
	return nil
}

// 找出人数最少的片区
func fewest(a [7][]string) int {
	b := min(len(a[0]), len(a[1]), len(a[2]), len(a[3]), len(a[4]), len(a[5]), len(a[6]))
	for i := range len(a) {
		if b == len(a[i]) {
			return i
		}
	}
	return -1 //error
}

// 找出人数最少的女生片区
func fewestF(a [7][]string) int {
	b := min(len(a[0]), len(a[1]), len(a[2]), len(a[3]))
	for i := range len(a) - 3 {
		if b == len(a[i]) {
			return i
		}
	}
	return -1 //error
}
