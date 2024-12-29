package model

type Member struct {
	ID       int    `csv:"工号"` //工号
	Name     string `csv:"姓名"`
	Sex      bool   `csv:"性别"` //不要把女生安排进男生宿舍，male=True
	FreeDay  int    `csv:"有空"` //哪天有空
	Access   int    `csv:"权限"` //遵循报修系统的access enum ,用来标注管理层
	Arranged bool   `csv:"-"`  //供分配程序使用的字段
	Note     int    `csv:"-"`  //正常=0,换班/补班=1,蹭班=2,供分配程序使用
}

// 权限：

const DEV = 0    //开发组
const HEAD = 1   //科长
const API = 2    //API
const GROUP = 3  //组长
const FORMAL = 4 //正式成员
const FRESH = 5  //实习成员
const PRE = 6    //前成员
