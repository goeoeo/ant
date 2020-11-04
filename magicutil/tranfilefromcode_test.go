package magicutil

import "testing"

func TestTranFileFromCode_Run(t *testing.T) {
	NewTranFileFromCode("locale_zh-CN.lang", "/home/yu/code/yungengxin2019/lwyapi/controllers").Run()
}
