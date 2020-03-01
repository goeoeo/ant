package validation

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

//数据验证器,用于验证结构体字段内容是否合法
//目标，只要注册了方法就可以直接使用

type ValidFun func(validValue interface{}, params ...string) bool

// A Validation context manages data validation and error messages.
type Validation struct {
	structTagField  string //结构体 验证器structTag名称
	structFieldName string //结构体字段名称

	validateOne bool //只验证一个字段

	validFuns      map[string]ValidFun
	requireFields  []string          //存放非零值字段
	priorityFields []string          //优先验证字段，可通过此属性，调整字段验证顺序
	fieldTag       map[string]string //手动设置验证规则

	failMessages map[string]string //字段验证失败的提示信息,当没有设置字段信息时，使用拼接的方式进行错误显示
	messageTmpls map[string]string //验证失败函数对应的模板消息
}

//外部调用实例
var Valid *Validation

func init() {
	Valid = NewValidation()
}

//new
func NewValidation() *Validation {
	this := &Validation{
		structTagField:  "valid",
		structFieldName: "field",
	}

	this.validFuns = make(map[string]ValidFun)
	this.failMessages = make(map[string]string)
	this.fieldTag = make(map[string]string)

	this.messageTmpls = MessageTmpls

	//注册函数
	this.RegisterFun("Max", Max).
		RegisterFun("Min", Min).
		RegisterFun("Range", Range).
		RegisterFun("MinSize", MinSize).
		RegisterFun("MaxSize", MaxSize).
		RegisterFun("Length", Length).
		RegisterFun("Alpha", Alpha).
		RegisterFun("Numeric", Numeric).
		RegisterFun("AlphaNumeric", AlphaNumeric).
		RegisterFun("AlphaDash", AlphaDash).
		RegisterFun("Email", Email).
		RegisterFun("IP", IP).
		RegisterFun("Mobile", Mobile).
		RegisterFun("Tel", Tel).
		RegisterFun("ZipCode", ZipCode).
		RegisterFun("Mac", Mac).
		RegisterFun("ChnDash", ChnDash).
		RegisterFun("ChnAlphaNumeric", ChnAlphaNumeric)

	return this

}

//设定函数对应的模板消息
func (this *Validation) SetMessageTmpls(messageTmpls map[string]string) *Validation {
	for k, v := range messageTmpls {
		this.messageTmpls[k] = v
	}

	return this
}

//设定字段出错时的错误消息
func (this *Validation) SetFailMessages(failMessages map[string]string) *Validation {
	for k, v := range failMessages {
		this.failMessages[k] = v
	}

	return this

}

//注册函数
func (this *Validation) RegisterFun(funcName string, validFunc ValidFun) *Validation {
	if len(this.validFuns) == 0 {
		this.validFuns = make(map[string]ValidFun)
	}
	this.validFuns[funcName] = validFunc

	return this
}

//获取字段对应的错误信息
func (this *Validation) getError(field string, name string, msg string) error {

	//自定义的错误提示优先
	if msg, ok := this.failMessages[field]; ok {
		return errors.New(msg)
	}

	//使用拼接的方式提示错误
	return errors.New(name + msg)
}

//手动设置FieldTag
func (this *Validation) FieldTag(fieldTag map[string]string) *Validation {
	this.fieldTag = fieldTag

	return this
}

//指定不为空字段
func (this *Validation) Require(fields ...string) *Validation {
	this.requireFields = fields
	return this
}

//设置优先验证字段，可通过此方法调整字段验证顺序
func (this *Validation) Priority(fields ...string) *Validation {
	this.priorityFields = fields
	return this
}

//结构体验证遇到一个验证不通过就会退出
func (this *Validation) Valid(obj interface{}) error {
	this.validateOne = true
	res := this.ValidAll(obj)

	for _, v := range res {
		return v
	}

	return nil
}

