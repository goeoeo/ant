package util

import (
	"log"
	"runtime"
	"strings"
)

func PosLog(err error) {
	var (
		s string
	)
	if err == nil {
		return
	}

	pc, _, _, _ := runtime.Caller(1)
	a := runtime.FuncForPC(pc).Name()

	arr := strings.Split(a, "/")

	if len(arr) > 0 {
		s = arr[len(arr)-1]
	}

	log.Printf("%s 错误: %s\n", s, err.Error())

}
