# 设计
本系统计划完成以下业务：
## 值班表生成
提供所有本学期应值班成员的表格，系统自动生成每一天的值班名单，其中：
- 区分片区负责人，正式成员，实习成员
- 标记出现在表中的管理层
- 每隔一段时间轮换片区
- 女生不在男生片区值班
- 标记值班备注（日常，换蹭补）
## 换班处理
发起人将通过web API发起换班请求，包含换班日期（需要符合值班的日期），意向换班日期，返回一个换班id,其他人想和发起人换班的，输入发起人的换班id,系统将自动处理换班
## 补班处理
发起人填写补班日期，和意向补班的空闲日期，系统将自动处理补班的日期
## 蹭班处理
发起人填写蹭班日期，系统将自动处理
## 管理
管理API使得绕过系统正常的流程直接控制排班的结果：
- 删除选定日期的选定人员值班
- 增加选定日期的选定人员值班
## 鉴权
可以让他们在gitea上注册，也可以使用basic auth

## 技术

配置：viper
数据库：gorm
web框架:echo
### 选型
生成值班表时可以随用随读`member.csv`,换班补班等信息可以使用sqlite来存储，或许应该将生成的结果也应该用SQLite缓存
