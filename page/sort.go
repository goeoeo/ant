package page

import (
	"encoding/json"
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
func SortSlice(slicePtr interface{}, sortFields string /* Field Aes */) error {

	var (
		err            error
		sortFieldSlice []sortField
		arrSlice       []interface{}
		content        []byte
	)

	if sortFieldSlice, err = parseField(sortFields); err != nil {
		return err
	}

	//没有排序字段
	if len(sortFieldSlice) == 0 {
		return nil
	}

	if arrSlice, err = toSlice(slicePtr); err != nil {
		return err
	}

	//执行排序
	sort.Slice(arrSlice, func(i, j int) bool {

		for _, v := range sortFieldSlice {
			a := getObjVal(arrSlice[i], v.field)
			b := getObjVal(arrSlice[j], v.field)

			//当前排序字段值相等跳过
			if reflect.DeepEqual(a, b) {
				continue
			}

			switch v.val {
			case aes:

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

			case desc:
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
			}

		}

		return false

	})

	//将排序内容转回去
	if content, err = json.Marshal(arrSlice); err != nil {
		return err
	}

	//slicePtr=unsafe.Pointer(&arrSlice)

	return json.Unmarshal(content, slicePtr)

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
func toSlice(arr interface{}) ([]interface{}, error) {
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
func getObjVal(obj interface{}, field string) interface{} {
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
