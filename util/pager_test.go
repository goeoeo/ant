package util

import (
	"fmt"
	"testing"
)

func TestPager_Pagination(t *testing.T) {
	type User struct {
		Id   int
		Name int
	}

	total := int64(0)
	users := []User{}

	for i := 0; i <= 100; i++ {
		users = append(users, User{i, i})
	}

	page := &Pager{
		Page:     2,
		PageSize: 50,
	}

	if err := page.Pagination(&users).Total(&total).Error; err != nil {
		t.Error(err)
		return
	}

	fmt.Println("total:", total)
	fmt.Println("users:", len(users))
	if len(users) > 0 {
		fmt.Println(users[0].Id)
	}

}
