# 通用API
## 说明
这里是本项目的API调用文档

uri一般为`/api/...`
## 用户
### getAssignment
- 路径:`/api/getAssignment`
- 方法:`GET`
- 功能：获取值班表
- 返回:HTML格式，没有CSS
- 参数：`date`,需要一个格式为`yyyy-mm-dd`的日期,没有参数的话默认返回当前时间的值班表
