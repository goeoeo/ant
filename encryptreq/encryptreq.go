package encryptreq

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type (
	EncryptReq struct {
		debug bool
	}

	//http参数
	HttpOption struct {
		TimeOut           int  //请求server超时时间 单位秒
		DisableKeepAlives bool //禁用长链接
	}

	HttpOptionFun func(options *HttpOption)

	//响应的数据格式
	Response struct {
		Code uint32      //错误码
		Msg  string      //错误信息
		Data interface{} //响应数据
	}

	//请求的数据格式
	Request struct {
		Code string //认证字符串
		Time string //请求的时间
		Data string //传输的数据：请求数据的json格式数据
	}
)

//设置timeout参数
func WithTimeout(second int) HttpOptionFun {
	return func(options *HttpOption) {
		options.TimeOut = second
	}
}

//解密请求数据
func (this *EncryptReq) DecryptRequest(data interface{}, req Request) (err error) {
	var (
		aesKey      string
		decryptData string
		ok          bool
	)
	if data == nil {
		return
	}

	if !this.isStructPtr(data) {
		return errors.New("输出结构体必须为指针")
	}

	aesKey = this.getAesKey(req.Code, req.Time)

	if decryptData, ok = AESDecrypt([]byte(aesKey), req.Data); !ok {
		return errors.New("数据解密失败")
	}

	//输出
	if err = json.Unmarshal([]byte(decryptData), data); err != nil {
		return
	}

	return
}

func (this *EncryptReq) Debug() *EncryptReq {
	this.debug = true
	return this
}

//加密请求数据
func (this *EncryptReq) encryptRequest(data interface{}) (req Request, err error) {
	var (
		jsonData []byte
		ok       bool
	)
	if data == nil {
		return req, errors.New("请求数据不能为空")

	}

	req = Request{
		Code: this.randomString(22),
		Time: strconv.Itoa(int(time.Now().Unix())),
	}

	aesKey := this.getAesKey(req.Code, req.Time)

	if jsonData, err = json.Marshal(data); err != nil {
		return
	}

	//加密数据
	if req.Data, ok = AESEncrypt([]byte(aesKey), string(jsonData)); !ok {
		return req, errors.New("数据加密失败")
	}

	return
}

//进行加密请求
func (this *EncryptReq) EncryptPost(url string, requestData interface{}, resp *Response, httpOptionFuns ...HttpOptionFun) (err error) {

	var (
		req      Request
		reqByte  []byte
		bodyByte []byte
	)
	//加密请求内容
	if req, err = this.encryptRequest(requestData); err != nil {
		return
	}

	//序列化请求数据
	if reqByte, err = json.Marshal(req); err != nil {
		return
	}

	//发送请求
	if bodyByte, err = this.httpPost(url, reqByte, httpOptionFuns...); err != nil {
		return
	}

	if this.debug {
		fmt.Println("请求地址：" + url)
		fmt.Printf("请求数据:%+v \n", requestData)
		fmt.Println("请求加密数据", string(reqByte))
		fmt.Println("响应数据：" + string(bodyByte))
	}

	//解析响应
	if err = json.Unmarshal(bodyByte, &resp); err != nil {
		return
	}

	return nil
}

//生成随机字符串
func (this *EncryptReq) randomString(num int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, num)
	rand.Seed(time.Now().Unix())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//获取aes加密key
func (this *EncryptReq) getAesKey(code string, time string) string {

	s := code + time
	strLen := len(s)

	evenStr := []byte{}

	for i := 0; i < strLen; i++ {
		if i%2 != 0 {
			evenStr = append(evenStr, []byte(s)[i])
		}
	}

	tmp := append(evenStr, []byte(s)...)

	return string(tmp[:16])
}

//httpPost请求
func (this *EncryptReq) httpPost(url string, postParams []byte, httpOptionFuns ...HttpOptionFun) (bodyByte []byte, err error) {

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

//json，协议byte数据
func (this *EncryptReq) PostCheckCode(url string, req interface{}, ack interface{}, httpOptionFuns ...HttpOptionFun) error {
	var (
		err      error
		response Response
	)

	if ack != nil && !this.isStructPtr(ack) {
		return errors.New("ack必须为指针")
	}

	response.Data = ack

	if err = this.EncryptPost(url, req, &response, httpOptionFuns...); err != nil {
		return err
	}

	if response.Code != 0 {
		return errors.New(response.Msg)
	}

	return nil
}

//判定是否为结构体指针
func (this *EncryptReq) isStructPtr(obj interface{}) bool {
	t := reflect.TypeOf(obj)
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}
