package validation

/**
import (
	"regexp"
)

//定义示例
func init() {
	//注册函数
	DefultValidationConfig = NewValidationConfig()

	DefultValidationConfig.
		RegisterFun("OpenTaskName", OpenTaskName, "只支持数字,字母,汉字,-或_或.的组合").
		RegisterFun("Chn", Chn, "只支持汉字")
}

func NewValidate() *Validation {
	return New()
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

**/
