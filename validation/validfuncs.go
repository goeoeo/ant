package validation

import (
	"strconv"
)

//验证函数

//验证函数未通过,对应的错误提示
var MessageTmpls = map[string]string{
	"Required":          "不能为空",
	"Max":          "最大为%d",
}


//限制数字最大值
func Max(validValue interface{},params ...interface{}) bool {
	var max int
	var err error
	if len(params)!= 1 {
		return false
	}

	if maxStr,ok:=params[0].(string); ok {
		max,err=strconv.Atoi(maxStr)
		if err != nil {
			return false
		}

	}

	var v int

	switch tmp:=validValue.(type) {
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

	return v<=max
}