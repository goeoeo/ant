package ocr

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/url"
	"sort"
	"strings"
	"time"
)

type (
	TencentOcr struct {
		appId int
		appKey string

	}
	baseAck struct {
		Ret int
		Msg string
		Data interface{}
	}
)

func NewTencentOcr(appId int,appKey string) *TencentOcr {
	this:=&TencentOcr{
		appId:  appId,
		appKey: appKey,
	}
	return this
}

//通用OCR
func (this *TencentOcr)Universal(filePath string)(res string,err error)  {
	var (
		imageContent []byte
		ack baseAck
	)

	if imageContent,err=ioutil.ReadFile(filePath);err!= nil {
		return
	}

	encodeToString := base64.StdEncoding.EncodeToString(imageContent)

	params:=this.baseSign()
	params["image"]= url.QueryEscape(encodeToString)
	params["sign"]=this.createSign(params)

	if err=Post("https://api.ai.qq.com/fcgi-bin/ocr/ocr_generalocr",params,&ack);err!= nil {
		return
	}

	JsonPrint(ack)

	return
}

func (this *TencentOcr)baseSign() map[string]interface{} {
	params:=make(map[string]interface{})
	//基础参数
	params["app_id"]=this.appId
	params["time_stamp"]=time.Now().Unix()
	params["nonce_str"]=RandomString(16)
	return params
}

func (this *TencentOcr)createSign(params map[string]interface{}) string {
	var (
		keys []string
		appendString string
	)

	for k:=range params {
		keys=append(keys,k)
	}

	//key 排序
	sort.Strings(keys)
	//拼接
	for _,k:=range keys {
		//参数值为空不参与签名
		if params[k]== nil {
			continue
		}
		appendString+=fmt.Sprintf("%s=%v&",k,params[k])
	}

	//拼接秘钥
	appendString+=fmt.Sprintf("app_key=%v&",this.appKey)

	// MD5
	hash := md5.New()
	hash.Write([]byte(appendString))
	return strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))
}


