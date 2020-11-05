package util

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

type Pager struct {
	Page       int   //当前页
	PageSize   int   //页大小 ,页大小 小于 0 返回所有数据
	Error      error `json:"-"`
	total      int64 `json:"-"` //总量
	strictMode bool  //严格模式，会返回真实计算出的偏移量。如果偏移量不合法，那么原切片会被置空
}

//是否进行分页
func (this *Pager) PageEnable() bool {
	return this.PageSize > 0 && this.Page > 0
}

//开启严格模式
func (this *Pager) StrictMode() *Pager {
	this.strictMode = true
	return this
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
	var (
		v  reflect.Value
		ve reflect.Value
	)
	v = reflect.ValueOf(arr)
	if v.Kind() != reflect.Ptr {
		return errors.New("分页源数据必须为切片指针")
	}

	ve = v.Elem()
	if ve.Kind() != reflect.Slice {
		return errors.New("分页源数据必须为切片指针.")
	}

	//数据总量
	this.total = int64(ve.Len())
	//偏移量
	offset := this.Offset()

	//检查偏移量是否越界
	if !this.PageEnable() || offset >= int(this.total) {
		if this.strictMode {
			//严格模式，切片置空
			ve.SetLen(0)
		}
		return
	}

	sliHeader := (*reflect.SliceHeader)(unsafe.Pointer(v.Pointer()))

	//取切片元素大小
	sizeItem := reflect.TypeOf(arr).Elem().Elem().Size()

	//移动指针
	sliHeader.Data += uintptr(offset) * sizeItem

	limit := this.Limit()
	//防止越界
	if int(this.total)-offset < limit {
		limit = int(this.total) - offset
	}
	sliHeader.Len = limit

	return
}
