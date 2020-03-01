package util

import (
	"math/rand"
	"time"
)

//生成随机数字
func RandomNumber(maxNum int) int {
	rand.Seed(time.Now().UnixNano()) //以时间作为初始化种子
	return rand.Intn(maxNum)
}
