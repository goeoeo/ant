package util

import (
	"fmt"
	"testing"
)

func TestScanPath(t *testing.T) {
	res,err:=ScanPath("/home/yu/code/ant",1)
	if err != nil {
		t.Fatal(err)
	}

	for k:=range res {
		fmt.Println(k)
	}
}

func TestScanDir(t *testing.T) {
	res:=ScanDir("/home/yu/code/ant",1)

	for _,v:=range res {
		fmt.Println(v)
	}
}
