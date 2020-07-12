package autodoc

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type (
	ParseFunc func(file string) (outs []Comment)

	//通过正则解析出控制器注释和方法名称
	Comment struct {
		File   string //从哪个文件解析出来的
		Remark string //注释名称
		Method string //方法名称
	}

	//ParseFile
	ParseFileConfig struct {
		ScanDir       string    //扫描目录
		MatchCallBack ParseFunc //匹配解析函数
	}
	ParseFile struct {
		Config    ParseFileConfig
		scanFiles []string //需要扫描的文件
	}
)

var (
	//默认配置
	DefaultParseFileConfig = ParseFileConfig{
		ScanDir:       "./",
		MatchCallBack: BeegoControllerParseFun,
	}
)

func NewParseFile() *ParseFile {
	return &ParseFile{Config: DefaultParseFileConfig}
}

//执行扫描
func (this *ParseFile) Do() (outs []Comment, err error) {
	if err = this.loadFiles(); err != nil {
		return
	}

	for _, file := range this.scanFiles {
		outs = append(outs, this.Config.MatchCallBack(strings.TrimRight(this.Config.ScanDir, "/")+"/"+file)...)
	}

	return
}

//载入分析文件
func (this *ParseFile) loadFiles() (err error) {
	files, err := ioutil.ReadDir(this.Config.ScanDir)
	if err != nil {
		return err
	}

	for _, f := range files {
		if strings.Contains(f.Name(), "_test.go") || !strings.Contains(f.Name(), ".go") {
			continue
		}

		this.scanFiles = append(this.scanFiles, f.Name())
	}

	return
}

//beego框架，控制器文件解析器
/*
内容：

//删除图标
func (this *DesktopIconController) Del() {

解析结果：
{
	Remark:"删除图标",
	Method:"Del"
}
*/
func BeegoControllerParseFun(file string) (outs []Comment) {

	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	re := regexp.MustCompile(fmt.Sprintf(`//([^\n]+)\nfunc \(this \*([A-Z][A-Z0-9a-z]*Controller)\) ([A-Z][A-Z0-9a-z]*)\(\) \{`))
	res := re.FindAllStringSubmatch(string(content), -1)

	for _, v := range res {

		if len(v) != 4 {
			continue
		}

		outs = append(outs, Comment{
			File:   file,
			Remark: v[1],
			Method: v[3],
		})

	}

	return
}

//gin框架，控制器文件解析器
/*
内容：

//获取分块上传整体状态
func MultipartUploadStatus(this *gin.Context) {

解析结果：
{
	Remark:"获取分块上传整体状态",
	Method:"MultipartUploadStatus"
}
*/
func GinControllerParseFun(file string) (outs []Comment) {

	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	re := regexp.MustCompile(fmt.Sprintf(`//([^\n]+)\nfunc \(([A-Z][A-Z0-9a-z]*Controller)\) ([A-Z][A-Z0-9a-z]*)\(ctx \*gin.Context\) \{`))
	res := re.FindAllStringSubmatch(string(content), -1)

	for _, v := range res {

		if len(v) != 4 {
			continue
		}

		outs = append(outs, Comment{
			File:   file,
			Remark: v[1],
			Method: fmt.Sprintf("%s_%s",v[2],v[3]),
		})

	}

	return
}
