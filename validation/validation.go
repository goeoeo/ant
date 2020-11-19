package validation

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
)

type (

	//数据验证器,用于验证结构体字段内容是否合法
	Validation struct {
		Config *ValidationConfig

		requireFields []string                  //存放非零值字段
		fieldTag      map[string]string         //手动设置验证规则
		failMessages  map[string] /*字段*/ string /*自定义的错误消息*/ //字段验证失败的提示信息,当没有设置字段信息时，使用拼接的方式进行错误显示
	}
)

//new
func NewValidation(config *ValidationConfig) *Validation {
	this := &Validation{
		Config:       config,
		failMessages: make(map[string]string),
		fieldTag:     make(map[string]string),
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

//获取字段对应的错误信息
func (this *Validation) getError(field string, name string, msg string) error {

	//自定义的错误提示优先
	if msg, ok := this.failMessages[field]; ok {
		return errors.New(this.Config.TranFunc(msg))
	}

	//使用拼接的方式提示错误
	if name == "" {
		name = field
	}

	if nameArr := strings.Split(name, ";"); len(nameArr) > 0 {
		name = nameArr[0]
	}

	reg := regexp.MustCompile(`\p{Han}+`)
	//非汉字加空格
	if !reg.Match([]byte(msg)) {
		msg = " " + msg
	}
	return errors.New(this.Config.TranFunc(name) + msg)
}

//设置翻译回调函数
func (this *Validation) SetTrFun(f func(s string, params ...interface{}) string) *Validation {
	this.Config.TranFunc = f
	return this
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

//结构体验证遇到一个验证不通过就会退出
func (this *Validation) Valid(obj interface{}) error {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)

	if objT.Kind() == reflect.Ptr {
		objT = objT.Elem()
	}

	if objT.Kind() != reflect.Struct {
		return errors.New(this.Config.TranFunc("the verification object can only be a structure or a structure pointer"))
	}

	l := objT.NumField()
	for i := 0; i < l; i++ {
		if err := this.validStructField("", objT.Field(i), objV.Field(i)); err != nil {
			return err
		}
	}

	return nil
}

//验证结构体字段
func (this *Validation) validStructField(parentFiledName string, structFiled reflect.StructField, rv reflect.Value) error {

	//指针处理
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if parentFiledName != "" {
		parentFiledName += "." + structFiled.Name
	} else {
		parentFiledName = structFiled.Name
	}
	switch rv.Kind() {
	case reflect.Struct:
		l := rv.NumField()
		for i := 0; i < l; i++ {

			if err := this.validStructField(parentFiledName, rv.Type().Field(i), rv.Field(i)); err != nil {
				return err
			}
		}
		return nil

	case reflect.Invalid, reflect.Complex64, reflect.Complex128, reflect.Chan, reflect.Func, reflect.UnsafePointer, reflect.Ptr: //不支持的验证
		return nil

	}

	//零值验证
	if this.inArray(parentFiledName, this.requireFields) && rv.IsZero() {
		return this.getError(parentFiledName, structFiled.Tag.Get(this.Config.StructFieldName), this.Config.TranFunc(this.Config.messageTmpls["Required"]))
	}

	//定义了验证函数
	if this.getTagString(parentFiledName, structFiled.Tag) != "" {

		//当前字段需要执行的验证函数
		funcsMap := this.parseFunc(structFiled.Tag.Get(this.Config.StructTagField))
		for k, v := range funcsMap {
			if tmpFunc, ok := this.Config.validFuns[k]; ok {
				//验证未通过
				if !tmpFunc(rv.Interface(), v...) {
					name := structFiled.Tag.Get(this.Config.StructFieldName)
					if msg, ok := this.Config.messageTmpls[k]; ok {
						//定义了验证不通过的错误消息
						formatParams := []interface{}{}
						for _, v1 := range v {
							formatParams = append(formatParams, v1)
						}
						msg = this.Config.TranFunc(msg, formatParams...)
						return this.getError(parentFiledName, name, msg)
					} else {
						return this.getError(parentFiledName, name, this.Config.TranFunc("verification failed"))
					}
				}
			}

		}

	}

	return nil
}

//获取字段tag
func (this *Validation) getTagString(fieldName string, structTag reflect.StructTag) string {
	//自定义配置优先
	if tag, ok := this.fieldTag[fieldName]; ok {
		structTag = reflect.StructTag(tag)
	}

	return structTag.Get(this.Config.StructTagField)
}

//解析函数返回函数名和参数的k-v结构
func (this *Validation) parseFunc(structTag string) map[string][]string /*函数名=>参数*/ {
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

//判定某个值是否在数组里面
func (this *Validation) inArray(field string, arr []string) bool {
	for _, v := range arr {
		if v == field {
			return true
		}
	}

	return false
}
