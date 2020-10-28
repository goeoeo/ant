package magicutil

import (
	"fmt"
	"github.com/bregydoc/gtranslate"
	"github.com/phpdi/ant/util"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
)

type TranFileFromCode struct {
}


//执行翻译
func (this *TranFileFromCode) Run(targetFile string, filePaths ...string) {
	areadyTrwords := this.getTargetFileWords(targetFile)
	appendContent := this.tran(areadyTrwords, filePaths...)
	oldContent, _ := ioutil.ReadFile(targetFile)

	newContent := string(oldContent) + appendContent

	if err := ioutil.WriteFile(targetFile, []byte(newContent), os.ModePerm); err != nil {
		fmt.Println(err)
	}

	fmt.Println("完成.")
}

//执行翻译，依赖google-api，先科学上网
func (this *TranFileFromCode) tran(areadyTrwords []string, filePaths ...string) string {
	var (
		wordsMap    map[string]string
		words       []string
		content     string
		uniqueWords []string
	)

	wordsMap = make(map[string]string)
	files := this.getGofile(filePaths...)

	for _, v := range files {
		words = append(words, util.ParseChnFromGolang(v)...)
	}

	//gtranslate.GoogleHost="google.cn"

	for _, v := range words {
		v = strings.TrimLeft(v, `"`)

		wordsMap[v] = ""
	}

	for k := range wordsMap {
		uniqueWords = append(uniqueWords, k)
	}

	if uniqueWords != nil {
		sort.Strings(uniqueWords)
	}

	fmt.Println("扫描文件总数:", len(files))
	fmt.Println("词汇数量：", len(uniqueWords))
	tranCompleteNum := 0
	go func() {
		fmt.Println("正在翻译...")
		for {
			if tranCompleteNum == len(uniqueWords) {
				break
			}

			process := 0
			if len(uniqueWords) > 0 {
				process = tranCompleteNum * 100 / len(uniqueWords)
			}

			fmt.Printf("已完成:%d,进度:%d%%\n", tranCompleteNum, process)

			time.Sleep(3 * time.Second)
		}

	}()
	for _, k := range uniqueWords {
		tranCompleteNum++
		//不需要翻译
		if this.in(k, areadyTrwords) {
			continue
		}
		translated := ""
		var err error
		translated, err = gtranslate.TranslateWithParams(
			k,
			gtranslate.TranslationParams{
				From: "zh-CN",
				To:   "en",
			},
		)

		if err != nil {
			fmt.Println(err)
		}

		translated = strings.ToLower(strings.Replace(translated, " ", "_", -1))

		content += fmt.Sprintf("%s = %s\n", translated, k)

	}

	return content
}

//递归扫描目录获取文件列表
func (this *TranFileFromCode) scanpath(dir string) (res []string) {

	if util.IsFile(dir) {
		return []string{dir}
	}

	if r, err := util.ScanPath(dir, 2); err == nil {
		for k := range r {
			res = append(res, k)
		}
	}

	return
}

//过滤go文件
func (this *TranFileFromCode) getGofile(filePaths ...string) (res []string) {
	var files []string

	for _, v := range filePaths {
		files = append(files, this.scanpath(v)...)
	}

	sort.Strings(files)

	for _, v := range files {

		if !strings.HasSuffix(v, ".go") || strings.HasSuffix(v, "_test.go") {
			continue
		}

		res = append(res, v)
	}

	return
}

//从目标文件中获取已经翻译好的词句
func (this *TranFileFromCode) getTargetFileWords(filePath string) (words []string) {
	content, _ := ioutil.ReadFile(filePath)
	if arr := strings.Split(string(content), "\n"); len(arr) > 0 {
		for _, v := range arr {
			if arr1 := strings.Split(v, "="); len(arr1) == 2 {
				words = append(words, strings.Trim(arr1[1], " "))
			}
		}
	}
	return
}

func (this *TranFileFromCode) in(s string, ss []string) bool {
	for _, v := range ss {
		if s == v {
			return true
		}
	}

	return false
}
