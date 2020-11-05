/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/5 0:20
 */
package email

import (
	"gopkg.in/gomail.v2"
)

type Email struct {
	o *Options
}

func NewEmail(options *Options) *Email {
	return &Email{options}
}

func (e *Email) SendEmail(mailTo []string, subject string, body string) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(e.o.Username, "reminder")) //这种方式可以添加别名，即“XX官方”
	m.SetHeader("To", mailTo...)                                   //发送给多个用户
	m.SetHeader("Cc", m.FormatAddress(e.o.Username, "reminder"))   //抄送
	m.SetHeader("Subject", subject)                                //设置邮件主题
	m.SetBody("text/html", body)                                   //设置邮件正文
	d := gomail.NewDialer(
		e.o.Host,
		e.o.Port,
		e.o.Username,
		e.o.Password)
	err = d.DialAndSend(m)
	return err
}
