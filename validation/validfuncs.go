package validation

import (
	"reflect"
	"regexp"
	"strconv"
	"unicode/utf8"
)

//验证函数，通过ValidationConfig.RegisterFun进行注册

//限制数字最大值
func Max(validValue interface{}, params ...string) bool {
	var max int
	var err error
	if len(params) != 1 {
		return false
	}

	max, err = strconv.Atoi(params[0])
	if err != nil {
		return false
	}

	var v int

	switch tmp := validValue.(type) {
	case int64:
		v = int(tmp)
	case int32:
		v = int(tmp)
	case int16:
		v = int(tmp)
	case int8:
		v = int(tmp)
	case int:
		v = int(tmp)
	case uint64:
		v = int(tmp)
	case uint32:
		v = int(tmp)
	case uint16:
		v = int(tmp)
	case uint8:
		v = int(tmp)
	case uint:
		v = int(tmp)
	default:
		return false
	}

	return v <= max
}

//限制数字最小值
func Min(validValue interface{}, params ...string) bool {
	var min int
	var err error
	if len(params) != 1 {
		return false
	}

	min, err = strconv.Atoi(params[0])
	if err != nil {
		return false
	}

	var v int

	switch tmp := validValue.(type) {
	case int64:
		v = int(tmp)
	case int32:
		v = int(tmp)
	case int16:
		v = int(tmp)
	case int8:
		v = int(tmp)
	case int:
		v = int(tmp)
	case uint64:
		v = int(tmp)
	case uint32:
		v = int(tmp)
	case uint16:
		v = int(tmp)
	case uint8:
		v = int(tmp)
	case uint:
		v = int(tmp)
	default:
		return false
	}

	return v >= min
}

//范围
func Range(validValue interface{}, params ...string) bool {
	var min int
	var max int
	var err error
	if len(params) != 2 {
		return false
	}

	min, err = strconv.Atoi(params[0])
	if err != nil {
		return false
	}

	max, err = strconv.Atoi(params[1])
	if err != nil {
		return false
	}

	var v int

	switch tmp := validValue.(type) {
	case int64:
		v = int(tmp)
	case int32:
		v = int(tmp)
	case int16:
		v = int(tmp)
	case int8:
		v = int(tmp)
	case int:
		v = int(tmp)
	case uint64:
		v = int(tmp)
	case uint32:
		v = int(tmp)
	case uint16:
		v = int(tmp)
	case uint8:
		v = int(tmp)
	case uint:
		v = int(tmp)
	default:
		return false
	}

	if v >= min && v <= max {
		return true
	}

	return false
}

//最小长度 有效类型：string slice
func MinSize(validValue interface{}, params ...string) bool {
	if len(params) != 1 {
		return false
	}

	var min int
	min, err := strconv.Atoi(params[0])
	if err != nil {
		return false
	}

	if str, ok := validValue.(string); ok {
		return utf8.RuneCountInString(str) >= min
	}
	v := reflect.ValueOf(validValue)
	if v.Kind() == reflect.Slice {
		return v.Len() >= min
	}
	return false
}

//最大长度,有效类型：string slice，
func MaxSize(validValue interface{}, params ...string) bool {
	if len(params) != 1 {
		return false
	}

	var max int
	max, err := strconv.Atoi(params[0])
	if err != nil {
		return false
	}

	if str, ok := validValue.(string); ok {
		return utf8.RuneCountInString(str) <= max
	}
	v := reflect.ValueOf(validValue)
	if v.Kind() == reflect.Slice {
		return v.Len() <= max
	}
	return false
}

//指定长度，有效类型：string slice
func Length(validValue interface{}, params ...string) bool {
	if len(params) != 1 {
		return false
	}

	var lenNum int
	lenNum, err := strconv.Atoi(params[0])
	if err != nil {
		return false
	}

	if str, ok := validValue.(string); ok {
		return utf8.RuneCountInString(str) == lenNum
	}
	v := reflect.ValueOf(validValue)
	if v.Kind() == reflect.Slice {
		return v.Len() == lenNum
	}
	return false
}

