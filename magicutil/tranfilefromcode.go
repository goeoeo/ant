package magicutil

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/bregydoc/gtranslate"
	"github.com/phpdi/ant/util"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

//扫描文件中文，翻译成对应英文
type TranFileFromCode struct {
	skipWords      [][2]string
	uniqueWords    []string //扫描出的词汇组合
	outUniqueWords [][2]string

	filePaths []string //扫描目录

	targetFile string //目标文件
}

func NewTranFileFromCode(scanPaths ...string) *TranFileFromCode {
	this := &TranFileFromCode{
		filePaths:  scanPaths,
		targetFile: "locale.lang",
	}
	return this
}

//设置目标文件
func (this *TranFileFromCode) SetTargetFile(file string) *TranFileFromCode {
	this.targetFile = file

	return this
}

//执行翻译,追加方式翻译
func (this *TranFileFromCode) AppendRun() *TranFileFromCode {
	var (
		oldword [][2]string
		newWord [][2]string
	)
	//文件中数据
	oldword = this.getTargetFileWords(this.targetFile)

	//扫描出的数据
	sourceWords := this.GetKvs()

	in := func(word [2]string) bool {
		for _, v := range oldword {
			if v[1] == word[1] {
				return true
			}
		}

		return false
	}

	for _, v := range sourceWords {
		if !in(v) {
			newWord = append(newWord, v)
		}
	}

	this.writeFile(append(oldword, newWord...), this.targetFile)

	return this
}

//执行翻译
func (this *TranFileFromCode) Run() *TranFileFromCode {
	sourceWords := this.GetKvs()
	this.writeFile(sourceWords, this.targetFile)
	fmt.Println("翻译完成.")

	return this
}

//扫描文件
func (this *TranFileFromCode) scanWords() {
	var (
		wordsMap map[string]string
		words    []string
	)

	wordsMap = make(map[string]string)
	files := this.getGofile(this.filePaths...)

	for _, v := range files {
		words = append(words, util.ParseChnFromGolang(v)...)
	}

	//gtranslate.GoogleHost="google.cn"

	for _, v := range words {
		v = strings.Trim(v, `"`)

		wordsMap[v] = ""
	}

	for k := range wordsMap {
		this.uniqueWords = append(this.uniqueWords, k)
	}
	sort.Strings(this.uniqueWords)
}

//预览
func (this *TranFileFromCode) Preview() *TranFileFromCode {
	this.scanWords()

	fmt.Println("扫描总量:", len(this.uniqueWords))
	for _, v := range this.uniqueWords {
		fmt.Println(v)
	}
	return this
}

//扫描目标文件
func (this *TranFileFromCode) getWordsFromTargetFile(word string) [2]string {
	if this.skipWords == nil {
		this.skipWords = this.getTargetFileWords(this.targetFile)
	}

	for _, v := range this.skipWords {
		if word == v[1] {
			return v
		}
	}

	return [2]string{}

}

//获取英文=>中文组合
func (this *TranFileFromCode) GetKvs() [][2]string {

	//扫描目标文件
	this.scanWords()

	getWord := func(word string) [2]string {
		if w := this.getWordsFromTargetFile(word); w[0] != "" {
			return w
		}

		//调用google进行翻译
		translated, err := gtranslate.TranslateWithParams(
			word,
			gtranslate.TranslationParams{
				From: "zh-CN",
				To:   "en",
			},
		)

		if err != nil {
			fmt.Println(err)
			return [2]string{}
		}

		return [2]string{translated, word}
	}

	//进度条
	var bar ProgressBar
	bar.NewOption(0, int64(len(this.uniqueWords)))
	fmt.Println("正在翻译...")
	for k, word := range this.uniqueWords {
		bar.Play(int64(k + 1))
		if words := getWord(word); words[0] != "" {
			this.outUniqueWords = append(this.outUniqueWords, words)
		}

	}

	return this.outUniqueWords
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
func (this *TranFileFromCode) getTargetFileWords(filePath string) (words [][2]string) {
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

//生成中英文文件
func (this *TranFileFromCode) MakeFile() {
	words := this.getTargetFileWords(this.targetFile)
	var (
		enWords [][2]string
		zhWords [][2]string
	)

	for _, v := range words {
		key := Md5(v[1])
		enWords = append(enWords, [2]string{key, strings.ToLower(v[0])})
		zhWords = append(zhWords, [2]string{key, v[1]})
	}

	this.writeFile(enWords, "locale_us-EN.lang")
	this.writeFile(zhWords, "locale_zh-CN.lang")

}

func (this *TranFileFromCode) writeFile(source [][2]string, fileName string) {
	var content string
	for _, v := range source {
		content += fmt.Sprintf("%s = %s\n", v[0], v[1])
	}

	ioutil.WriteFile(fileName, []byte(content), os.ModePerm)
}

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
