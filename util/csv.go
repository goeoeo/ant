package util

import (
	"bytes"
	"encoding/csv"
)

type CsvUtil struct {
	buff *bytes.Buffer
	Wr   *csv.Writer
}

func NewCsv(title ...string) *CsvUtil {
	this := &CsvUtil{buff: new(bytes.Buffer)}
	this.Wr = csv.NewWriter(this.buff)

	//防止word打开出现乱码
	this.buff.WriteString("\xEF\xBB\xBF")

	//写入了表头
	if len(title) > 0 {
		this.Row(title...)
	}

	return this
}

//内容行
func (this *CsvUtil) Row(row ...string) error {
	return this.Wr.Write(row)
}

//输出数据
func (this *CsvUtil) Bytes() []byte {

	this.Wr.Flush()
	return this.buff.Bytes()
}
