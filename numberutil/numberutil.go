package numberutil

import (
	"math/rand"
	"time"
)

//判定某个值是否在数组里面
func InSliceInt(field int, arr []int) bool {
	for _, v := range arr {
		if v == field {
			return true
		}
	}

	return false
}


//生成随机数字
func RandomNumber(maxNum int) int {
	rand.Seed(time.Now().UnixNano()) //以时间作为初始化种子
	return rand.Intn(maxNum)
}
