package validation

import (
	"fmt"
	"github.com/phpdi/ant"
	"testing"
)

func TestValidation_Require(t *testing.T) {
	a := ant.StockHsas{}

	v := New()

	//零值,必填,验证
	err := v.Require("Code").Valid(a)
	if err == nil {
		t.Error("零值,必填验证,失败")

	}
	t.Log("零值,必填验证,通过")

	a.Code = "111"
	err = v.Valid(a)
	if err != nil {
		t.Error("非零值,必填验证,失败")
	}
	t.Log("非零值,必填验证,通过")

}

func TestValidation_parseFunc(t *testing.T) {
	res := New().parseFunc("Name(Id);Max(20);Min(5)")
	fmt.Printf("%+v", res)
}

func TestNewValidation(t *testing.T) {
	v := New()
	v.Config.SetMessageTmpls(map[string]string{
		"Max": "max is %v",
	})

	t.Log(v.Config.messageTmpls["Max"])
}

func TestValidation_SetFailMessages(t *testing.T) {
	a := ant.StockHsas{Id: 100}
	v := New().SetFailMessages(map[string]string{"Id": "Id 必须到5到20之间"})
	err := v.Valid(a)
	if err != nil {
		t.Error(err)
	}
}

func TestValidation_Require2(t *testing.T) {
	a := ant.StockHsas{Id: 100}
	v := New().Require("Code")
	err := v.Valid(a)
	if err != nil {
		t.Error(err)
	}
}
