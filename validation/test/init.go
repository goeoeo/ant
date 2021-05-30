package test

import (
	"github.com/phpdi/ant/validation"
	"regexp"
)

var defaultValidationConfig *validation.ValidationConfig

//定义示例
func init() {
	//注册函数
	defaultValidationConfig = validation.NewValidationConfig()

	//扩展函数
	defaultValidationConfig.
		RegisterFun("OpenTaskName", OpenTaskName, "只支持数字,字母,汉字,-或_或.的组合").
		RegisterFun("Chn", Chn, "只支持汉字")
}

func NewValidate() *validation.Validation {
	return validation.NewValidation(defaultValidationConfig)
}

//中文,数字,字母,下划线,点
func OpenTaskName(validValue interface{}, params ...string) bool {
	if str, ok := validValue.(string); ok {
		re := regexp.MustCompile("^[\u4e00-\u9fa50-9a-zA-Z_.-]+$")
		return re.MatchString(str)
	}

	return false

}

//中文
func Chn(validValue interface{}, params ...string) bool {
	if str, ok := validValue.(string); ok {
		re := regexp.MustCompile("^[\u4e00-\u9fa50]+$")
		return re.MatchString(str)
	}

	return false
}
