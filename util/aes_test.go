package util

import (
	"testing"
)

var key = []byte("6aaf508205a7e5cca6b03ee5747a92f3")

func TestAESEncrypt(t *testing.T) {

	// 每次得到的结果都不同，但是都可以解密
	msg, ok := AESEncrypt(key, "abcd")
	if !ok {
		t.Error("加密失败")
	}
	t.Log("msg=", msg)

}

func TestAESDecrypt(t *testing.T) {

	// 每次得到的结果都不同，但是都可以解密
	msg, ok := AESDecrypt(key, "3YRtSZVmlGk3KH62yJyBwdBgFb4=")
	if !ok {
		t.Error("加密失败")
	}
	t.Log("msg=", msg)

}
