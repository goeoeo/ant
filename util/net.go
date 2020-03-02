package util

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os/exec"
	"strings"
	"time"
)

type (
	HttpOptionFun func(options *HttpOption)

	//http参数
	HttpOption struct {
		TimeOut           int  //请求server超时时间 单位秒
		DisableKeepAlives bool //禁用长链接
	}
)

//检测网络状态
func NetWorkStatus(ip string) bool {
	cmd := exec.Command("ping", ip, "-c", "1", "-W", "5")
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

//httpPost请求
func HttpPost(url string, postParams []byte, httpOptionFuns ...HttpOptionFun) (bodyByte []byte, err error) {
	var (
		cookieJar *cookiejar.Jar
		req       *http.Request
		resp      *http.Response

		option HttpOption
	)
	//http默认参数
	option = HttpOption{
		TimeOut:           5,
		DisableKeepAlives: true,
	}
	for _, f := range httpOptionFuns {
		f(&option)
	}

	tr := &http.Transport{
		DisableKeepAlives: option.DisableKeepAlives,              //禁用长连接
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true}, //跳过证书验证
	}
	//http cookie接口
	if cookieJar, err = cookiejar.New(nil); err != nil {
		return
	}

	client := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
		Timeout:   time.Duration(option.TimeOut) * time.Second,
	}

	if req, err = http.NewRequest("POST", url, strings.NewReader(string(postParams))); err != nil {
		return
	}

	req.Header.Set("Connection", "close")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if resp, err = client.Do(req); err != nil {
		return
	}
	resp.Header.Set("Connection", "close")
	defer resp.Body.Close()

	if bodyByte, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	return
}
