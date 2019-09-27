package reflectutil

import (
	"fmt"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	a:= ""
	fmt.Println(IsEmpty(a))
}

func TestInArray(t *testing.T) {
	a:="1111"
	b:=[]string{"111","222"}
	fmt.Println(InArray(a,b))
}