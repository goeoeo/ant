package mail

import "gopkg.in/gomail.v2"

func SendMail()  {
	m := gomail.NewMessage()
	m.SetHeader("From", "rentmaterial@163.com")                     //发件人
	m.SetHeader("To", "977564830@qq.com")           //收件人
	m.SetHeader("Subject", "Hello!")                     //邮件标题
	m.SetBody("text/html", "使用Go测试发送邮件!")     //邮件内容

	d := gomail.NewDialer("smtp.163.com", 465, "rentmaterial@163.com", "UWDHGKJKRHVPYWSA")
	//邮件发送服务器信息,使用授权码而非密码
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}