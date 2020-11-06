package i18n

import (
	"github.com/phpdi/ant/util"
	"io/ioutil"
	"log"
	"strings"
)

const zh_CN = "zh-CN"

type I18n struct {
	fileDir string            //语言包所在的文件夹
	chnMap  map[string]string //中文=>key映射
}

func NewI18n(dir string) *I18n {
	i := &I18n{
		fileDir: dir,
		chnMap:  make(map[string]string),
	}

	//初始化语言包
	i.initLang()
	return i
}

//初始化语言包
func (i *I18n) initLang() {

	files, err := util.ScanPath(i.fileDir, 2)
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
			if langType == zh_CN {
				log.Println(i.initChnMap(file))
			}
			log.Println(i18n.SetMessage(langType, file))
		}

	}
}

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
