package validation

import (
	"encoding/json"
	"fmt"
	"strings"
)

//验证配置
type (
	//验证方法函数
	ValidFun func(validValue interface{}, params ...string) bool

	//翻译函数回调
	TranFun func(format string, params ...interface{}) string

	//验证器基础配置
	ValidationConfig struct {
		RequiredField   string //必填验证校验函数名
		StructTagField  string //结构体 验证器structTag名称 valid
		StructFieldName string //结构体字段名称 field

		validFuns map[string]ValidFun //验证函数

		messageTmpls map[string]string //验证失败函数对应的模板消息

		TranFunc TranFun
	}
)

//生成默认的配置
func NewValidationConfig() *ValidationConfig {
	this := &ValidationConfig{
		RequiredField:   "Required",
		StructTagField:  "valid",
		StructFieldName: "field",
		validFuns:       make(map[string]ValidFun),
		messageTmpls:    make(map[string]string),
		TranFunc: func(format string, params ...interface{}) string {
			num := strings.Count(format, `%v`)
			if num == 0 {
				return format
			}
			return fmt.Sprintf(format, params...)
		},
	}

	//注册函数
	this.RegisterFun("Required", Required, "不能为空").
		RegisterFun("Max", Max, "最大为%v").
		RegisterFun("Min", Min, "最小为%v").
		RegisterFun("Range", Range, "范围为%v到%v").
		RegisterFun("MinSize", MinSize, "最小长度为%v").
		RegisterFun("MaxSize", MaxSize, "最大长度为%v").
		RegisterFun("Length", Length, "长度必须为%v").
		RegisterFun("Alpha", Alpha, "必须为字母").
		RegisterFun("Numeric", Numeric, "必须为数字").
		RegisterFun("AlphaNumeric", AlphaNumeric, "必须为字母、数字").
		RegisterFun("AlphaDash", AlphaDash, "必须为字母、数字、-或_").
		RegisterFun("Email", Email, "必须为有效的邮箱地址").
		RegisterFun("IP", IP, "必须为有效的IP地址").
		RegisterFun("Mobile", Mobile, "必须为有效的手机号码").
		RegisterFun("Tel", Tel, "必须为有效的固定电话").
		RegisterFun("ZipCode", ZipCode, "必须是有效的邮政编码").
		RegisterFun("Mac", Mac, "必须是有效的mac地址").
		RegisterFun("ChnDash", ChnDash, "必须是数字,字母,汉字,-或_的组合").
		RegisterFun("ChnAlphaNumeric", ChnAlphaNumeric, "必须是数字,字母,汉字的组合").
		RegisterFun("NumericDot", NumericDot, "必须是数字,点的组合")

	return this
}

//注册函数
func (this *ValidationConfig) RegisterFun(funcName string, validFunc ValidFun, failMsg ...string) *ValidationConfig {
	if len(this.validFuns) == 0 {
		this.validFuns = make(map[string]ValidFun)
	}
	this.validFuns[funcName] = validFunc

	if len(failMsg) > 0 {
		this.SetMessageTmpls(map[string]string{funcName: failMsg[0]})
	}

	return this
}

//批量设置模板消息
func (this *ValidationConfig) SetMessageTmpls(messageTmpls map[string]string) *ValidationConfig {

	for k, v := range messageTmpls {
		this.messageTmpls[k] = v
	}

	return this
}

//打印出所有支持的函数
func (this *ValidationConfig) SupportFuns() {
	tmp, _ := json.MarshalIndent(this.messageTmpls, "", "     ")
	fmt.Println(string(tmp))
}

func (this *ValidationConfig) GetMessageTmpls(key string) string {
	return this.messageTmpls[key]
}
