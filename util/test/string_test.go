package test

import (
	"github.com/phpdi/ant/util"
	"testing"
)

func TestMd5(t *testing.T) {
	str := util.Md5([]byte("123456"))
	if str != "e10adc3949ba59abbe56e057f20f883e" {
		t.Error("md5 fail")
	}
}

func TestKeep0Add(t *testing.T) {
	s := "001"
	ns := util.Keep0Add(s, 2)
	if ns != "003" {
		t.Error("fail")
	}
}
