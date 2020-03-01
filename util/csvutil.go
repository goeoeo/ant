package util

import (
	"bytes"
	"encoding/csv"
)

type CsvUtil struct {
	fields []string
	rows   [][]string
	title  []string
}

func NewCsvUtil(title []string, rows [][]string) *CsvUtil {
	this := new(CsvUtil)
	this.rows = rows
	this.title = title

	return this
}

//csv导出工具
func (this *CsvUtil) Export() []byte {
	b := new(bytes.Buffer)
	wr := csv.NewWriter(b)

	// 写入UTF-8 BOM，防止excel打开时中文乱码
	b.WriteString("\xEF\xBB\xBF")
	//写入标题
	wr.Write(this.title)

	//按行写入数据
	for i := 0; i < len(this.rows); i++ {
		wr.Write(this.rows[i])
	}

	wr.Flush()

	return b.Bytes()

}
