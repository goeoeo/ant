package autodoc

import (
	"testing"
)

func TestNewParseFile(t *testing.T) {
	p := NewParseFile()

	//beego框架控制器
	p.Config.ScanDir = "/home/yu/code/yungengxin2019/lwyapi/controllers"

	//gin 框架控制器
	//p.Config.ScanDir="/home/yu/gomod/cloud-storage/controller"
	//p.Config.MatchCallBack=GinControllerParseFun

	res, err := p.Do()
	if err != nil {
		t.Error(err)
		return
	}
	jsonFmt(res)
}

func TestGinControllerParseFun(t *testing.T) {
	jsonFmt(GinControllerParseFun("/home/yu/code/flea/api/controller/goods.go"))
}