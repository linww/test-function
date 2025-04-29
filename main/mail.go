package main

import (
	"github.com/aobco/log"
	"net/smtp"
	"strings"
	"xsky.com/ocpf/pkg/alert/model"
	"xsky.com/ocpf/pkg/alert/service/mail"
	"xsky.com/ocpf/pkg/config"
	"xsky.com/ocpf/pkg/pg"
)

func SendToMail(user, password, host, subject, date, body, mailtype, replyToAddress string, to, cc, bcc []string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	cc_address := strings.Join(cc, ";")
	bcc_address := strings.Join(bcc, ";")
	to_address := strings.Join(to, ";")
	msg := []byte("To: " + to_address + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\nDate: " + date + "\r\nReply-To: " + replyToAddress + "\r\nCc: " + cc_address + "\r\nBcc: " + bcc_address + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := MergeSlice(to, cc)
	send_to = MergeSlice(send_to, bcc)
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func MergeSlice(s1 []string, s2 []string) []string {
	slice := make([]string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}

//func main() {
//	dialer := gomail.NewDialer("smtp.feishu.cn", 465, "kang.caiqin@xsky.com", "r5xjCPzsTUl0RvUw")
//	dial, err := dialer.Dial()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer dial.Close()
//	if err = dial.Send("kang.caiqin@xsky.com", []string{"itman_wei@163.com"}, strings.NewReader("hello")); err != nil {
//		log.Fatal(err)
//	}
//}

//func main() {
//	var email string
//	flag.StringVar(&email, "email", "", "要验证的邮箱")
//	flag.Parse()
//	if email == "" {
//		email = "2856dfsdfdsfsdf7582aaaaa2@qq.com"
//	}
//	configuration, err := truemail.NewConfiguration(truemail.ConfigurationAttr{
//		VerifierEmail: "itman_wei@163.com", // 替换为您的验证邮箱
//		//SmtpPort:       465,
//		//VerifierDomain:    "smtp.feishu.cn", // 替换为您的验证邮箱的域名
//		ConnectionTimeout: 2,
//		ResponseTimeout:   2,
//	})
//	if err != nil {
//		fmt.Println("初始化配置失败:", err)
//		return
//	}
//	r, err := truemail.Validate(email, configuration)
//	if err != nil {
//		fmt.Println("验证失败:", err)
//		return
//	}
//	log.Infof("邮箱验证结果: %+v", r)
//}

//func main() {
//	// {"host":"smtp.feishu.cn","port":465,"user":"kang.caiqin@xsky.com","password":"r5xjCPzsTUl0RvUw","ssl_enabled":false,"skip_tls_cert_verify":true}
//	err := util2.ValidateHostAndUser("smtp.feishu.cn", "itman_wei@163.com", "2856758aaaaa22@qq.com")
//	if err != nil {
//		log.Fatal(err)
//	}
//}

//var (
//	verifier = emailverifier.NewVerifier()
//)

//func main() {
//	//email := "itman_wei111@163.com"
//
//	if err := verifier.EnableAPIVerifier("yahoo"); err != nil {
//		fmt.Println("enable yahoo verifier failed, error is: ", err)
//		return
//	}
//	verifier.EnableSMTPCheck()
//	verifier.FromEmail("kang.caiqin@xsky.com")
//	verifier.DisableCatchAllCheck()
//	ret, err := verifier.CheckSMTP("163.com", "aaaaaa")
//	if err != nil {
//		fmt.Println("verify email address failed, error is: ", err)
//		return
//	}
//
//	log.Infof("email validation result: %+v", ret)
//	/*
//		result is:
//		{
//			"email":"example@exampledomain.org",
//			"disposable":false,
//			"reachable":"unknown",
//			"role_account":false,
//			"free":false,
//			"syntax":{
//			"username":"example",
//				"domain":"exampledomain.org",
//				"valid":true
//			},
//			"has_mx_records":true,
//			"smtp":null,
//			"gravatar":null
//		}
//	*/
//}

func main() {
	pg.Init(&pg.Config{
		Url:     "postgres://console_u:xsky@2021@127.0.0.1:5432/datasource?sslmode=disable",
		ShowSql: false,
		MaxIdle: 10,
		MaxOpen: 10,
	})
	err := config.Configuration.Init()
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	if err = mail.NotifyService.Init(config.Configuration); err != nil {
		log.Fatalf("%v", err)
		return
	}
	message := mail.NotifyService.CreateMessage("1111", "22222", true, &model.AlertContact{
		Email:    "itman_wei123@111163.com",
		Dingding: nil,
		Wechat:   nil,
	})
	if message == nil {
		log.Fatalf("create message failed")
		return
	}
	if err = mail.NotifyService.Send(message); err != nil {
		log.Fatalf("%v", err)
		return
	}
	log.Infof("send email success")
}
