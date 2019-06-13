package uniqueparse

import (
	"ant/reflectutil"
	"fmt"
	"regexp"
	"strings"
)

//数据库重复性错误解析工具,此解析工具只用于单字段unique的错误解析

//场景,在以前的开发中,数据表某个字段重复,是在程序中判定的,这样并发情况下,依然会有数据插入
//进阶版本1:为当前逻辑加锁(加锁这种方式是推荐的,因为加锁会直接降低系统的并发性),但项目中修改数据表的入口可能会有多个,这样并发依然可以插入
//进阶版本2:数据库字段加唯一索引,为保证友好提示,在插入数据之前判定数据是否重复,但在并发情况下依然会有数据表字段重复的错误

//最终版本:程序中不加锁,依靠数据库保证数据唯一性,对数据库错误进行解析,以达到友好提示的目的

//Error 1062: Duplicate entry '23421' for key 'username'

type UniqueMysql struct {
	TableFieldMap map[string]map[string]string //存储数据表和字段的映射关系 tableName:fieldName:fieldDesc

}

func NewUniqueMysql(obj ...interface{}) *UniqueMysql {
	this:=new(UniqueMysql)
	this.TableFieldMap=make(map[string]map[string]string)


	for _,v:=range obj {
		this.SetTableFieldMapFromStruct(v)
	}
	return this
}


//直接设置映射关系
func (this *UniqueMysql)SetTableFieldMap(tableFieldMap map[string]map[string]string)  *UniqueMysql{
	for k,v:=range tableFieldMap {
		this.TableFieldMap[k]=v
	}

	return this
}

//通过结构体设置映射关系
func (this *UniqueMysql)SetTableFieldMapFromStruct(obj interface{})*UniqueMysql  {
	objT, _, _ := reflectutil.GetStructTV(obj)

	tableName:=this.getStructName(obj)
	if _,ok:=this.TableFieldMap[tableName];ok {
		return this
	}

	tableMap:=make(map[string]string)

	for i:=0;i<objT.NumField();i++{
		structTag:=objT.Field(i).Tag
		fieldName:=reflectutil.GetStructTagFuncContent(structTag,"orm","column")
		if fieldName== "" {
			continue
		}

		fieldDesc:=structTag.Get("field")

		tableMap[fieldName]=fieldDesc
	}

	this.TableFieldMap[tableName]=tableMap

	return this
}

//解析数据库字段重复错误,返回友好提示错误
func (this *UniqueMysql)Parse(obj interface{},err error,errTpl ...string) error {

	if len(errTpl)== 0 {
		errTpl=[]string{"%s`%s`已经存在"}
	}

	tableName:=this.getStructName(obj)

	es:=err.Error()

	//非重复性错误,不解析
	if !strings.Contains(es, "1062") {
		return err
	}

	//通过正则提取字段名
	re:=regexp.MustCompile(`'(.*)' for key '([\d\w_]+)'`)
	fieldName:=re.FindStringSubmatch(es)

	if len(fieldName)< 3 {
		return err
	}

	fieldDesc:=this.TableFieldMap[tableName][fieldName[2]]

	if fieldDesc!= "" {
		return fmt.Errorf(errTpl[0],fieldDesc,fieldName[1])
	}

	return err
}

//获取结构体名称
func (this *UniqueMysql)getStructName(obj interface{})string {
	objT, _, _ := reflectutil.GetStructTV(obj)

	return objT.String()
}