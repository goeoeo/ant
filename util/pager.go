package util

import (
	"bytes"
	"encoding/gob"
	"errors"
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
func (this *Pager) Pagination(data interface{}) *Pager {
	var (
		resSlice []interface{}
		err      error
		contentBuffer bytes.Buffer
	)

	defer func() {
		if err != nil {
			this.Error = err
		}
	}()

	if resSlice, err = this.toSlice(data); err != nil {
		return this
	}

	//不进行分页
	if !this.PageEnable() {
		return this
	}


	//这里不用json序列化的原因是如果遇到proto结构体 omitempty tag 会导致后面的页字段零值情况使用到第一页数据的非零值
	//编码
	enc := gob.NewEncoder(&contentBuffer)
	if err = enc.Encode(resSlice); err == nil {
		return this
	}

	//解码
	dec := gob.NewDecoder(&contentBuffer)
	if err = dec.Decode(data); err != nil {
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
func (this *Pager) toSlice(arr interface{}) (res []interface{}, err error) {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Ptr {
		return []interface{}{}, errors.New("分页源数据必须为切片指针")
	}

	ve := reflect.ValueOf(arr).Elem()
	if ve.Kind() != reflect.Slice {
		return []interface{}{}, errors.New("分页源数据必须为切片指针.")
	}

	//数据总量
	this.total = int64(ve.Len())

	if !this.PageEnable() {
		return
	}

	res = []interface{}{}

	limit := this.Limit()
	offset := this.Offset()
	start := 0

	for {
		if start >= limit || start+offset >= int(this.total) {
			break
		}

		res = append(res, ve.Index(start+offset).Interface())
		start++
	}

	return
}
