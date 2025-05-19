package handler

//
// import (
// 	"github.com/gocarina/gocsv"
// 	"os"
// 	"zsxyww.com/scheduler/config"
// 	//"zsxyww.com/scheduler/database"
// 	"zsxyww.com/scheduler/model"
// )
//
// func init() {
// 	allMember, err := loadMembers()
// 	if err != nil {
// 		panic(err)
// 	}
// 	_ = allMember
// }
//
// func loadMembers() ([]model.Member, error) {
// 	data, err := os.OpenFile(config.File, os.O_RDWR|os.O_CREATE, os.ModePerm)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer data.Close()
//
// 	var m []model.Member
//
// 	err = gocsv.UnmarshalFile(data, m)
// 	if err != nil {
// 		return nil, err
// 	}
// 	//for index, member := range *m {
// 	//	fmt.Printf("%v:%v\n", index, member) // for debug concerns
// 	//}
// 	return m, nil
// }
