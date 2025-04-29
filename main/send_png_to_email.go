package main

import (
	"crypto/tls"
	"fmt"
	"github.com/aobco/log"
	"gopkg.in/gomail.v2"
	"path/filepath"
	"xsky.com/ocpf/pkg/alert/service/mail"
	"xsky.com/ocpf/pkg/config"
	"xsky.com/ocpf/pkg/pg"
)

func main() {
	pg.Init(&pg.Config{
		Url:     "postgres://console_u:Xsky@123@127.0.0.1:5432/console?sslmode=disable",
		ShowSql: false,
	})
	if err := config.Configuration.Init(); err != nil {
		panic(err)
	}
	notifyService := mail.NotifyService
	if err := notifyService.Init(config.Configuration); err != nil {
		panic(err)
	}
	body := `
		<h1>产品展示</h1>
		<p>这是我们的最新产品：</p>
		<img src="cid:datareport" alt="产品图片" width="300">
		<p>欢迎咨询！</p>
	`
	m := gomail.NewMessage()
	m.SetHeader("From", "kang.caiqin@xsky.com")
	m.SetHeader("To", "lin.wanwei@xsky.com")
	m.SetHeader("Subject", "test")
	m.SetBody("text/html", body)

	imgPath, err := filepath.Abs("./output.png") // 转换为绝对路径
	if err != nil {
		log.Fatal("获取图片路径失败:", err)
	}

	m.Embed(imgPath, gomail.SetHeader(map[string][]string{
		"Content-ID": {fmt.Sprintf("<%s>", "datareport")}, // 设置CID
	}))

	dialer := gomail.NewDialer("smtp.feishu.cn", 465, "kang.caiqin@xsky.com", "r5xjCPzsTUl0RvUw")

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err = dialer.DialAndSend(m)
	if err != nil {
		log.Fatal("发送邮件失败:", err)
	}
}
