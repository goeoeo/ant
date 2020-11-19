package test

import (
	"testing"
)

func TestValidation_Require(t *testing.T) {
	type C struct {
		Id int `valid:"Max(20);Min(5)"`
	}

	type B struct {
		Code string
		CC   C

		C
	}

	a := B{}
	a.Id = 5
	a.CC.Id = 6

	v := NewValidate()

	//零值,必填,验证
	err := v.Require("Code", "C.Id").Valid(a)
	if err == nil {
		t.Fatal("零值,必填验证,失败:", err)
	}

	a.Code = "111"
	err = v.Valid(a)
	if err != nil {
		t.Fatal("非零值,必填验证,失败:", err)
	}

	a.CC.Id = 8
	v = NewValidate().Require("CC.Id")
	err = v.Valid(a)
	if err != nil {
		t.Fatal("零值,必填验证,失败:", err)
	}

	a.CC.Id = 1
	err = v.Valid(a)
	if err == nil {
		t.Fatal("非零值,必填验证,失败:", err)
	}

}

func TestNewValidation(t *testing.T) {
	v := NewValidate()
	v.Config.SetMessageTmpls(map[string]string{
		"Max": "max is %v",
	})

}

func TestValidation_SetFailMessages(t *testing.T) {
	a := StockHsas{Id: 100}
	v := NewValidate().SetFailMessages(map[string]string{"Id": "Id 必须到5到20之间"})
	err := v.Valid(a)
	if err != nil {
		t.Error(err)
	}
}

func TestValidation_Require2(t *testing.T) {
	a := StockHsas{Id: 1}
	a.Sp.Id = 1
	v := NewValidate().Require("Sp.Id")
	err := v.Valid(a)
	if err != nil {
		t.Error(err)
	}
}
