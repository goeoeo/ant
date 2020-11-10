package test

import (
	"fmt"
	"github.com/phpdi/ant/i18n"
	"testing"
)

func TestNewI18nWithDir(t *testing.T) {
	i18 := i18n.NewI18nWithDir("")

	fmt.Println(i18.Tr("zh-CN", ""))
}
