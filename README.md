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

### validation(基于structTag的结构体格式验证器)
> 思路来源于beego框架,beego的表单验证器扩展很复杂,于是自己实现了这个验证器,扩展简单.
> 可将验证器定义为全局变量,以便在程序的任意一个地方调用,不用反复实例化

例如:扩展一个最长度的验证规则
1.实现函数MinSize,params参数和你定义到structTag中MinSize()函数的参数一一对应
```go
func MinSize(validValue interface{}, params ...string) bool {
	...
}
```
2.将此函数注册到你的验证器上
```go
type User struct {
	Id               int     
	Code             string                
	Name             string  `field:"姓名";valid:"MinSize(6)"` //定义Name最小长度为6       
}

user:=User{Name:"test"}
v:=NewValidation()
v.RegisterFun("MinSize", MinSize) //注册函数到验证器
v.SetMessageTmpls(map[string]string{"MinSize":""最小长度为%v""}) //注册出错的模板消息
err:=validate.Valid(user) //err 姓名最小长度为6
```


* Min(min int) 最小值，有效类型：int，其他类型都将不能通过验证
* Max(max int) 最大值，有效类型：int，其他类型都将不能通过验证
* Range(min, max int) 数值的范围，有效类型：int，他类型都将不能通过验证
* MinSize(min int) 最小长度，有效类型：string slice，其他类型都将不能通过验证
* MaxSize(max int) 最大长度，有效类型：string slice，其他类型都将不能通过验证
* Length(length int) 指定长度，有效类型：string slice，其他类型都将不能通过验证
* Alpha alpha字符，有效类型：string，其他类型都将不能通过验证
* Numeric 数字，有效类型：string，其他类型都将不能通过验证
* AlphaNumeric alpha 字符或数字，有效类型：string，其他类型都将不能通过验证
* Match(pattern string) 正则匹配，有效类型：string，其他类型都将被转成字符串再匹配(fmt.Sprintf(“%v”, obj).Match)
* AlphaDash alpha 字符或数字或横杠 -_，有效类型：string，其他类型都将不能通过验证
* Email 邮箱格式，有效类型：string，其他类型都将不能通过验证
* IP IP 格式，目前只支持 IPv4 格式验证，有效类型：string，其他类型都将不能通过验证
* Mobile 手机号，有效类型：string，其他类型都将不能通过验证
* Tel 固定电话号，有效类型：string，其他类型都将不能通过验证
* Phone 手机号或固定电话号，有效类型：string，其他类型都将不能通过验证
* ZipCode 邮政编码，有效类型：string，其他类型都将不能通过验证
* Mac Mac地址，有效类型：string，其他类型都将不能通过验证
* ChnDash 中文,数字,字母,下划线,有效类型：string，其他类型都将不能通过验证

###uniqueparse数据库重复字段错误解析工具
>* 此工具用于解析数据库唯一字段的重复性错误
>* 此解析工具只用于单字段unique的错误解析

例如:

数据库错误为://Error 1062: Duplicate entry '23421' for key 'username'

解析后://账号`23421`已经存在

```go
unique:=NewUniqueMysql(new(ant.User))

err:=errors.New("Error 1062: Duplicate entry '23421' for key 'username'")
err=unique.Parse(ant.User{},err)
fmt.Println(err)
//账号`23421`已经存在
```
