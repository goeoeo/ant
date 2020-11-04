package magicutil

import (
	"fmt"
	"testing"
)

func TestTranFileFromCode_Run(t *testing.T) {
	NewTranFileFromCode("locale_zh-CN.lang", "/home/yu/code/yungengxin2019/lwyapi/controllers").Run()
}

func TestCrc(t *testing.T) {
	fmt.Println(Crc("交换机厂商不存在"))
}
