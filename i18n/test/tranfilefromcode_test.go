package test

import (
	"fmt"
	"github.com/phpdi/ant/i18n"
	"testing"
)

func TestTranFileFromCode_Run(t *testing.T) {
	i18n.NewTranFileFromCode("locale_zh-CN.lang", "/home/yu/code/yungengxin2019/lwyapi/controllers").Run()
}

func TestCrc(t *testing.T) {
	fmt.Println(i18n.Crc("交换机厂商不存在"))
}
