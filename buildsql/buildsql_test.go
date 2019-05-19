package buildsql

import (
	"fmt"
	"testing"
)

//股票最新数据模型
type StockHsas struct {
	Id               int     `orm:"column(id);table(rms_hsas)" `
	Code             string  `orm:"column(code)"`               //股票代码
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

//插入
func TestBuildSql_Insert(t *testing.T) {
	sql, err := NewModel(StockHsas{Code:"131"}).Insert()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sql)
}

//更新
func TestBuildSql_Update(t *testing.T) {
	sql, err := NewModel(StockHsas{Code:"131"}).Where("id",1).Update()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sql)
}

//删除
func TestBuildSql_Delete(t *testing.T) {
	sql, err := NewModel(StockHsas{}).Where("id",1).Delete()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sql)
}


//查询
func TestBuildSql_Select(t *testing.T) {
	sql, err := NewModel(StockHsas{}).Where("id",1).Select()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sql)
}

func TestBuildSql_LeftJoin(t *testing.T) {

	sql, err := NewModel(StockHsas{},"T").
		LeftJoin(StockParse{},"B","A.code=B.code").
		Where("A.id",1).
		Select()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sql)

}

func TestBuildSql_Field(t *testing.T) {

	sql, err := NewModel(StockHsas{},"A").
		LeftJoin(StockParse{},"B","A.code=B.code").
		Where("A.id",1).
		Field("A.id","A.code","B.id").
		Select()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sql)

}


//func TestIsSatisfied(t *testing.T)  {
//
//	res:=IsSatisfied(float64(0))
//	fmt.Println(res)
//}

func TestGetColumnName(t *testing.T) {
	str:=GetColumnName("column(id);table(rms_hsas)","table")


	fmt.Println(str)
}

func TestBuildSql_GetTableNameFromModel(t *testing.T) {

	tableName,err:=GetTableNameFromModel(StockHsas{})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(tableName)

}