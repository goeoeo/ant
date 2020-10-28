package md

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMdToHtmlWithCss(t *testing.T) {
	file, _ := ioutil.ReadFile("./markdown.md")

	out := MdToHtmlWithCssNoDir(file)

	ioutil.WriteFile("markdown.html", []byte(out), os.ModePerm)
}
