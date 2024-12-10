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

func GetAssignment(i echo.Context) error {
	if (carbon.Now().ToDateString() != signals.Table.LastUpdated.ToDateString()) || signals.Table.NeedUpdate == true {
		fmt.Printf("At %v:start regenerate table", carbon.Now())
		err := generateTable()
		if err != nil {
			i.String(http.StatusInternalServerError, err.Error())
			return echo.ErrInternalServerError
		}
		return nil
	}
	return nil
}
func generateTable() error {

	table := [7][]string{}       //结果放入这里
	members := []*model.Member{} //包含所有成员信息的切片

	err := readTableData(&members)
	if err != nil {
		return err
	}

	dayOfWeek := carbon.Now().DayOfWeek() //今天星期几
	today := []*model.Member{}            //今天值班的人
	female := []*model.Member{}           //今天的女生
	male := []*model.Member{}             //今天的男生

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
	//为女生宿舍分配负责人：
	for i := 0; i < 4; i++ {
		for _, x := range female {
			if (x.Arranged == false) && (x.Access < model.FRESH) {
				table[i] = append(table[i], x.Name)
				x.Arranged = true
			}
		}
	}
	//男生负责人：
	for i := 0; i < 7; i++ {
		for _, x := range male {
			if (x.Arranged == false) && (x.Access < model.FRESH) {
				table[i] = append(table[i], x.Name)
				x.Arranged = true
			}
		}
	}
	//女生成员
	for c, i := range female {
		if i.Arranged == false {
			table[(c % 4)] = append(table[(c%4)], i.Name)
		}
	}
	//男生成员
	for c, i := range male {
		if i.Arranged == false {
			table[(c % 7)] = append(table[(c%7)], i.Name)
		}
	}
	fmt.Printf("today:%v\n", today)
	fmt.Printf("table:%v\n", table)
	return nil
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
