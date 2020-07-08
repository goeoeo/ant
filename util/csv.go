package util

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"reflect"
)

type CsvUtil struct {
	buff *bytes.Buffer
	Wr   *csv.Writer
}

func NewCsv(title ...string) *CsvUtil {
	this := &CsvUtil{buff: new(bytes.Buffer)}
	this.Wr = csv.NewWriter(this.buff)

	//防止word打开出现乱码
	this.buff.WriteString("\xEF\xBB\xBF")

	//写入了表头
	if len(title) > 0 {
		this.Row(title...)
	}

	return this
}

//内容行
func (this *CsvUtil) Row(row ...string) error {
	return this.Wr.Write(row)
}

//输出数据
func (this *CsvUtil) Bytes() []byte {

	this.Wr.Flush()
	return this.buff.Bytes()
}

//从slice中选中字段进行导出
func (this *CsvUtil)Slice2Csv(arr interface{},fields ...string) (content []byte,err error){

	arrT:=reflect.TypeOf(arr)
	if arrT.Kind()!=reflect.Slice {
		return content,errors.New("arr必须为slice")
	}

	if len(fields) == 0 {
		return
	}

	arrV:=reflect.ValueOf(arr)

	for i:=0;i<arrV.Len();i++ {
		item:=arrV.Index(i)
		//指针转换
		if item.Kind() == reflect.Ptr {
			item=item.Elem()
		}

		row:=[]string{}
		for _,field:=range fields {
			row=append(row,fmt.Sprintf("%v",item.FieldByName(field).Interface()))
		}

		if err=this.Row(row...);err!= nil {
			return
		}
	}

	content=this.Bytes()
	return
}
