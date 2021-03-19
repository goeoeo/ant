package ocr

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	r "math/rand"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

func Post(url string,postParams map[string]interface{},ack interface{}) (err error) {

	var (
		postStr   string
		cookieJar *cookiejar.Jar
		req       *http.Request
		resp      *http.Response
		bodyByte []byte

	)

	for k,v:=range postParams {
		postStr+=fmt.Sprintf("%s=%v&",k,v)
	}

	postStr=strings.TrimRight(postStr,"&")

	tr := &http.Transport{
		DisableKeepAlives: true,                                  //禁用长连接
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true}, //跳过证书验证
	}
	//http cookie接口
	if cookieJar, err = cookiejar.New(nil); err != nil {
		return
	}

	client := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}

	if req, err = http.NewRequest("POST", url, strings.NewReader(postStr)); err != nil {
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
	fmt.Println(string(bodyByte))

	if ack!=nil && len(bodyByte)> 0 {
		if err=json.Unmarshal(bodyByte,ack);err!= nil {
			return
		}
	}

	return
}

// RandomCreateBytes generate random []byte by specify chars.
func RandomString(n int, alphabets ...byte) string {
	if len(alphabets) == 0 {
		alphabets = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)
	}
	var bytes = make([]byte, n)
	var randBy bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randBy = true
	}
	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[r.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return string(bytes)
}

//json方式打印结构体
func JsonPrint(obj interface{}) {
	tmp, _ := json.MarshalIndent(obj, "", "     ")
	fmt.Println(string(tmp))
}



