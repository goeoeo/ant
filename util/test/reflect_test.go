package test

import (
	"fmt"
	"github.com/phpdi/ant/util"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	a := ""
	fmt.Println(util.IsEmpty(a))
}

func TestInArray(t *testing.T) {
	a := "1111"
	b := []string{"111", "222"}
	fmt.Println(util.InArray(a, b))
}
