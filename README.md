### ant
积累自己开发过的go模块,使用go-mod管理依赖


### bulidsql(基于结构体的sql构建器) 
* Insert 插入
* Update 更新
* Delete 删除
* Select 查询
* Emphasize 强调字段,默认情况,零值字段不会参与构建,通过此函数强调后,则可参与构建
* Field 指定查询字段
* Where  
* OrWhere 
* GroupBy 
* OrderBy 
* LeftJoin 
* Limit 

### reflectutil(反射工具)
* IsEmpty 判定一个interface的值是否为空值
* IsStruct 判定是否为结构体
* IsStructPtr 判定是否为结构体指针
* GetStructTV 获取结构体或者指针的类型和值
* GetStructTagFuncContent 获取获取结构体中structTag中函数参数内容
* GetNotEmptyFields 获取结构体的非零值字段

### numberutil(数字工具)
* InSliceInt  判定某个值是否在数组里面 
* RandomNumber 生成随机数字 

### stringutil(字符串工具)
* InSliceString 判定某个值是否在数组里面
* Md5 
* Ip2Long ip转换函数,字符串转数字
* Long2Ip ip转换函数,数字转字符串
* RandomString 生成随机字符串
* Keep0Add 保留左边0的字符串加法

