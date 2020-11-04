package magicutil

import "fmt"

//进度条
type ProgressBar struct {
	percent int64  //百分比
	cur     int64  //当前进度位置
	total   int64  //总进度
	rate    string //进度条
	graph   string //显示符号
}

func (this *ProgressBar) NewOption(start, total int64) {
	this.cur = start
	this.total = total
	if this.graph == "" {
		this.graph = "█"
	}
	this.percent = this.getPercent()

	for i := 0; i < int(this.percent); i += 2 {
		this.rate += this.graph
	}
}

func (this *ProgressBar) NewOptionWithGraph(start, total int64, graph string) {
	this.graph = graph
	this.NewOption(start, total)
}

func (this *ProgressBar) getPercent() int64 {
	return int64(float32(this.cur) / float32(this.total) * 100)
}

func (this *ProgressBar) Play(cur int64) {
	this.cur = cur
	last := this.percent
	this.percent = this.getPercent()
	if this.percent != last && this.percent%2 == 0 {
		this.rate += this.graph
	}

	fmt.Printf("\r[%-50s]%3d%%  %8d/%d", this.rate, this.percent, this.cur, this.total)

	if this.cur == this.total {
		fmt.Println()
	}
}
