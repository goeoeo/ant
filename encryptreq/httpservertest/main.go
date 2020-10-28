package main

import (
	"encoding/json"
	"fmt"
	"github.com/phpdi/ant/encryptreq"
	"io/ioutil"
	"net/http"
)

type StockParse struct {
	Id   int    `orm:"column(id);table(rms_parse)" `
	Code string `orm:"column(code)"` //股票代码
	Name string `orm:"column(name)"` //股票名称
}

func main() {
	http.HandleFunc("/", handle)
	//ListenAndServe监听srv.Addr指定的TCP地址，并且会调用Serve方法接收到的连接。如果srv.Addr为空字符串，会使用":http"。
	err := http.ListenAndServe("0.0.0.0:8080", nil)

	if err != nil {
		fmt.Println("http listen failed")
	}
}

func handle(writer http.ResponseWriter, r *http.Request) {
	var (
		request encryptreq.Request
		resp    encryptreq.Response
		req     StockParse

		outBody []byte
	)
	body, _ := ioutil.ReadAll(r.Body)

	entcryptReq := encryptreq.EncryptReq{}
	//    r.Body.Close()
	fmt.Println("请求数据", string(body))

	json.Unmarshal(body, &request)

	entcryptReq.DecryptRequest(&req, request)

	fmt.Println(req)

	resp.Data = req

	outBody, _ = json.Marshal(resp)

	writer.Write(outBody)

}
