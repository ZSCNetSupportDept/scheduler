package model

type Member struct {
	ID      int //工号
	Name    string
	Sex     bool //不要把女生安排进男生宿舍，male=True
	FreeDay int  //哪天有空
	access  int  // 遵循报修系统的access enum ,用来标注管理层
}
