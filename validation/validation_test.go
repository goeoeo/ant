package validation

import (
	"ant"
	"testing"
)

func TestValidation_Require(t *testing.T) {
	a:=ant.StockHsas{}
	err:=NewValidation().Require("Code").Valid(a)
	if err != nil {
		t.Error(err)
	}


}

func TestValidation_Max(t *testing.T) {



}