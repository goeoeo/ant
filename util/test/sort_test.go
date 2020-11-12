package test

import (
	"fmt"
	"github.com/phpdi/ant/util"
	"testing"
)

type baseInfo struct {
	Machine string
	Admin   uint32
	Ip      int

	B struct {
		Id   int
		Name string
	}
	Hi Hi
}

type Hi struct {
	Id   int
	Name string
}

func Test_SortSlice(t *testing.T) {

	clients := []baseInfo{
		{Machine: "BBB", Admin: 0, Ip: 100},
		{Machine: "CCB", Admin: 0, Ip: 109},
		{Machine: "DDD", Admin: 1, Ip: 103},
		{Machine: "ZZZ", Admin: 0, Ip: 108},
		{Machine: "EEE", Admin: 0, Ip: 108},
		{Machine: "WWW", Admin: 0, Ip: 105},
		{Machine: "XXX", Admin: 1, Ip: 106},
		{Machine: "AAA", Admin: 0, Ip: 107},
	}

	for _, c := range clients {
		fmt.Println(c)
	}

	fmt.Println("__________________________")
	if err := util.SortSlice(&clients, "Admin aes,Ip desc"); err != nil {
		t.Error(err)
	}

	for _, c := range clients {
		fmt.Println(c)
	}

}

//点预防排序取值
func Test_SortSlice1(t *testing.T) {
	clients := []baseInfo{
		{Hi: Hi{
			Id:   1,
			Name: "aaa",
		}}, {Hi: Hi{
			Id:   2,
			Name: "cc",
		}}, {Hi: Hi{
			Id:   2,
			Name: "bb",
		}},
	}

	if err := util.SortSlice(&clients, "Hi.Id desc,Hi.Name aes"); err != nil {
		t.Error(err)
	}

	for _, c := range clients {
		fmt.Println(c.Hi)
	}

	client1s := []baseInfo{
		{B: struct {
			Id   int
			Name string
		}{Id: 1, Name: "aaa"}},
		{B: struct {
			Id   int
			Name string
		}{Id: 2, Name: "cc"}},
		{B: struct {
			Id   int
			Name string
		}{Id: 2, Name: "bb"}},
	}

	if err := util.SortSlice(&client1s, "B.Id desc,B.Name aes"); err != nil {
		t.Error(err)
	}

	for _, c := range client1s {
		fmt.Println(c.B)
	}
}
