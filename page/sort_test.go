package page

import (
	"fmt"
	"testing"
)

func Test_SortSlice(t *testing.T)  {

	type baseInfo struct {
		Machine string
		Admin uint32
		Ip int
	}

	clients := []baseInfo{
		{Machine:"BBB", Admin:0, Ip:100},
		{Machine:"CCB", Admin:0, Ip:109},
		{Machine:"DDD", Admin:1, Ip:103},
		{Machine:"ZZZ", Admin:0, Ip:108},
		{Machine:"EEE", Admin:0, Ip:108},
		{Machine:"WWW", Admin:0, Ip:105},
		{Machine:"XXX", Admin:1, Ip:106},
		{Machine:"AAA", Admin:0, Ip:107},
	}

	for _, c := range clients {
		fmt.Println(c)
	}

	fmt.Println("__________________________")
	if err:=SortSlice(&clients,"Admin desc,Ip desc,Machine aes");err!=nil {
		t.Error(err)
	}

	for _, c := range clients {
		fmt.Println(c)
	}


}