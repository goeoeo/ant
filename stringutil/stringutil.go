package stringutil

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//判定某个值是否在数组里面
func InSliceString(field string, arr []string) bool {
	for _, v := range arr {
		if v == field {
			return true
		}
	}

	return false
}

// 生成md5
func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//ip转换函数,字符串转数字
func Ip2Long(ipstr string) (ip uint32) {
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return
	}
	ips := reg.FindStringSubmatch(ipstr)
	if ips == nil {
		return
	}

	ip1, _ := strconv.Atoi(ips[1])
	ip2, _ := strconv.Atoi(ips[2])
	ip3, _ := strconv.Atoi(ips[3])
	ip4, _ := strconv.Atoi(ips[4])

	if ip1 > 255 || ip2 > 255 || ip3 > 255 || ip4 > 255 {
		return
	}

	ip += uint32(ip1 * 0x1000000)
	ip += uint32(ip2 * 0x10000)
	ip += uint32(ip3 * 0x100)
	ip += uint32(ip4)

	return
}

//ip转换函数,数字转字符串
func Long2Ip(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}

//生成随机字符串
func RandomString(num int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, num)
	rand.Seed(time.Now().Unix())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}


//保留左边0的字符串加法
//001+1=002
func Keep0Add(s string, sep int) string {
	//必须全数字
	re := regexp.MustCompile(`^\d+$`)
	if !re.MatchString(s) {
		return ""
	}

	sl := len(s)

	s = strings.TrimLeft(s, "0")

	sint, err := strconv.Atoi(s)
	if err != nil {
		return ""
	}

	sint += sep

	ns := strconv.Itoa(sint)
	nsl := len(ns)
	if sl > nsl {
		//新字符串长度小于旧字符串长度
		//左边补0
		var s0 string
		for i := 0; i < (sl - nsl); i++ {
			s0 += "0"
		}

		return s0 + ns
	}

	return ns
}