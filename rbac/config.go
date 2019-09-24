package rbac

import (
	"github.com/phpdi/ant/reflectutil"
	"errors"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"reflect"
	"time"
)

type TransactionClosure = func(o orm.Ormer) error

//k-v结构体
type Kv struct {
	Key   interface{}
	Value interface{}
}



type Config struct {
	Orm orm.Ormer
	PermTable string
	RoleTable string
	RolePermTable string
	RoleUserTable string

	UserPermCachePreFix string 	//缓存前缀
	UserPermCacheTime time.Duration //缓存时间
}

var RoleService roleService
var ConfigService Config
var PermService permService
var CacheService cache.Cache

//初始化rbac服务
func InitRbac()  {

}

//事务闭包
func Transaction(f TransactionClosure) error {
	o := orm.NewOrm()

	err := o.Begin()
	if err != nil {
		return err
	}

	err = f(o)

	if err != nil {
		terr := o.Rollback()
		if terr != nil {
			logs.Warn("事务回滚异常:" + terr.Error())
		}

		return err
	}

	terr := o.Commit()
	if terr != nil {
		logs.Warn("事务提交异常:" + terr.Error())
	}

	return nil

}

//获取kv结构
func GetKv(obj interface{}, key string, value string) []Kv {

	res := []Kv{}

	objV := reflect.ValueOf(obj)
	if objV.Kind() != reflect.Slice {
		return res
	}

	l := objV.Len()
	for i := 0; i < l; i++ {
		res = append(res, Kv{
			Key:   objV.Index(i).FieldByName(key).Interface(),
			Value: objV.Index(i).FieldByName(value).Interface(),
		})
	}

	return res
}

//数据重复
func Repeat(obj interface{}, table string, pk string, field string) (error, bool) {

	objT := reflect.TypeOf(obj)
	if !reflectutil.IsStructPtr(objT) {
		return errors.New("obj 必须为结构体指针"), true
	}

	objV := reflect.ValueOf(obj).Elem()

	q := ConfigService.Orm.QueryTable(table).Filter(field, objV.FieldByName(field).Interface())

	pkVal := objV.FieldByName(pk).Int()
	if pkVal != 0 {
		q = q.Exclude(pk, pkVal)
	}

	err := q.One(obj)

	if err != nil {
		//没有重复的数据
		return nil, false
	}

	return nil, true
}