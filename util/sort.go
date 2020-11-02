package util

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"
)

type sortField struct {
	field string //排序字段
	val   string //aes=升序，desc=降序
}

const (
	aes  = "aes"
	desc = "desc"
)

//多字段排序
//SliceSort([]User ,"Name desc,Id aes"})
func SortSlice(slicePtr interface{}, sortFields string /* Field Aes */) (err error) {

	var (
		sortFieldSlice   []sortField
		arrSlice         []reflect.Value
		slicePtrValue    reflect.Value
		newSlicePtrValue reflect.Value
	)

	if slicePtrValue = reflect.ValueOf(slicePtr); slicePtrValue.Kind() != reflect.Ptr {
		return errors.New("排序源数据必须为切片指针")
	}

	//取指针对应的值
	slicePtrValue = slicePtrValue.Elem()
	if slicePtrValue.Kind() != reflect.Slice {
		return errors.New("排序源数据必须为切片指针.")
	}

	if slicePtrValue.Len() == 0 {
		return
	}

	//制造相同数组
	newSlicePtrValue = reflect.MakeSlice(slicePtrValue.Type(), slicePtrValue.Len(), slicePtrValue.Cap())

	//解析排序字段
	if sortFieldSlice, err = parseField(sortFields); err != nil {
		return err
	}

	//没有排序字段
	if len(sortFieldSlice) == 0 {
		return nil
	}

	for i := 0; i < slicePtrValue.Len(); i++ {
		//拷贝值
		newSlicePtrValue.Index(i).Set(slicePtrValue.Index(i))

		arrSlice = append(arrSlice, newSlicePtrValue.Index(i))
	}

	//执行排序
	sort.Slice(arrSlice, func(i, j int) bool {

		for _, v := range sortFieldSlice {
			if v.val != aes && v.val != desc {
				continue
			}

			a := arrSlice[i].FieldByName(v.field)
			b := arrSlice[j].FieldByName(v.field)

			//当前排序字段值相等跳过
			if reflect.DeepEqual(a.Interface(), b.Interface()) {
				continue
			}

			if v.val == aes {
				//升序
				return lessValue(a, b)
			} else {
				//降序
				return !lessValue(a, b)
			}

		}

		return false

	})

	//将排序内容转回去
	for k, v := range arrSlice {
		slicePtrValue.Index(k).Set(v)
	}

	return nil

}

//解析排序字段
func parseField(sortFields string) (sortFieldsSlice []sortField, err error) {
	var (
		sortFieldsArr []string
	)
	sortFieldsArr = strings.Split(sortFields, ",")

	for _, v := range sortFieldsArr {
		tmp := strings.Split(v, " ")
		if len(tmp) != 2 {
			return nil, errors.New("排序字段解析错误")
		}
		//升降序指令，统一转小写
		tmp[1] = strings.ToLower(tmp[1])
		if !inArray(tmp[1], []string{aes, desc}) {
			return nil, errors.New(fmt.Sprintf("排序字段解析错误,排序指令只支持:%s,%s", aes, desc))
		}

		sortFieldsSlice = append(sortFieldsSlice, sortField{field: tmp[0], val: tmp[1]})
	}

	return

}

func inArray(item string, items []string) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}

//slice interface 变数组
func toSlice(arr interface{}) ([]reflect.Value, error) {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Ptr {
		return nil, errors.New("排序源数据必须为切片指针")
	}

	ve := reflect.ValueOf(arr).Elem()
	if ve.Kind() != reflect.Slice {
		return nil, errors.New("排序源数据必须为切片指针.")
	}

	l := ve.Len()
	ret := make([]reflect.Value, l)
	for i := 0; i < l; i++ {
		ret[i] = ve.Index(i)
	}
	return ret, nil
}

//a<b 判定
func lessValue(a reflect.Value, b reflect.Value) bool {

	switch a.Kind() {
	case reflect.String:
		return a.String() < b.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return a.Int() < b.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return a.Uint() < b.Uint()
	case reflect.Bool:
		//bool小于的情况
		if a.Bool() == false && b.Bool() == true {
			return true
		}
	case reflect.Float32, reflect.Float64:
		return a.Float() < b.Float()
	case reflect.Struct:
		//时间比较
		if _, ok := a.Interface().(time.Time); ok {
			return a.Interface().(time.Time).Unix() < b.Interface().(time.Time).Unix()
		}

	}

	return false
}
