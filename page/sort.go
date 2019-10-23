package page

import (
	"encoding/json"
	"errors"
	"reflect"
	"sort"
	"time"
)

//排序结构体
type Sort struct {
	SortField string //排序字段
	SortVal   int    //0=不排序，1=升序，2=降序
}

const (
	SortValNot  = 0
	SortValAsc  = 1
	SortValDesc = 2
)

//排序
func (this *Sort) SortSlice(slicePtr interface{}) error {
	//不进行排序
	if this.SortVal == SortValNot || slicePtr == nil {
		return nil
	}

	var (
		arrSlice []interface{}
		err      error
		content  []byte
	)

	if arrSlice, err = this.toSlice(slicePtr); err != nil {
		return err
	}

	sort.Slice(arrSlice, func(i, j int) bool {
		//升序
		a := this.getObjVal(arrSlice[i], this.SortField)
		b := this.getObjVal(arrSlice[j], this.SortField)

		switch this.SortVal {
		case SortValAsc:

			switch at := a.(type) {
			case string:

				bt, _ := b.(string)
				return at < bt
			case int64:
				bt, _ := b.(int64)
				return at < bt
			case uint64:
				bt, _ := b.(uint64)
				return at < bt
			}
			return false

		case SortValDesc:
			//降序
			switch at := a.(type) {
			case string:
				bt, _ := b.(string)
				return at > bt
			case int64:
				bt, _ := b.(int64)
				return at > bt
			case uint64:
				bt, _ := b.(uint64)
				return at > bt
			}

			return false

		default:
			return false
		}
	})

	if content, err = json.Marshal(arrSlice); err != nil {
		return err
	}

	return json.Unmarshal(content, slicePtr)

}

//slice interface 变数组
func (this *Sort) toSlice(arr interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Ptr {
		return []interface{}{}, errors.New("排序源数据必须为切片指针")
	}

	ve := reflect.ValueOf(arr).Elem()
	if ve.Kind() != reflect.Slice {
		return []interface{}{}, errors.New("排序源数据必须为切片指针.")
	}

	l := ve.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = ve.Index(i).Interface()
	}
	return ret, nil
}

//获取对象属性值
func (this *Sort) getObjVal(obj interface{}, field string) interface{} {
	if obj == nil {
		return ""
	}
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	if objT.Kind() == reflect.Ptr && objT.Elem().Kind() == reflect.Struct {
		objV = objV.Elem()
	}

	switch objV.FieldByName(field).Kind() {
	case reflect.String:
		return objV.FieldByName(field).String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return objV.FieldByName(field).Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return objV.FieldByName(field).Uint()
	case reflect.Struct:
		//时间比较
		if timeTmp, ok := objV.FieldByName(field).Interface().(time.Time); ok {
			return timeTmp.String()
		}

	}
	return ""
}
