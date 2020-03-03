package autodoc

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

const (
	template = `
####%s

请求地址：%s

请求方式：%s

####请求参数
| 参数   | 类型  | 必填 | 说明 |
| :---:   | :---: | :---: | :---: |
%s
#### 类型备注
%s
####响应参数
%s
`

	structTagField = "field"
)

type (
	//自动化文档,主要用于Http请求，生成文档代码，简化开发
	AutoDoc struct {
		method string //请求方式
		url    string // 请求地址
		title  string //接口标题

		req interface{} //请求参数
		ack interface{} //响应参数

		requiredFields []string //必填字段

		requestParams []RequestParam //存储请求数据

		requestRemark string //请求备注数据

		responseString string //生成的响应数据

	}

	//请求参数
	RequestParam struct {
		Field     string //字段
		FieldType string //字段类型
		Required  bool   //必填字段
		Desc      string //字段说明
	}
)

func New(req interface{}, ack interface{}) *AutoDoc {
	this := &AutoDoc{
		url:    "接口地址",
		method: "接口请求方法",
		title:  "接口标题",
		req:    req,
		ack:    ack,
	}
	return this
}

//设置额外信息
func (this *AutoDoc) SetExtInfo(url string, method string, title string) *AutoDoc {
	this.url = url
	this.method = method
	this.title = title

	return this
}

//必填字段
func (this *AutoDoc) Require(fields ...string) *AutoDoc {
	this.requiredFields = append(this.requiredFields, fields...)

	return this
}

func (this *AutoDoc) setRequestRecursive(t reflect.Type) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("字段名称：", t.Name())
			panic(err)
		}
	}()
	var (
		requestParams []RequestParam
	)

	if t.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		//不处理
		if field.Tag.Get("json") == "-" {
			continue
		}

		fieldType := field.Type.String()

		if strings.Contains(fieldType, ".") {
			fieldTypeArr := strings.Split(fieldType, ".")
			if strings.Contains(fieldType, "[]") {
				fieldType = "[]" + fieldTypeArr[len(fieldTypeArr)-1]
			} else {
				fieldType = fieldTypeArr[len(fieldTypeArr)-1]
			}
		}

		item := RequestParam{
			Field:     field.Name,
			FieldType: fieldType,
			Required:  inArray(field.Name, this.requiredFields),
			Desc:      field.Tag.Get(structTagField),
		}

		if item.Field == item.FieldType {
			this.setRequestRecursive(field.Type)
			continue
		}

		switch field.Type.Kind() {
		case reflect.Struct:
			this.setRequestRemark(field.Type)
		case reflect.Slice, reflect.Ptr:
			this.setRequestRemark(field.Type.Elem())
		}

		if item.Desc == "" {
			item.Desc = item.Field
		}

		requestParams = append(requestParams, item)
	}

	this.requestParams = append(this.requestParams, requestParams...)

}

//执行计算
func (this *AutoDoc) Do() (content string, err error) {
	//计算requestParams
	reqT, _, _ := getStructTV(this.req)
	this.setRequestRecursive(reqT)

	//计算responseString
	ackT, _, _ := getStructTV(this.ack)
	this.responseString = this.responseStringRecursive(ackT, "", "    ", false)

	this.responseString = strings.Trim(strings.Trim(this.responseString, " "), "\n")

	if this.requestRemark != "" {
		this.requestRemark = fmt.Sprintf("\n```\n%s```\n", this.requestRemark)
	}

	content = fmt.Sprintf(template, this.title, this.url, this.method, this.getRequestParamString(), this.requestRemark, this.responseString)

	return
}

//设置备注
func (this *AutoDoc) setRequestRemark(t reflect.Type) {

	if t.Kind() == reflect.Slice {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return
	}

	this.requestRemark += fmt.Sprintf("%s:{\n", t.Name())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("json") == "-" {
			continue
		}

		tag := field.Tag.Get(structTagField)
		if tag != "" {
			tag = "//" + tag
		}
		this.requestRemark += fmt.Sprintf("    %s %s %s\n", field.Name, field.Type.String(), tag)
	}

	this.requestRemark += "}\n"

}

func (this *AutoDoc) getRequestParamString() (requestString string) {
	boolConvert := func(a bool) int {
		if a {
			return 1
		}

		return 0

	}

	//必填顺序调整
	sort.SliceStable(this.requestParams, func(i, j int) bool {
		return boolConvert(this.requestParams[i].Required) > boolConvert(this.requestParams[j].Required)
	})

	for _, v := range this.requestParams {
		required := ""
		if v.Required {
			required = "是"
		}
		requestString += fmt.Sprintf("| %s | %s | %s |  %s |\n", v.Field, v.FieldType, required, v.Desc)
	}

	return
}

//递归生成输出参数
func (this *AutoDoc) responseStringRecursive(t reflect.Type, name string, space string, embed bool) (s string) {

	if !embed {

		if t.Kind() == reflect.Slice {

			if name != "" {
				name += ":"
			}

			name += "[]"

			t = t.Elem()

			switch t.Kind() {
			case reflect.Ptr:
				t = t.Elem()
			case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Bool:

				s += fmt.Sprintf("%s%s %s \n", space, name, t.Name())

				return

			}

		} else {
			if name != "" {
				name += ":"
			}

		}

		s = fmt.Sprintf("%s%s{\n", space, name)
	}

	if t.Kind() != reflect.Struct {
		return ""
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Tag.Get("json") == "-" {
			continue
		}

		switch field.Type.Kind() {
		case reflect.Ptr:
			fieldType := field.Type.String()

			if strings.Contains(fieldType, ".") {
				fieldTypeArr := strings.Split(fieldType, ".")
				if strings.Contains(fieldType, "[]") {
					fieldType = "[]" + fieldTypeArr[len(fieldTypeArr)-1]
				} else {
					fieldType = fieldTypeArr[len(fieldTypeArr)-1]
				}
			}

			if fieldType == field.Name {
				//嵌入的结构体
				s += this.responseStringRecursive(field.Type, "", space, true)

			} else {
				s += this.responseStringRecursive(field.Type, field.Name, "    "+space, false)
			}

		case reflect.Slice, reflect.Struct, reflect.Interface:
			s += this.responseStringRecursive(field.Type, field.Name, "    "+space, false)

		case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Bool:

			tag := field.Tag.Get(structTagField)
			if tag != "" {
				tag = "//" + tag
			}
			s += fmt.Sprintf("%s%s %s %s\n", "    "+space, field.Name, field.Type.String(), tag)
		}

	}

	if !embed {
		s += space + "}\n"
	}

	return s

}

func inArray(item interface{}, array interface{}) bool {
	if reflect.TypeOf(array).Kind() != reflect.Slice {
		return false
	}

	n := reflect.ValueOf(array).Len()
	for i := 0; i < n; i++ {
		if reflect.ValueOf(array).Index(i).Interface() == reflect.ValueOf(item).Interface() {
			return true
		}
	}

	return false

}

func isStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

//判定是否为结构体指针
func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

//获取结构体或者指针的类型和值
func getStructTV(obj interface{}) (reflect.Type, reflect.Value, error) {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)

	switch {
	case isStruct(objT):
	case isStructPtr(objT):
		objT = objT.Elem()
		objV = objV.Elem()
	default:
		return objT, objV, fmt.Errorf("%v must be a struct or a struct pointer", obj)
	}

	return objT, objV, nil
}
