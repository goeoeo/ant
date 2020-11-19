package test

import "time"

//测试需要的结构体

//股票最新数据模型
type StockHsas struct {
	Id               int     `field:"ID" orm:"column(id);" valid:"Name(Id);Max(20);Min(5)" `
	Code             string  `orm:"-" valid:"Name(股票代码)"`       //股票代码
	Name             string  `orm:"column(name)"`               //股票名称
	Date             string  `orm:"column(date)"`               //日期
	OpenToday        float64 ``                                 //今开
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
	Sp               StockParse
}

type StockParse struct {
	Id   int    `field:"分析id"orm:"column(id);table(rms_parse)" valid:"Max(20);Min(5)"`
	Code string `orm:"column(code)"` //股票代码
	Name string `orm:"column(name)"` //股票名称
}

type User struct {
	Id         int       `orm:"size(11);column(id);"`
	UserName   string    `field:"账号" orm:"size(50);column(username)"`                                          //用户名：代理商、代理商员工（手机号登录），管理员、运营商、网吧（账号登录）
	Password   string    `orm:"size(50);column(password)"`                                                     //密码
	Salt       string    `orm:"size(32);column(salt)"`                                                         //密码加密盐
	Name       string    `field:"姓名" orm:"size(50);column(name)" valid:"Chn;MinSize(2);MaxSize(12);Sensitive"` //姓名
	UserType   int       `orm:"size(4);column(user_type)"`                                                     //用户类型：1=超级管理员，2=平台管理员，3=运营商，4=代理商，5=代理商员工，6=网吧
	Mobile     string    `field:"手机号" orm:"size(15);column(mobile)" valid:"Mobile"`                            //手机号码
	Email      string    `field:"邮箱" orm:"size(50);column(email) " valid:"Email"`                              //邮箱
	Memo       string    `field:"备注" orm:"size(255);column(memo)" valid:"MaxSize(20);Sensitive"`               //备注
	Status     int       `orm:"size(4);column(status)"`                                                        //状态：1=可用，2=禁用
	LastIp     string    `orm:"size(15);column(last_ip)"`                                                      //最后登录ip
	CreateTime time.Time `orm:"auto_now_add;type(datetime);column(create_time)"`                               //创建时间
	UpdateTime time.Time `orm:"auto_now_add;type(datetime);column(update_time)"`                               //更新时间

}
