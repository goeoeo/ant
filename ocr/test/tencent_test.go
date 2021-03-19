package test

import (
	"fmt"
	ocr2 "github.com/phpdi/ant/ocr"
	"testing"
)


func TestUniversal(t *testing.T)  {
	ocr:=ocr2.NewTencentOcr(2160713637,"MEIgU2fna3OS8WyU")

	res,err:=ocr.Universal("/home/yu/code/clockin/screen.png")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}