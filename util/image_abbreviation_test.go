package util

import "testing"

func TestMakeThumbnail(t *testing.T) {
	i := NewImageAbbreviation(320, 240)
	i.MakeThumbnail("/home/yu/图片/IMG_20200115_173549.jpg", "/home/yu/图片/1.jpg")
}
