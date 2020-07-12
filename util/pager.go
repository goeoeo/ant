package util

import (
	"encoding/json"
	"errors"
	"reflect"
)

type Pager struct {
	Page      int         //当前页
	PageSize  int         //页大小 ,页大小 小于 0 返回所有数据
	total     int64       //总量
}

//是否进行分页
func (this *Pager) PageEnable() bool {
	return this.PageSize > 0
}

//切片分页 pageSize=-1 返回全部数据
func (this *Pager) Pagination(data interface{}) error {

	var (
		resSlice []interface{}
		arrSlice []interface{}
		start    int64
		end      int64
		err      error
		content  []byte
	)

	if arrSlice, err = this.toSlice(data); err != nil {
		return err
	}

	this.total = int64(len(arrSlice))

	//分页大小=-1 或者 当前页大小不大于0 不进行分页处理
	if this.PageSize == -1 || this.Page <= 0 {
		return nil
	}

	start = int64((this.Page - 1) * this.PageSize)
	end = int64(this.Page * this.PageSize)
	if end > this.total {
		end = this.total
	}

	for start < end {
		resSlice = append(resSlice, arrSlice[start])
		start++
	}

	if content, err = json.Marshal(resSlice); err != nil {
		return err
	}

	return json.Unmarshal(content, data)
}


//总量
func (this *Pager)Total() int64 {
	return this.total
}

//页大小
func (this *Pager) Limit() int {

	//分页量小于0，返回全部数据
	if this.PageSize < 0 {
		return int(this.total)
	}

	return this.PageSize
}

//偏移量
func (this *Pager) Offset() int {

	if this.PageSize < 0 {
		return 0
	}

	offset := (this.Page - 1) * this.PageSize
	if offset < 0 {
		offset = 0
	}

	return offset
}

//slice interface 变数组
func (this *Pager) toSlice(arr interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Ptr {
		return []interface{}{}, errors.New("分页源数据必须为切片指针")
	}

	ve := reflect.ValueOf(arr).Elem()
	if ve.Kind() != reflect.Slice {
		return []interface{}{}, errors.New("分页源数据必须为切片指针.")
	}

	l := ve.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = ve.Index(i).Interface()
	}
	return ret, nil
}
