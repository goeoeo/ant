package captcha

import (
	"math/rand"
	"time"
)


//随机字符串
func getRandStr(n int) (randStr string) {
	chars := "ABCDEFGHIJKMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz23456789"
	charsLen := len(chars)
	if n > 10 {
		n = 10
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		randIndex := rand.Intn(charsLen)
		randStr += chars[randIndex : randIndex+1]
	}
	return randStr
}


//创建图形验证码
func CreateCode(n int, size ...int) (text string, imgByte []byte) {
	text = getRandStr(n)
	width := 180
	height := 60
	if len(size) >= 2 {
		width = size[0]
		height = size[1]
	}

	imgByte = ImgText(width, height, text)

	return text, imgByte
}
