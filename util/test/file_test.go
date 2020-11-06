package test

import (
	"fmt"
	"github.com/phpdi/ant/util"
	"testing"
)

func TestScanPath(t *testing.T) {
	res, err := util.ScanPath("/home/yu/code/ant", 1)
	if err != nil {
		t.Fatal(err)
	}

	for k := range res {
		fmt.Println(k)
	}
}

func TestScanDir(t *testing.T) {
	res := util.ScanDir("/home/yu/code/ant", 1)

	for _, v := range res {
		fmt.Println(v)
	}
}

func TestParseChnFromGolang(t *testing.T) {
	util.ParseChnFromGolang("/home/yu/code/ant/util/file.go")
}
