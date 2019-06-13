package uniqueparse

import (
	"ant"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

var Umysql *UniqueMysql

func init() {
	// 注册驱动
	err:=orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Println(err)
	}
	// 设置默认数据库
	err=orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/lwyweb?charset=utf8", 30)
	if err != nil {
		fmt.Println(err)
	}

	// 注册定义的 model
	orm.RegisterModelWithPrefix("t_",new(ant.User))
	orm.Debug=true
	// 创建 table
	//orm.RunSyncdb("default", false, true)

	Umysql=NewUniqueMysql(ant.User{})
	//fmt.Println(Umysql.TableFieldMap)
}


func TestUniqueMysql_Parse(t *testing.T) {
	unique:=NewUniqueMysql(new(ant.User))

	err:=errors.New("Error 1062: Duplicate entry '23421' for key 'username'")
	err=unique.Parse(ant.User{},err)
	fmt.Println(err)
	//账号`23421`已经存在
}

func TestUniqueMysql_Parse2(t *testing.T) {

	user:=ant.User{UserName:"admin"}
	o:=orm.NewOrm()

	_,err:=o.Insert(&user)

	if err != nil {
		fmt.Println(Umysql.Parse(user,err))
	}

}