//会验证所有的字段
func (this *Validation) ValidAll(obj interface{}) map[string]error {

	res := make(map[string]error)

	objT, objV, err := this.GetStructTV(obj)
	if err != nil {
		return res
	}

	//设置验证结构体
	for i := 0; i < objT.NumField(); i++ {
		field := objT.Field(i).Name
		if this.inArray(field, this.priorityFields) {
			continue
		}
		this.priorityFields = append(this.priorityFields, field)
	}

	//优先字段,验证
	for _, v := range this.priorityFields {
		//字段类型
		fieldType, ok := objT.FieldByName(v)
		if !ok {
			continue
		}

		//字段值
		fieldValue := objV.FieldByName(v)

		if fieldTag, ok := this.fieldTag[v]; ok {
			fieldType.Tag = reflect.StructTag(fieldTag)
		}

		//执行此字段的相应验证规则
		err := this.validField(v, fieldType, fieldValue)

		if err != nil {
			res[v] = err

			//只验证一个字段
			if this.validateOne {
				break
			}
		}
	}

	return res
}

//验证字段
func (this *Validation) validField(field string, fieldType reflect.StructField, fieldValue reflect.Value) error {

	//零值验证
	if this.IsEmpty(fieldValue.Interface()) {
		if this.inArray(field, this.requireFields) {
			return this.getError(field, fieldType.Tag.Get(this.structFieldName), this.messageTmpls["Required"])
		}
		return nil
	}

	//定义了验证函数
	if fieldType.Tag.Get(this.structTagField) != "" {
		funcsMap := this.parseFunc(fieldType.Tag.Get(this.structTagField))
		for k, v := range funcsMap {
			if tmpFunc, ok := this.validFuns[k]; ok {
				if !tmpFunc(fieldValue.Interface(), v...) {
					//验证未通过
					name := fieldType.Tag.Get(this.structFieldName)
					if msg, ok := this.messageTmpls[k]; ok {
						//设置了提示信息
						formatParams := []interface{}{}
						for _, v1 := range v {
							formatParams = append(formatParams, v1)

						}
						return this.getError(field, name, fmt.Sprintf(msg, formatParams...))
					} else {
						return this.getError(field, name, "验证不通过")
					}

				}
			}

		}

	}

	return nil

}

//解析函数返回函数名和参数的k-v结构
func (this *Validation) parseFunc(structTag string) map[string][]string {
	funcs := strings.Split(structTag, ";")
	funcsMap := make(map[string][]string)
	for _, v := range funcs {
		start := strings.Index(v, "(")
		if start == -1 {
			funcsMap[v] = []string{}
			continue
		}

		end := strings.Index(v, ")")

		if end == -1 || start >= end {
			continue
		}

		funcName := strings.TrimSpace(v[:start])

		if funcName == "" {
			continue
		}

		funcParams := strings.TrimSpace(v[start+1 : end])
		funcsMap[funcName] = strings.Split(funcParams, ",")
	}

	return funcsMap
}

//判定是否为结构体
func (this *Validation) IsStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

//判定是否为结构体指针
func (this *Validation) IsStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

//获取结构体或者指针的类型和值
func (this *Validation) GetStructTV(obj interface{}) (reflect.Type, reflect.Value, error) {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)

	switch {
	case this.IsStruct(objT):
	case this.IsStructPtr(objT):
		objT = objT.Elem()
		objV = objV.Elem()
	default:
		return objT, objV, fmt.Errorf("%v must be a struct or a struct pointer", obj)
	}

	return objT, objV, nil
}

//判定一个interface的值是否为空值
func (this *Validation) IsEmpty(obj interface{}) bool {
	if obj == nil {
		return true
	}

	switch val := obj.(type) {
	case string:
		return len(strings.TrimSpace(val)) == 0
	case bool:
		return true
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val) == "0"
	case float32:
		return val == float32(0)
	case float64:
		return val == float64(0)
	case time.Time:
		return val.IsZero()
	}

	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Slice {
		return v.Len() == 0
	}
	return false
}

//判定某个值是否在数组里面
func (this *Validation) inArray(field string, arr []string) bool {
	for _, v := range arr {
		if v == field {
			return true
		}
	}

	return false
}