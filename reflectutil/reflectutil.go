package reflectutil

import (
	"ant/stringutil"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

//判定一个interface的值是否为空值
func IsEmpty(obj interface{}) bool {
	if obj == nil {
		return true
	}

	switch val:=obj.(type) {
	case string:
		return len(strings.TrimSpace(val)) == 0
	case bool:
		return true
	case int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64:
		return val==0
	case float32:
		return val==float32(0)
	case float64:
		return val==float64(0)
	case time.Time:
		return val.IsZero()
	}

	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Slice {
		return v.Len() == 0
	}
	return false
}

//判定是否为结构体
func IsStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

//判定是否为结构体指针
func IsStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

//获取结构体或者指针的类型和值
func GetStructTV(obj interface{})(reflect.Type,reflect.Value,error) {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)

	switch {
	case IsStruct(objT):
	case IsStructPtr(objT):
		objT = objT.Elem()
		objV = objV.Elem()
	default:
		return objT,objV, fmt.Errorf("%v must be a struct or a struct pointer", obj)
	}

	return objT,objV,nil
}

//获取获取结构体中structTag中函数参数内容
func GetStructTagFuncContent(structTag reflect.StructTag,field string,funcName string) string {
	tag:=structTag.Get(field)

	re := regexp.MustCompile(fmt.Sprintf(`%s\(([^(]*)\)`,funcName))

	res := re.FindStringSubmatch(tag)

	if len(res) > 0 {
		return res[1]
	}

	return ""
}

//获取结构体的非零值字段
func GetNotEmptyFields(obj interface{}, fields ...string) []string {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)

	objT,objV,err:=GetStructTV(obj)
	if err != nil {
		return fields
	}

	for i := 0; i < objT.NumField(); i++ {
		//字段名称
		currentField := objT.Field(i).Name

		//字段值
		currentFieldValue := objV.Field(i).Interface()
		if objV.Field(i).Kind() == reflect.Ptr {
			if objV.Field(i).IsNil() {
				currentFieldValue = ""
			} else {
				currentFieldValue = objV.Field(i).Elem().Interface()
			}
		}

		if !IsEmpty(currentFieldValue) && !stringutil.InSliceString(currentField, fields) {
			fields = append(fields, currentField)
		}

	}
	return fields
}