package services

import (
	"ant/reflectutil"
	"ant/stringutil"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)


//底层错误页面拒绝显示的Code
var RejectCode = []int{}

//lwagent 接口服务服务
type lwagentService struct {
	Debug bool
}

//响应的数据格式
type Response struct {
	Code int         `field:"错误码" json:"code"`
	Msg  string      `field:"错误信息" json:"msg"`
	Data interface{} `field:"交互数据" json:"data"`
}

//请求的数据格式
type Request struct {
	Code string `field:"认证字符串"`
	Time string `field:"请求的时间"`
	Data string `field:"传输的数据"`
}

//获取数据
func (this *lwagentService) Post(url string,requestData interface{}, response *Response) error {

	req := httplib.Post(url).SetTimeout(3*time.Second, 10*time.Second)

	
	//包装请求数据
	request,err:=this.makeRequest(requestData)
	if err != nil {
		return err
	}
	
	requestStr, _ := json.Marshal(request)
	req.Body(requestStr)

	bodyByte, err := req.Bytes()

	if this.Debug {
		fmt.Println("响应数据:", string(bodyByte))
	}

	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyByte, response)
	if err != nil {
		return err
	}

	return err
}

//代理请求函数,会对code进行判定
func (this *lwagentService)PostProxy(url string,requestData,responseData interface{}) error {
	response:=new(Response)
	if responseData!= nil {
		if !reflectutil.IsStructPtr(reflect.TypeOf(responseData)) {
			return errors.New("responseData 必须为指针")
		}
		response.Data=responseData
	}


	err:=this.Post(url,requestData,response)
	if err != nil {
		return err
	}

	if response.Code!= 0 {
		return errors.New(response.Msg)
	}

	return nil

}

//请求的数据格式
func (this *lwagentService) makeRequest(data interface{}) (Request, error) {

	var err error
	res := Request{
		Code: this.randomString(32),
		Time: strconv.Itoa(int(time.Now().Unix())),
	}

	aesKey := this.getAesKey(res.Code, res.Time)

	res.Data, err = this.getAesData(aesKey, data)
	if err != nil {
		return res, err
	}

	return res, nil

}

//生成随机字符串
func (this *lwagentService) randomString(num int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, num)
	rand.Seed(time.Now().Unix())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//获取aes加密key
func (this *lwagentService) getAesKey(code string, time string) string {

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

//获取加密后的data数据
func (this *lwagentService) getAesData(key string, data interface{}) (string, error) {

	jsonData, err := json.Marshal(data)

	if this.Debug {
		fmt.Println("请求数据:", string(jsonData))
	}
	if err != nil {
		return "", err
	}

	str, ok := stringutil.AESEncrypt([]byte(key), string(jsonData))
	if !ok {
		return "", errors.New("数据加密失败")
	}

	return str, nil

}
