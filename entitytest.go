package ant

//测试需要的结构体

//股票最新数据模型
type StockHsas struct {
	Id               int     `orm:"column(id);table(rms_hsas)" valid:"Name(Id);Max(20);Min(5)" `
	Code             string  `orm:"column(code)" valid:"Name(股票代码)"`               //股票代码
	Name             string  `orm:"column(name)"`               //股票名称
	Date             string  `orm:"column(date)"`               //日期
	OpenToday        float64 `orm:"column(open_today)"`         //今开
	Highest          float64 `orm:"column(highest)"`            //最高
	Lowest           float64 `orm:"column(lowest)"`             //最低
	Yesterday        float64 `orm:"column(yesterday)"`          //收盘价
	RiseAndFall      float64 `orm:"column(rise_and_fall)"`      //涨跌额
	QuoteChange      float64 `orm:"column(quote_change)"`       //涨跌幅
	Volume           int     `orm:"column(volume)"`             //成交量
	Turnover         int     `orm:"column(turnover)"`           //成交额
	Amplitude        float64 `orm:"column(amplitude)"`          //振幅
	HandTurnoverRate float64 `orm:"column(hand_turnover_rate)"` //换手率
	Diff             float64 `orm:"column(diff)"`               //diff
	Dea              float64 `orm:"column(dea)"`                //dea
}

type StockParse struct {
	Id               int     `orm:"column(id);table(rms_parse)" `
	Code             string  `orm:"column(code)"`               //股票代码
	Name             string  `orm:"column(name)"`               //股票名称
}