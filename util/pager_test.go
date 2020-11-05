package util

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestPager_Pagination(t *testing.T) {
	type User struct {
		Id   int
		Name int
	}

	total := int64(0)
	users := []User{}

	for i := 0; i < 100; i++ {
		users = append(users, User{i, i})
	}

	page := &Pager{
		Page:     2,
		PageSize: 1000,
	}

	if err := page.Pagination(&users).Total(&total).Error; err != nil {
		t.Error(err)
		return
	}

	fmt.Println("total:", total)
	fmt.Println("users:", len(users))
	for _, v := range users {
		fmt.Println("id:", v.Id)
	}

}

func TestPager_A(t *testing.T) {
	type User struct {
		Id   int
		Name int
	}
	users := []User{}
	for i := 0; i < 10; i++ {

		item := User{
			Id:   i,
			Name: i,
		}

		users = append(users, item)
	}

	for k := range users {
		fmt.Println(">>>", unsafe.Sizeof(users[k]))
	}
	a(&users)

}

type User struct {
	Id   int
	Name string
}

func a(obj interface{}) {
	v := reflect.TypeOf(obj).Elem().Elem()

	fmt.Println(v.Size())

}