//alpha字符(全字母)，有效类型：string
func Alpha(validValue interface{}, params ...string) bool {
	if str, ok := validValue.(string); ok {
		for _, v := range str {
			if ('Z' < v || v < 'A') && ('z' < v || v < 'a') {
				return false
			}
		}
		return true
	}
	return false
}

//数字，有效类型：string
func Numeric(validValue interface{}, params ...string) bool {
	if str, ok := validValue.(string); ok {
		for _, v := range str {
			if '9' < v || v < '0' {
				return false
			}
		}
		return true
	}
	return false
}

//alpha 字符(字母)或数字，有效类型：string
func AlphaNumeric(validValue interface{}, params ...string) bool {
	if str, ok := validValue.(string); ok {
		for _, v := range str {
			if ('Z' < v || v < 'A') && ('z' < v || v < 'a') && ('9' < v || v < '0') {
				return false
			}
		}
		return true
	}
	return false
}

//alpha 字符或数字或横杠 -_，有效类型：string，
func AlphaDash(validValue interface{}, params ...string) bool {
	if str, ok := validValue.(string); ok {
		alphaDashPattern := regexp.MustCompile(`^[\d\w-_]+$`)
		return alphaDashPattern.MatchString(str)
	}
	return false
}

// 邮箱格式，有效类型：string
func Email(validValue interface{}, params ...string) bool {
	if str, ok := validValue.(string); ok {
		emailPattern := regexp.MustCompile(`^[\w!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\.[\w!#$%&'*+/=?^_` + "`" + `{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[a-zA-Z0-9](?:[\w-]*[\w])?$`)
		return emailPattern.MatchString(str)
	}
	return false
}

// IP 格式，目前只支持 IPv4 格式验证，有效类型：string
func IP(validValue interface{}, params ...string) bool {
	if str, ok := validValue.(string); ok {
		ipPattern := regexp.MustCompile(`^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$`)
		return ipPattern.MatchString(str)
	}
	return false
}

// 手机号，有效类型：string
func Mobile(validValue interface{}, params ...string) bool {
	if str, ok := validValue.(string); ok {
		mobilePattern := regexp.MustCompile(`^((\+86)|(86))?(1([3456789]))\d{9}$`)
		return mobilePattern.MatchString(str)
	}
	return false
}

//固定电话号，有效类型：string
func Tel(validValue interface{}, params ...string) bool {
	if str, ok := validValue.(string); ok {
		telPattern := regexp.MustCompile(`^(0\d{2,3}(\-)?)?\d{7,8}$`)
		return telPattern.MatchString(str)
	}
	return false
}

//邮政编码，有效类型：string
func ZipCode(validValue interface{}, params ...string) bool {
	if str, ok := validValue.(string); ok {
		zipCodePattern := regexp.MustCompile(`^[1-9]\d{5}$`)
		return zipCodePattern.MatchString(str)
	}
	return false
}

//mac地址校验
func Mac(validValue interface{}, params ...string) bool {
	if str, ok := validValue.(string); ok {
		macReg := regexp.MustCompile(`^([A-Z0-9]{2}-){5}[A-Z0-9]{2}$`)
		return macReg.MatchString(str)
	}
	return false
}

//中文,数字,字母,下划线
func ChnDash(validValue interface{}, params ...string) bool {
	if str, ok := validValue.(string); ok {
		re := regexp.MustCompile("^[\u4e00-\u9fa50-9a-zA-Z_-]+$")
		return re.MatchString(str)
	}

	return false

}

//中文,数字,字母
func ChnAlphaNumeric(validValue interface{}, params ...string) bool {
	if str, ok := validValue.(string); ok {
		re := regexp.MustCompile("^[\u4e00-\u9fa50-9a-zA-Z]+$")
		return re.MatchString(str)
	}

	return false
}
