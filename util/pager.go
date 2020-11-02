package util

import (
	"errors"
	"fmt"
	"reflect"
)

type Pager struct {
	Page     int   //当前页
	PageSize int   //页大小 ,页大小 小于 0 返回所有数据
	Error    error `json:"-"`
	total    int64 `json:"-"` //总量
}

//是否进行分页
func (this *Pager) PageEnable() bool {
	return this.PageSize > 0 && this.Page > 0
}

//切片分页 pageSize=-1 返回全部数据
func (this *Pager) Pagination(src interface{}) *Pager {
	if err := this.pagination(src); err != nil {
		this.Error = fmt.Errorf("分页错误:%s", err)
		return this
	}

	return this
}

//总量
func (this *Pager) Total(total *int64) *Pager {
	*total = this.total
	return this
}

//页大小
func (this *Pager) Limit() int {

	if !this.PageEnable() {
		return 0
	}

	return this.PageSize
}

//偏移量
func (this *Pager) Offset() int {

	if !this.PageEnable() {
		return 0
	}

	offset := (this.Page - 1) * this.PageSize
	if offset < 0 {
		offset = 0
	}

	return offset
}

//slice interface 变数组
func (this *Pager) pagination(arr interface{}) (err error) {
	ve := reflect.ValueOf(arr)
	if ve.Kind() != reflect.Ptr {
		return errors.New("分页源数据必须为切片指针")
	}

	ve = ve.Elem()
	if ve.Kind() != reflect.Slice {
		return errors.New("分页源数据必须为切片指针.")
	}

	//数据总量
	this.total = int64(ve.Len())

	if !this.PageEnable() {
		return
	}

	limit := this.Limit()
	offset := this.Offset()
	start := 0

	for {
		if start >= limit || start+offset >= int(this.total) {
			break
		}

		ve.Index(start).Set(ve.Index(start + offset))
		start++
	}
	ve.SetLen(start)

	return
}
