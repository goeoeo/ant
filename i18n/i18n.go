package i18n

import (
	"fmt"
	"github.com/phpdi/ant/util"
	"io/ioutil"
	"log"
	"strings"
)

const Zh_CN = "zh-CN"

type I18n struct {
	chnMap   map[string] /*value*/ string                        /*key*/ //中文=>key映射
	langData map[string] /*langType*/ map[string] /*key*/ string /*value*/
}

func NewI18n() *I18n {
	i := &I18n{
		chnMap:   make(map[string]string),
		langData: make(map[string]map[string]string),
	}

	return i
}

func NewI18nWithDir(dir string) *I18n {
	i := NewI18n()

	//初始化语言包
	i.initLang(dir)
	return i
}

//初始化语言包
func (i *I18n) initLang(fileDir string) {

	files, err := util.ScanPath(fileDir, 2)
	if err != nil {
		log.Println(err)
		return
	}

	//解析语言包类型 locale_zh-CN.lang =>zh-CN
	parseLangType := func(fileName string) string {

		if !strings.HasSuffix(fileName, ".lang") {
			return ""
		}

		if arr := strings.Split(fileName, "/"); len(arr) == 0 {
			return ""
		} else {
			suffix := arr[len(arr)-1]
			suffix = strings.Replace(suffix, "locale_", "", 1)
			return strings.Replace(suffix, ".lang", "", 1)
		}

	}

	for file := range files {
		if langType := parseLangType(file); langType != "" {
			if langType == Zh_CN {
				log.Println(i.initChnMap(file))
			}
			i.SetMessage(langType, file)
		}
	}
}

//设置消息
func (i *I18n) SetMessage(langType string, file string) {
	words := i.getWordsFromFile(file)
	if _,ok:=i.langData[langType];!ok {
		i.langData[langType]=make(map[string]string)
	}
	for _, v := range words {

		i.langData[langType][v[0]] = v[1]
	}
}

//直接翻译中文
func (i *I18n) TrChn(langType, format string, args ...interface{}) string {
	//通过中文找key
	if tmp, ok := i.chnMap[format]; ok {
		//执行翻译
		if trRes := i.Tr(langType, tmp, args...); trRes != tmp {
			//翻译成功
			return trRes
		}
	}

	return format

}

//执行翻译
func (i *I18n) Tr(langType, format string, args ...interface{}) string {
	if tmp, ok := i.langData[langType]; ok {
		if tmp1, ok := tmp[format]; ok {
			return fmt.Sprintf(tmp1, args...)
		}
	}
	return format
}

//解析中文映射
func (i *I18n) initChnMap(file string) (err error) {
	var (
		znWords [][2]string
	)

	if znWords, err = i.parseFile(file); err != nil {
		return
	}

	for _, v := range znWords {
		i.chnMap[v[1]] = v[0]
	}

	return
}

//解析文件
func (i *I18n) parseFile(file string) (words [][2]string, err error) {
	var content []byte
	if content, err = ioutil.ReadFile(file); err != nil {
		return
	}
	if arr := strings.Split(string(content), "\n"); len(arr) > 0 {
		for _, v := range arr {
			if arr1 := strings.Split(v, "="); len(arr1) == 2 {
				words = append(words, [2]string{strings.Trim(arr1[0], " "), strings.Trim(arr1[1], " ")})
			}
		}
	}

	return
}

//从文件中解析翻译数据
func (this *I18n) getWordsFromFile(filePath string) (words [][2]string) {
	content, _ := ioutil.ReadFile(filePath)
	if arr := strings.Split(string(content), "\n"); len(arr) > 0 {
		for _, v := range arr {
			if arr1 := strings.Split(v, "="); len(arr1) == 2 {

				words = append(words, [2]string{strings.Trim(arr1[0], " "), strings.Trim(arr1[1], " ")})
			}
		}
	}
	return
}
