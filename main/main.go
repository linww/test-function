package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/aobco/log"
	uuid2 "github.com/google/uuid"
	"github.com/vjeantet/grok"
	"net/smtp"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
	"xsky.com/ocpf/console/service/es"
	"xsky.com/ocpf/pkg/usercenter/model"
	"xsky.com/ocpf/pkg/util"
)

type status string

func main() {
	now := time.Now()
	nowTrim := now.Truncate(time.Hour)
	fmt.Println("now:", now)
	fmt.Println("nowTrim:", nowTrim)
}

func IsSubPath(dir, file string) bool {
	filePath := filepath.Dir(file)
	filePathAddSuffix := filePath + string(filepath.Separator)
	dirAddSuffix := dir + string(filepath.Separator)
	return filePathAddSuffix == dirAddSuffix || strings.HasPrefix(filePathAddSuffix, dirAddSuffix)
}

func sendEmail(to string) error {
	// {"host":"smtp.feishu.cn","port":465,"user":"kang.caiqin@xsky.com","password":"r5xjCPzsTUl0RvUw","ssl_enabled":false,"skip_tls_cert_verify":true}
	from := "kang.caiqin@xsky.com"
	password := "r5xjCPzsTUl0RvUw"
	smtpHost := "smtp.feishu.cn"
	smtpPort := "465"

	// Set up authentication information.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Set up the email message.
	msg := []byte("To: " + to + "\r\n" +
		"Subject: Test Email\r\n" +
		"\r\n" +
		"This is a test email.\r\n")

	// Connect to the SMTP server with TLS.
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsconfig)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %v", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %v", err)
	}
	defer client.Close()

	// Authenticate.
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("failed to authenticate: %v", err)
	}

	// Set the sender and recipient.
	if err = client.Mail(from); err != nil {
		return fmt.Errorf("failed to set sender: %v", err)
	}
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("failed to set recipient: %v", err)
	}

	// Send the email body.
	writer, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to send data: %v", err)
	}
	_, err = writer.Write(msg)
	if err != nil {
		return fmt.Errorf("failed to write message: %v", err)
	}

	return nil
}

func verifyEmail(email string) {
	// 验证用户（需要有效的 SMTP 服务器信息）
	serverHostName := "smtp.feishu.cn"          // 设置您的 SMTP 服务器
	serverMailAddress := "kang.caiqin@xsky.com" // 设置您的有效邮件地址

	err := util.ValidateHostAndUser(serverHostName, serverMailAddress, email)
	if err != nil {
		log.Fatalf("Failed to verify email: %v", err)
	} else {
		log.Infof("Email verification succeeded")
	}
}

func openLocalDir(path string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("open", path)
	case "linux":
		cmd = exec.Command("xdg-open", path)
	case "windows":
		cmd = exec.Command("explorer", path)
	default:
		log.Fatalf("unsupported platform")
	}

	err := cmd.Start()
	if err != nil {
		log.Fatalf("Failed to open directory: %v", err)
	}
}

func printShareLink() {
	//	template := `╭%s╮
	//│%s│
	//│%s│
	//╰%s╯`
	//	shareUrl := "   " + "http://10.16.64.161/share/182ea6ebcb5e4904896d872533db34f8" + "   "
	//	shareUrlLen := len(shareUrl)
	//	secondLine := "share url"
	//	firstLine := strings.Repeat("─", shareUrlLen)
	//	forthLine := strings.Repeat("─", shareUrlLen)
	//	secondLinePadLeft := (shareUrlLen - len(secondLine)) / 2
	//	secondLinePadRight := secondLinePadLeft
	//	if (shareUrlLen-len(secondLine))%2 != 0 {
	//		secondLinePadRight++
	//	}
	//	secondLine = strings.Repeat(" ", secondLinePadLeft) + secondLine + strings.Repeat(" ", secondLinePadRight)
	//	thirdLine := shareUrl
	//	result := fmt.Sprintf(template, firstLine, secondLine, thirdLine, forthLine)
	//	fmt.Println(result)
}

func wgTest(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	for i := 0; i < 10; i++ {
		log.Infof("i: %d", i)
	}
	wg.Add(1)
	go wgTest2(wg)
}

func wgTest2(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	for i := 0; i < 10; i++ {
		log.Infof("j: %d", i)
	}
	time.Sleep(10 * time.Second)
}

func uuid() string {
	return uuid2.Must(uuid2.NewRandom()).String()
}

func LdapListUsers() {
	server, err := model.NewLdapServer(&model.LdapConfigurationParam{
		Id:               100002,
		Name:             "test03",
		Address:          "10.255.101.60:389",
		BaseDn:           "dc=xsky,dc=ldaptest",
		BindUserDn:       "cn=Manager,dc=xsky,dc=ldaptest",
		BindUserPassword: "2943d88290cee8bbb7606cdda576c5a8",
		PropertyMapping: &model.LdapPropertyMapping{
			LoginName:   "uid",
			ObjectClass: []string{"person", "organizationalPerson"},
		},
	})
	if err != nil {
		log.Fatalf("new ldap server failed, %v", err)
	}
	users, err := server.ListUsers()
	if err != nil {
		log.Fatalf("list users failed, %v", err)
	}
	for _, user := range users {
		fmt.Println(user.EntryDn)
	}
}

func aggregateApi() {
	j := `[
        {
            "time": 1721311200000,
            "time_display": "2024-07-18T14:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 13484,
            "total_size": 7745526217,
            "agents": [
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 120,
                    "total_size": 1387739
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 1006,
                    "total_size": 1087009138
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 124,
                    "total_size": 162673872
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 1832,
                    "total_size": 939427305
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1523,
                    "total_size": 114446973
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 801,
                    "total_size": 444953821
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 32,
                    "total_size": 17270932
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 18,
                    "total_size": 12361424
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 264,
                    "total_size": 109001737
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 271,
                    "total_size": 10810589
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 38,
                    "total_size": 4998877
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 180,
                    "total_size": 1323536975
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 234,
                    "total_size": 172608705
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 334,
                    "total_size": 118370779
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 13,
                    "total_size": 70849
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 2464,
                    "total_size": 824432775
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 356,
                    "total_size": 615501403
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 91,
                    "total_size": 283476165
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 76,
                    "total_size": 39750170
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 13,
                    "total_size": 1424728
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 2520,
                    "total_size": 859992849
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 151,
                    "total_size": 1386347
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 134,
                    "total_size": 181249008
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 294,
                    "total_size": 105469662
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 467,
                    "total_size": 31171530
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 128,
                    "total_size": 282741865
                }
            ]
        },
        {
            "time": 1721304000000,
            "time_display": "2024-07-18T12:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 10090,
            "total_size": 5646719349,
            "agents": [
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 848,
                    "total_size": 1021314253
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 799,
                    "total_size": 447212788
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 130,
                    "total_size": 95983039
                },
                {
                    "agent_id": 100012,
                    "agent_name": "10_27_67_58",
                    "total_count": 40,
                    "total_size": 180865692
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 28,
                    "total_size": 15528659
                },
                {
                    "agent_id": 100007,
                    "agent_name": "10_27_67_51",
                    "total_count": 272,
                    "total_size": 11723094
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 240,
                    "total_size": 428022152
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 282,
                    "total_size": 11238004
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 36,
                    "total_size": 4737643
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 1268,
                    "total_size": 641384794
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 129,
                    "total_size": 1157908
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 56,
                    "total_size": 28929333
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 512,
                    "total_size": 32343047
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 56,
                    "total_size": 664844
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 15,
                    "total_size": 1645154
                },
                {
                    "agent_id": 100008,
                    "agent_name": "10_27_67_52",
                    "total_count": 20,
                    "total_size": 7448028
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 14,
                    "total_size": 9844722
                },
                {
                    "agent_id": 100015,
                    "agent_name": "10_27_67_53",
                    "total_count": 8,
                    "total_size": 942688
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 1456,
                    "total_size": 487221726
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 320,
                    "total_size": 116905918
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 146,
                    "total_size": 52724615
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 100,
                    "total_size": 736727283
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1560,
                    "total_size": 117527861
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 62,
                    "total_size": 81508795
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 68,
                    "total_size": 92535473
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 76,
                    "total_size": 143791276
                },
                {
                    "agent_id": 100013,
                    "agent_name": "10_27_67_56",
                    "total_count": 112,
                    "total_size": 2412464
                },
                {
                    "agent_id": 100010,
                    "agent_name": "10_27_67_57",
                    "total_count": 28,
                    "total_size": 120509375
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 1084,
                    "total_size": 369937770
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 96,
                    "total_size": 294296837
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 216,
                    "total_size": 89563298
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 13,
                    "total_size": 70816
                }
            ]
        },
        {
            "time": 1721289600000,
            "time_display": "2024-07-18T08:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 13741,
            "total_size": 8115762134,
            "agents": [
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 226,
                    "total_size": 9023056
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 19,
                    "total_size": 103519
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 160,
                    "total_size": 353450673
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1939,
                    "total_size": 146273956
                },
                {
                    "agent_id": 100015,
                    "agent_name": "10_27_67_53",
                    "total_count": 76,
                    "total_size": 9018892
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 160,
                    "total_size": 1877047
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 653,
                    "total_size": 363532472
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 86,
                    "total_size": 317601879
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 32,
                    "total_size": 15942351
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 240,
                    "total_size": 1767839256
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 18,
                    "total_size": 2036732
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 8,
                    "total_size": 5098763
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 1856,
                    "total_size": 629788792
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 551,
                    "total_size": 35511819
                },
                {
                    "agent_id": 100008,
                    "agent_name": "10_27_67_52",
                    "total_count": 88,
                    "total_size": 31594955
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 132,
                    "total_size": 173413021
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 130,
                    "total_size": 96624475
                },
                {
                    "agent_id": 100007,
                    "agent_name": "10_27_67_51",
                    "total_count": 1143,
                    "total_size": 48876490
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 306,
                    "total_size": 110261044
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 530,
                    "total_size": 637972154
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 904,
                    "total_size": 455392826
                },
                {
                    "agent_id": 100012,
                    "agent_name": "10_27_67_58",
                    "total_count": 91,
                    "total_size": 182576504
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 480,
                    "total_size": 872021391
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 26,
                    "total_size": 3418582
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 288,
                    "total_size": 118561291
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 123,
                    "total_size": 1122778
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 402,
                    "total_size": 142314723
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 16,
                    "total_size": 8460275
                },
                {
                    "agent_id": 100013,
                    "agent_name": "10_27_67_56",
                    "total_count": 283,
                    "total_size": 6151474
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 2592,
                    "total_size": 883048776
                },
                {
                    "agent_id": 100010,
                    "agent_name": "10_27_67_57",
                    "total_count": 75,
                    "total_size": 540687653
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 108,
                    "total_size": 146164515
                }
            ]
        },
        {
            "time": 1721282400000,
            "time_display": "2024-07-18T06:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 12799,
            "total_size": 6294990250,
            "agents": [
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 2556,
                    "total_size": 841496201
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 1506,
                    "total_size": 782405216
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 88,
                    "total_size": 195727081
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 270,
                    "total_size": 96920077
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 60,
                    "total_size": 439579070
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 120,
                    "total_size": 217396063
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 120,
                    "total_size": 49383795
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 265,
                    "total_size": 10611074
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 28,
                    "total_size": 15204016
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1270,
                    "total_size": 92476512
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 794,
                    "total_size": 984898484
                },
                {
                    "agent_id": 100008,
                    "agent_name": "10_27_67_52",
                    "total_count": 72,
                    "total_size": 25228820
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 124,
                    "total_size": 162499228
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 1860,
                    "total_size": 633657208
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 96,
                    "total_size": 129641250
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 12,
                    "total_size": 1332199
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 16,
                    "total_size": 9217167
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 730,
                    "total_size": 406129336
                },
                {
                    "agent_id": 100013,
                    "agent_name": "10_27_67_56",
                    "total_count": 252,
                    "total_size": 5451985
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 136,
                    "total_size": 1600624
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 128,
                    "total_size": 45622746
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 13,
                    "total_size": 70628
                },
                {
                    "agent_id": 100007,
                    "agent_name": "10_27_67_51",
                    "total_count": 977,
                    "total_size": 41767857
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 198,
                    "total_size": 361180410
                },
                {
                    "agent_id": 100010,
                    "agent_name": "10_27_67_57",
                    "total_count": 60,
                    "total_size": 360765712
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 411,
                    "total_size": 27327120
                },
                {
                    "agent_id": 100015,
                    "agent_name": "10_27_67_53",
                    "total_count": 68,
                    "total_size": 8074114
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 68,
                    "total_size": 33628017
                },
                {
                    "agent_id": 100012,
                    "agent_name": "10_27_67_58",
                    "total_count": 77,
                    "total_size": 122301074
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 256,
                    "total_size": 187936080
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 32,
                    "total_size": 4221116
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 136,
                    "total_size": 1239970
                }
            ]
        },
        {
            "time": 1721268000000,
            "time_display": "2024-07-18T02:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 11088,
            "total_size": 13521454219,
            "agents": [
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 96,
                    "total_size": 142207787
                },
                {
                    "agent_id": 100015,
                    "agent_name": "10_27_67_53",
                    "total_count": 72,
                    "total_size": 8460631
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 12,
                    "total_size": 6519482
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 1464,
                    "total_size": 494486865
                },
                {
                    "agent_id": 100012,
                    "agent_name": "10_27_67_58",
                    "total_count": 69,
                    "total_size": 62158666
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 233,
                    "total_size": 9322111
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 624,
                    "total_size": 349746474
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 489,
                    "total_size": 31635137
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 94,
                    "total_size": 123245960
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 12,
                    "total_size": 7648519
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 128,
                    "total_size": 281861111
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 134,
                    "total_size": 47854144
                },
                {
                    "agent_id": 100013,
                    "agent_name": "10_27_67_56",
                    "total_count": 280,
                    "total_size": 6046726
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 15,
                    "total_size": 1699570
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 234,
                    "total_size": 84151659
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 16,
                    "total_size": 87174
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 94,
                    "total_size": 126993651
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 636,
                    "total_size": 762030638
                },
                {
                    "agent_id": 100010,
                    "agent_name": "10_27_67_57",
                    "total_count": 347,
                    "total_size": 8386336931
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 114,
                    "total_size": 1020997
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 40,
                    "total_size": 19715941
                },
                {
                    "agent_id": 100007,
                    "agent_name": "10_27_67_51",
                    "total_count": 972,
                    "total_size": 41649252
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 156,
                    "total_size": 115569036
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 120,
                    "total_size": 885411630
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 120,
                    "total_size": 1415305
                },
                {
                    "agent_id": 100008,
                    "agent_name": "10_27_67_52",
                    "total_count": 69,
                    "total_size": 23694419
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 200,
                    "total_size": 366099984
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 1836,
                    "total_size": 609927282
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 120,
                    "total_size": 49495933
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1634,
                    "total_size": 118390306
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 632,
                    "total_size": 353135136
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 26,
                    "total_size": 3435762
                }
            ]
        },
        {
            "time": 1721228400000,
            "time_display": "2024-07-17T15:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 1109,
            "total_size": 510317658,
            "agents": [
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 252,
                    "total_size": 83607410
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 16,
                    "total_size": 5566492
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 540,
                    "total_size": 309550436
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 240,
                    "total_size": 81915294
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 42,
                    "total_size": 24165038
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 4,
                    "total_size": 158479
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 9,
                    "total_size": 80880
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 2,
                    "total_size": 2342
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 4,
                    "total_size": 5271287
                }
            ]
        },
        {
            "time": 1721260800000,
            "time_display": "2024-07-18T00:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 7193,
            "total_size": 15116878760,
            "agents": [
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 108,
                    "total_size": 38511276
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 924,
                    "total_size": 315361149
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 88,
                    "total_size": 3502119
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 16,
                    "total_size": 11335790
                },
                {
                    "agent_id": 100008,
                    "agent_name": "10_27_67_52",
                    "total_count": 68,
                    "total_size": 23421275
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 60,
                    "total_size": 80886291
                },
                {
                    "agent_id": 100010,
                    "agent_name": "10_27_67_57",
                    "total_count": 426,
                    "total_size": 11559647860
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 60,
                    "total_size": 31870708
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 96,
                    "total_size": 212547767
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 2,
                    "total_size": 2596935
                },
                {
                    "agent_id": 100015,
                    "agent_name": "10_27_67_53",
                    "total_count": 68,
                    "total_size": 7945816
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 104,
                    "total_size": 76413571
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 317,
                    "total_size": 20094525
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 32,
                    "total_size": 17306468
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 9,
                    "total_size": 988873
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 722,
                    "total_size": 817777101
                },
                {
                    "agent_id": 100012,
                    "agent_name": "10_27_67_58",
                    "total_count": 66,
                    "total_size": 540808595
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 9,
                    "total_size": 48114
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 1686,
                    "total_size": 932992351
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 72,
                    "total_size": 36300115
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 982,
                    "total_size": 71920298
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 88,
                    "total_size": 267535779
                },
                {
                    "agent_id": 100013,
                    "agent_name": "10_27_67_56",
                    "total_count": 168,
                    "total_size": 3648361
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 10,
                    "total_size": 1328343
                },
                {
                    "agent_id": 100007,
                    "agent_name": "10_27_67_51",
                    "total_count": 972,
                    "total_size": 41636439
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 40,
                    "total_size": 452841
                }
            ]
        },
        {
            "time": 1721253600000,
            "time_display": "2024-07-17T22:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 14490,
            "total_size": 9169883816,
            "agents": [
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 2664,
                    "total_size": 894316833
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 2708,
                    "total_size": 921182755
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 429,
                    "total_size": 26768577
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 13,
                    "total_size": 69567
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 797,
                    "total_size": 445745082
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 22,
                    "total_size": 14449175
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 123,
                    "total_size": 1419655
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 100,
                    "total_size": 198459219
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 164,
                    "total_size": 356550149
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 1166,
                    "total_size": 1398867668
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 14,
                    "total_size": 1575163
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 34,
                    "total_size": 4468820
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 240,
                    "total_size": 1765473661
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 2078,
                    "total_size": 1135347412
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 152,
                    "total_size": 205232945
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 276,
                    "total_size": 113322767
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 286,
                    "total_size": 211132033
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 40,
                    "total_size": 21833539
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 480,
                    "total_size": 863684313
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 352,
                    "total_size": 124036641
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 147,
                    "total_size": 1315235
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 256,
                    "total_size": 10271664
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1363,
                    "total_size": 99628189
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 134,
                    "total_size": 175561447
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 92,
                    "total_size": 48353623
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 360,
                    "total_size": 130817684
                }
            ]
        },
        {
            "time": 1721314800000,
            "time_display": "2024-07-18T15:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 11129,
            "total_size": 7910262327,
            "agents": [
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 472,
                    "total_size": 917618739
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 12,
                    "total_size": 65350
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 240,
                    "total_size": 1776830302
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 36,
                    "total_size": 19576187
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 33,
                    "total_size": 4359325
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 18,
                    "total_size": 13056849
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 136,
                    "total_size": 1234334
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 1712,
                    "total_size": 584527883
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 240,
                    "total_size": 99272455
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 104,
                    "total_size": 320729612
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 68,
                    "total_size": 35435408
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 13,
                    "total_size": 1427311
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 270,
                    "total_size": 197596109
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 1374,
                    "total_size": 698077319
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 297,
                    "total_size": 18911968
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 219,
                    "total_size": 8750986
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 332,
                    "total_size": 119120609
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 128,
                    "total_size": 173012348
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 2164,
                    "total_size": 714036324
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 390,
                    "total_size": 139081383
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 717,
                    "total_size": 399266402
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 134,
                    "total_size": 175723633
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 200,
                    "total_size": 458508880
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 744,
                    "total_size": 958072467
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 996,
                    "total_size": 74986994
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 80,
                    "total_size": 983150
                }
            ]
        },
        {
            "time": 1721300400000,
            "time_display": "2024-07-18T11:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 8727,
            "total_size": 5617542083,
            "agents": [
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 46,
                    "total_size": 60350270
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 126,
                    "total_size": 566955945
                },
                {
                    "agent_id": 100010,
                    "agent_name": "10_27_67_57",
                    "total_count": 30,
                    "total_size": 180385101
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 56,
                    "total_size": 657279
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 144,
                    "total_size": 60566033
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 78,
                    "total_size": 58360281
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 212,
                    "total_size": 256696284
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 162,
                    "total_size": 57964549
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 6,
                    "total_size": 3343358
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 1040,
                    "total_size": 351066848
                },
                {
                    "agent_id": 100007,
                    "agent_name": "10_27_67_51",
                    "total_count": 486,
                    "total_size": 20958846
                },
                {
                    "agent_id": 100015,
                    "agent_name": "10_27_67_53",
                    "total_count": 28,
                    "total_size": 3268393
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 140,
                    "total_size": 1032640421
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 416,
                    "total_size": 31478777
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 720,
                    "total_size": 396428719
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 28,
                    "total_size": 3689306
                },
                {
                    "agent_id": 100013,
                    "agent_name": "10_27_67_56",
                    "total_count": 151,
                    "total_size": 3282756
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 32,
                    "total_size": 14367442
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 5,
                    "total_size": 27223
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 1736,
                    "total_size": 589544822
                },
                {
                    "agent_id": 100012,
                    "agent_name": "10_27_67_58",
                    "total_count": 50,
                    "total_size": 61534489
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 220,
                    "total_size": 8762464
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 708,
                    "total_size": 263359580
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 174,
                    "total_size": 12248433
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 85,
                    "total_size": 289231352
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 1286,
                    "total_size": 645194063
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 84,
                    "total_size": 113736748
                },
                {
                    "agent_id": 100008,
                    "agent_name": "10_27_67_52",
                    "total_count": 28,
                    "total_size": 9977945
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 280,
                    "total_size": 509086637
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 20,
                    "total_size": 10690880
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 3,
                    "total_size": 350892
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 147,
                    "total_size": 1335947
                }
            ]
        },
        {
            "time": 1721296800000,
            "time_display": "2024-07-18T10:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 8302,
            "total_size": 7242946413,
            "agents": [
                {
                    "agent_id": 100012,
                    "agent_name": "10_27_67_58",
                    "total_count": 90,
                    "total_size": 242345840
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 32,
                    "total_size": 10981127
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 115,
                    "total_size": 1036595
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 24,
                    "total_size": 3165503
                },
                {
                    "agent_id": 100015,
                    "agent_name": "10_27_67_53",
                    "total_count": 80,
                    "total_size": 9433301
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 272,
                    "total_size": 91795126
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 322,
                    "total_size": 236267795
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 190,
                    "total_size": 7574042
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 28,
                    "total_size": 9374713
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 16,
                    "total_size": 87182
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 36,
                    "total_size": 12712871
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 1468,
                    "total_size": 747745313
                },
                {
                    "agent_id": 100007,
                    "agent_name": "10_27_67_51",
                    "total_count": 842,
                    "total_size": 36091332
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 642,
                    "total_size": 346452959
                },
                {
                    "agent_id": 100010,
                    "agent_name": "10_27_67_57",
                    "total_count": 142,
                    "total_size": 2636794404
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 128,
                    "total_size": 1551864
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 87,
                    "total_size": 382644498
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 16,
                    "total_size": 1778817
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 56,
                    "total_size": 677334228
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 20,
                    "total_size": 11513160
                },
                {
                    "agent_id": 100008,
                    "agent_name": "10_27_67_52",
                    "total_count": 72,
                    "total_size": 25498494
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1560,
                    "total_size": 117223276
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 1230,
                    "total_size": 1510880222
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 453,
                    "total_size": 28624538
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 24,
                    "total_size": 9890129
                },
                {
                    "agent_id": 100013,
                    "agent_name": "10_27_67_56",
                    "total_count": 241,
                    "total_size": 5202149
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 12,
                    "total_size": 15737059
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 64,
                    "total_size": 31661688
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 12,
                    "total_size": 16238540
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 28,
                    "total_size": 15309648
                }
            ]
        },
        {
            "time": 1721286000000,
            "time_display": "2024-07-18T07:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 11071,
            "total_size": 6274947584,
            "agents": [
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1538,
                    "total_size": 115537873
                },
                {
                    "agent_id": 100012,
                    "agent_name": "10_27_67_58",
                    "total_count": 30,
                    "total_size": 971165
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 120,
                    "total_size": 1085107
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 32,
                    "total_size": 16027761
                },
                {
                    "agent_id": 100007,
                    "agent_name": "10_27_67_51",
                    "total_count": 405,
                    "total_size": 17253271
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 104,
                    "total_size": 1213539
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 238,
                    "total_size": 85297132
                },
                {
                    "agent_id": 100008,
                    "agent_name": "10_27_67_52",
                    "total_count": 28,
                    "total_size": 9897690
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 320,
                    "total_size": 578165687
                },
                {
                    "agent_id": 100010,
                    "agent_name": "10_27_67_57",
                    "total_count": 36,
                    "total_size": 360065991
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 216,
                    "total_size": 89147552
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 30,
                    "total_size": 3947161
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 2336,
                    "total_size": 791227469
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 16,
                    "total_size": 8684967
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 140,
                    "total_size": 1029716014
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 156,
                    "total_size": 117090098
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 160,
                    "total_size": 256746080
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 91,
                    "total_size": 288626839
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 120,
                    "total_size": 162058912
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 252,
                    "total_size": 89571495
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 15,
                    "total_size": 1692392
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 14,
                    "total_size": 76196
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 894,
                    "total_size": 455091546
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 483,
                    "total_size": 30065533
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 1572,
                    "total_size": 521201230
                },
                {
                    "agent_id": 100015,
                    "agent_name": "10_27_67_53",
                    "total_count": 40,
                    "total_size": 4629009
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 620,
                    "total_size": 752691188
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 669,
                    "total_size": 373390044
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 74,
                    "total_size": 97022832
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 231,
                    "total_size": 9216140
                },
                {
                    "agent_id": 100013,
                    "agent_name": "10_27_67_56",
                    "total_count": 81,
                    "total_size": 1692185
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 10,
                    "total_size": 5847486
                }
            ]
        },
        {
            "time": 1721278800000,
            "time_display": "2024-07-18T05:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 10477,
            "total_size": 10073979086,
            "agents": [
                {
                    "agent_id": 100010,
                    "agent_name": "10_27_67_57",
                    "total_count": 204,
                    "total_size": 4134292521
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 84,
                    "total_size": 42236001
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 420,
                    "total_size": 138813139
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 18,
                    "total_size": 3128516
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 456,
                    "total_size": 254623272
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 605,
                    "total_size": 38606105
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 105,
                    "total_size": 960547
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 20,
                    "total_size": 11885942
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 252,
                    "total_size": 89139398
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 20,
                    "total_size": 26211787
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 260,
                    "total_size": 194021254
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 17,
                    "total_size": 1880867
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 54,
                    "total_size": 190311814
                },
                {
                    "agent_id": 100008,
                    "agent_name": "10_27_67_52",
                    "total_count": 77,
                    "total_size": 27599886
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 344,
                    "total_size": 642248164
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 22,
                    "total_size": 29765838
                },
                {
                    "agent_id": 100012,
                    "agent_name": "10_27_67_58",
                    "total_count": 85,
                    "total_size": 122555595
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 2096,
                    "total_size": 1131775469
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 40,
                    "total_size": 21832106
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 1130,
                    "total_size": 1307458869
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 17,
                    "total_size": 2248774
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 117,
                    "total_size": 1389490
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 17,
                    "total_size": 92613
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 72,
                    "total_size": 25895081
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 134,
                    "total_size": 5378447
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 180,
                    "total_size": 1324213829
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 144,
                    "total_size": 59376405
                },
                {
                    "agent_id": 100007,
                    "agent_name": "10_27_67_51",
                    "total_count": 1090,
                    "total_size": 46464173
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1876,
                    "total_size": 136194922
                },
                {
                    "agent_id": 100015,
                    "agent_name": "10_27_67_53",
                    "total_count": 80,
                    "total_size": 9330436
                },
                {
                    "agent_id": 100013,
                    "agent_name": "10_27_67_56",
                    "total_count": 297,
                    "total_size": 6452903
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 144,
                    "total_size": 47594923
                }
            ]
        },
        {
            "time": 1721264400000,
            "time_display": "2024-07-18T01:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 13295,
            "total_size": 9801007211,
            "agents": [
                {
                    "agent_id": 100013,
                    "agent_name": "10_27_67_56",
                    "total_count": 252,
                    "total_size": 5406444
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 288,
                    "total_size": 118759405
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 50,
                    "total_size": 67469331
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 442,
                    "total_size": 247787596
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 605,
                    "total_size": 38992100
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 322,
                    "total_size": 235453320
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 17,
                    "total_size": 1948309
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 121,
                    "total_size": 1084458
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 19,
                    "total_size": 2538764
                },
                {
                    "agent_id": 100008,
                    "agent_name": "10_27_67_52",
                    "total_count": 60,
                    "total_size": 22247125
                },
                {
                    "agent_id": 100010,
                    "agent_name": "10_27_67_57",
                    "total_count": 132,
                    "total_size": 2576616367
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 76,
                    "total_size": 37397623
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 1080,
                    "total_size": 1354075656
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 104,
                    "total_size": 4174280
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 15,
                    "total_size": 80186
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 18,
                    "total_size": 10895467
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 1068,
                    "total_size": 367027643
                },
                {
                    "agent_id": 100015,
                    "agent_name": "10_27_67_53",
                    "total_count": 60,
                    "total_size": 7094848
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 360,
                    "total_size": 654846568
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 44,
                    "total_size": 24087268
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 2300,
                    "total_size": 1276052467
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 109,
                    "total_size": 257665503
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 98,
                    "total_size": 128457807
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 352,
                    "total_size": 125113785
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 64,
                    "total_size": 141336716
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 120,
                    "total_size": 1414522
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 160,
                    "total_size": 1175382596
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 2004,
                    "total_size": 670916659
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 198,
                    "total_size": 72294609
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1929,
                    "total_size": 139666429
                },
                {
                    "agent_id": 100012,
                    "agent_name": "10_27_67_58",
                    "total_count": 72,
                    "total_size": 2421421
                },
                {
                    "agent_id": 100007,
                    "agent_name": "10_27_67_51",
                    "total_count": 756,
                    "total_size": 32301939
                }
            ]
        },
        {
            "time": 1721246400000,
            "time_display": "2024-07-17T20:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 10038,
            "total_size": 5631563642,
            "agents": [
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 1060,
                    "total_size": 1266031621
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 587,
                    "total_size": 36994018
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 244,
                    "total_size": 435281520
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 260,
                    "total_size": 10448704
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 156,
                    "total_size": 115816520
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 104,
                    "total_size": 1198265
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 35,
                    "total_size": 4599924
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 109,
                    "total_size": 191285692
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 68,
                    "total_size": 91726241
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 19,
                    "total_size": 101628
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 792,
                    "total_size": 439746204
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 88,
                    "total_size": 179813752
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 18,
                    "total_size": 12568171
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 48,
                    "total_size": 62857708
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 988,
                    "total_size": 332655054
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 152,
                    "total_size": 62822614
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 236,
                    "total_size": 84059352
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 24,
                    "total_size": 12971965
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 1276,
                    "total_size": 721177164
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 48,
                    "total_size": 24910790
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 144,
                    "total_size": 51575362
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 1396,
                    "total_size": 468970554
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 18,
                    "total_size": 2004790
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1900,
                    "total_size": 138946301
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 148,
                    "total_size": 1327931
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 120,
                    "total_size": 881671797
                }
            ]
        },
        {
            "time": 1721242800000,
            "time_display": "2024-07-17T19:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 11089,
            "total_size": 5854790987,
            "agents": [
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 536,
                    "total_size": 33190723
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 60,
                    "total_size": 135653738
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 230,
                    "total_size": 81473815
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1697,
                    "total_size": 122697021
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 606,
                    "total_size": 336308457
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 2112,
                    "total_size": 709720217
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 11,
                    "total_size": 58832
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 116,
                    "total_size": 151832696
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 140,
                    "total_size": 1030921642
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 32,
                    "total_size": 4216248
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 122,
                    "total_size": 164880957
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 104,
                    "total_size": 76755941
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 162,
                    "total_size": 353982869
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 286,
                    "total_size": 310161063
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 168,
                    "total_size": 69261090
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 1964,
                    "total_size": 665760897
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 14,
                    "total_size": 1542644
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 80,
                    "total_size": 949115
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 48,
                    "total_size": 24926274
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 306,
                    "total_size": 111448927
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 1688,
                    "total_size": 936594674
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 200,
                    "total_size": 8184259
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 93,
                    "total_size": 839920
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 28,
                    "total_size": 15493409
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 280,
                    "total_size": 503735642
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 6,
                    "total_size": 4199917
                }
            ]
        },
        {
            "time": 1721232000000,
            "time_display": "2024-07-17T16:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 13712,
            "total_size": 7668612085,
            "agents": [
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 192,
                    "total_size": 424882694
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 1464,
                    "total_size": 830993453
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 757,
                    "total_size": 897170791
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 136,
                    "total_size": 178485549
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 16,
                    "total_size": 11657804
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 360,
                    "total_size": 648874794
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 2692,
                    "total_size": 910497134
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 60,
                    "total_size": 31357490
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 2680,
                    "total_size": 889057746
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 150,
                    "total_size": 203258259
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 380,
                    "total_size": 137634155
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 128,
                    "total_size": 1504057
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 32,
                    "total_size": 17797552
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 182,
                    "total_size": 133428641
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 17,
                    "total_size": 1886345
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 146,
                    "total_size": 1314038
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 264,
                    "total_size": 108638176
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1791,
                    "total_size": 130586221
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 795,
                    "total_size": 446289828
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 33,
                    "total_size": 4363591
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 180,
                    "total_size": 1323679582
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 567,
                    "total_size": 35731913
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 233,
                    "total_size": 9373801
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 110,
                    "total_size": 173635919
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 17,
                    "total_size": 90913
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 330,
                    "total_size": 116421639
                }
            ]
        },
        {
            "time": 1721307600000,
            "time_display": "2024-07-18T13:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 13650,
            "total_size": 8273347471,
            "agents": [
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 162,
                    "total_size": 417389994
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1872,
                    "total_size": 140514339
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 420,
                    "total_size": 760930807
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 168,
                    "total_size": 1970391
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 133,
                    "total_size": 1207305
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 322,
                    "total_size": 116189876
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 128,
                    "total_size": 168093556
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 250,
                    "total_size": 9938790
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 216,
                    "total_size": 89211812
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 2544,
                    "total_size": 844879825
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 2128,
                    "total_size": 1078296620
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 1972,
                    "total_size": 679410117
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 742,
                    "total_size": 408970200
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 16,
                    "total_size": 10493302
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 798,
                    "total_size": 928757889
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 34,
                    "total_size": 4461906
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 220,
                    "total_size": 1615975806
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 18,
                    "total_size": 1955506
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 605,
                    "total_size": 39605656
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 116,
                    "total_size": 157402616
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 208,
                    "total_size": 153383512
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 72,
                    "total_size": 37320774
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 342,
                    "total_size": 122425837
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 105,
                    "total_size": 463262627
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 40,
                    "total_size": 21194917
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 19,
                    "total_size": 103491
                }
            ]
        },
        {
            "time": 1721293200000,
            "time_display": "2024-07-18T09:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 10264,
            "total_size": 7672114374,
            "agents": [
                {
                    "agent_id": 100010,
                    "agent_name": "10_27_67_57",
                    "total_count": 120,
                    "total_size": 2396719938
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 87,
                    "total_size": 795104
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 122,
                    "total_size": 159965265
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 200,
                    "total_size": 367249115
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 159,
                    "total_size": 6411197
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 120,
                    "total_size": 884417930
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 9,
                    "total_size": 49045
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 48,
                    "total_size": 19774491
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 973,
                    "total_size": 73227315
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 306,
                    "total_size": 109594060
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 10,
                    "total_size": 6257412
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 10,
                    "total_size": 1141597
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 20,
                    "total_size": 2630591
                },
                {
                    "agent_id": 100008,
                    "agent_name": "10_27_67_52",
                    "total_count": 20,
                    "total_size": 7272663
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 1876,
                    "total_size": 638874810
                },
                {
                    "agent_id": 100012,
                    "agent_name": "10_27_67_58",
                    "total_count": 22,
                    "total_size": 767463
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 26,
                    "total_size": 19112559
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 144,
                    "total_size": 318769628
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 64,
                    "total_size": 758283
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 254,
                    "total_size": 280845341
                },
                {
                    "agent_id": 100007,
                    "agent_name": "10_27_67_51",
                    "total_count": 346,
                    "total_size": 14778180
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 564,
                    "total_size": 305838634
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 76,
                    "total_size": 287400120
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 24,
                    "total_size": 13014866
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 96,
                    "total_size": 34592405
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 2620,
                    "total_size": 885604979
                },
                {
                    "agent_id": 100015,
                    "agent_name": "10_27_67_53",
                    "total_count": 24,
                    "total_size": 2804350
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 1236,
                    "total_size": 626648267
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 112,
                    "total_size": 151332250
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 56,
                    "total_size": 28061601
                },
                {
                    "agent_id": 100013,
                    "agent_name": "10_27_67_56",
                    "total_count": 168,
                    "total_size": 3606586
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 352,
                    "total_size": 23798329
                }
            ]
        },
        {
            "time": 1721271600000,
            "time_display": "2024-07-18T03:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 7281,
            "total_size": 6120456180,
            "agents": [
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 198,
                    "total_size": 71104934
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 120,
                    "total_size": 49486655
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 7,
                    "total_size": 38122
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 700,
                    "total_size": 51422798
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 118,
                    "total_size": 1073353
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 120,
                    "total_size": 881295757
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 64,
                    "total_size": 750905
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 231,
                    "total_size": 15061037
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 296,
                    "total_size": 162493827
                },
                {
                    "agent_id": 100015,
                    "agent_name": "10_27_67_53",
                    "total_count": 20,
                    "total_size": 2312319
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 70,
                    "total_size": 247931917
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 28,
                    "total_size": 3691736
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 80,
                    "total_size": 104790532
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 7,
                    "total_size": 777946
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 98,
                    "total_size": 212947663
                },
                {
                    "agent_id": 100013,
                    "agent_name": "10_27_67_56",
                    "total_count": 112,
                    "total_size": 2434101
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 14,
                    "total_size": 8088199
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 2,
                    "total_size": 1585315
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 1744,
                    "total_size": 575853970
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 198,
                    "total_size": 71692440
                },
                {
                    "agent_id": 100007,
                    "agent_name": "10_27_67_51",
                    "total_count": 324,
                    "total_size": 13829133
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 4,
                    "total_size": 2136612
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 684,
                    "total_size": 383868149
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 240,
                    "total_size": 444198430
                },
                {
                    "agent_id": 100008,
                    "agent_name": "10_27_67_52",
                    "total_count": 28,
                    "total_size": 10391500
                },
                {
                    "agent_id": 100010,
                    "agent_name": "10_27_67_57",
                    "total_count": 85,
                    "total_size": 2096497275
                },
                {
                    "agent_id": 100012,
                    "agent_name": "10_27_67_58",
                    "total_count": 30,
                    "total_size": 120701268
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 197,
                    "total_size": 7938459
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 8,
                    "total_size": 3985334
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 1376,
                    "total_size": 466734510
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 78,
                    "total_size": 105341984
                }
            ]
        },
        {
            "time": 1721250000000,
            "time_display": "2024-07-17T21:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 13730,
            "total_size": 8142877737,
            "agents": [
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 157,
                    "total_size": 1795381
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 306,
                    "total_size": 112255189
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 16,
                    "total_size": 1774929
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 1898,
                    "total_size": 1068304109
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 561,
                    "total_size": 34615604
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 108,
                    "total_size": 267163051
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 200,
                    "total_size": 1470012076
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 124,
                    "total_size": 167378909
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 36,
                    "total_size": 19628542
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 244,
                    "total_size": 100700605
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 268,
                    "total_size": 10755501
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 18,
                    "total_size": 13556659
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 2444,
                    "total_size": 825826710
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1777,
                    "total_size": 129025707
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 33,
                    "total_size": 4335329
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 116,
                    "total_size": 151935381
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 158,
                    "total_size": 321336488
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 2304,
                    "total_size": 780769069
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 72,
                    "total_size": 37682787
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 972,
                    "total_size": 1159132187
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 396,
                    "total_size": 723506380
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 332,
                    "total_size": 117786045
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 143,
                    "total_size": 1276680
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 260,
                    "total_size": 191751253
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 17,
                    "total_size": 90899
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 770,
                    "total_size": 430482267
                }
            ]
        },
        {
            "time": 1721275200000,
            "time_display": "2024-07-18T04:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 5948,
            "total_size": 8845524905,
            "agents": [
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 56,
                    "total_size": 79636145
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 85,
                    "total_size": 263381943
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 24,
                    "total_size": 13256026
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 20,
                    "total_size": 146831439
                },
                {
                    "agent_id": 100007,
                    "agent_name": "10_27_67_51",
                    "total_count": 612,
                    "total_size": 26090150
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 173,
                    "total_size": 10891987
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 96,
                    "total_size": 39733879
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 6,
                    "total_size": 3528108
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 256,
                    "total_size": 84669464
                },
                {
                    "agent_id": 100010,
                    "agent_name": "10_27_67_57",
                    "total_count": 242,
                    "total_size": 6169643745
                },
                {
                    "agent_id": 100015,
                    "agent_name": "10_27_67_53",
                    "total_count": 48,
                    "total_size": 5743752
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 637,
                    "total_size": 358732496
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 74,
                    "total_size": 26028772
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 550,
                    "total_size": 40073261
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 720,
                    "total_size": 244359473
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 6,
                    "total_size": 681021
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 28,
                    "total_size": 70128657
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 23,
                    "total_size": 3038700
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 28,
                    "total_size": 37852869
                },
                {
                    "agent_id": 100013,
                    "agent_name": "10_27_67_56",
                    "total_count": 151,
                    "total_size": 3216386
                },
                {
                    "agent_id": 100012,
                    "agent_name": "10_27_67_58",
                    "total_count": 51,
                    "total_size": 61531585
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 5,
                    "total_size": 27253
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 78,
                    "total_size": 58311189
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 188,
                    "total_size": 7634414
                },
                {
                    "agent_id": 100008,
                    "agent_name": "10_27_67_52",
                    "total_count": 44,
                    "total_size": 15780346
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 1182,
                    "total_size": 642522035
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 304,
                    "total_size": 373724929
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 59,
                    "total_size": 699947
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 28,
                    "total_size": 14138892
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 54,
                    "total_size": 19142552
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 102,
                    "total_size": 921906
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 18,
                    "total_size": 23571584
                }
            ]
        },
        {
            "time": 1721257200000,
            "time_display": "2024-07-17T23:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 8622,
            "total_size": 5295059241,
            "agents": [
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 742,
                    "total_size": 893673207
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 128,
                    "total_size": 307049628
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 96,
                    "total_size": 864489
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 17,
                    "total_size": 2255807
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 56,
                    "total_size": 27763731
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 72,
                    "total_size": 29772781
                },
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 152,
                    "total_size": 6085068
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 14,
                    "total_size": 74911
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 860,
                    "total_size": 466354917
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 120,
                    "total_size": 880283323
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 1396,
                    "total_size": 472393983
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 1648,
                    "total_size": 563240040
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 14,
                    "total_size": 9339680
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 12,
                    "total_size": 1356746
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1168,
                    "total_size": 86002643
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 583,
                    "total_size": 328129924
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 180,
                    "total_size": 63497514
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 112,
                    "total_size": 1288045
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 20,
                    "total_size": 11166195
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 96,
                    "total_size": 214189907
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 198,
                    "total_size": 72095505
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 80,
                    "total_size": 108101958
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 86,
                    "total_size": 112612137
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 260,
                    "total_size": 192748115
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 240,
                    "total_size": 428136666
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 272,
                    "total_size": 16582321
                }
            ]
        },
        {
            "time": 1721239200000,
            "time_display": "2024-07-17T18:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 10533,
            "total_size": 8197906806,
            "agents": [
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 84,
                    "total_size": 44000414
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 338,
                    "total_size": 248111355
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 36,
                    "total_size": 19670624
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 15,
                    "total_size": 133030
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 434,
                    "total_size": 154125201
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 1684,
                    "total_size": 929818343
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 220,
                    "total_size": 1618122901
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 440,
                    "total_size": 790898285
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 312,
                    "total_size": 129872345
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 126,
                    "total_size": 164915001
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 306,
                    "total_size": 111057182
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 148,
                    "total_size": 445055440
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 20,
                    "total_size": 14114889
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 30,
                    "total_size": 17073071
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 56,
                    "total_size": 187661893
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 2672,
                    "total_size": 887951416
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 2272,
                    "total_size": 780435786
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 1228,
                    "total_size": 1503688111
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 112,
                    "total_size": 151201519
                }
            ]
        },
        {
            "time": 1721235600000,
            "time_display": "2024-07-17T17:00:00.000000Z",
            "index": "ocpf-eos-metadata-100000-100003-43",
            "datasource_id": 100003,
            "datasource_name": "yc2svolt",
            "resource_id": 43,
            "resource_name": "yc2svolt",
            "total_count": 12867,
            "total_size": 7646840491,
            "agents": [
                {
                    "agent_id": 100005,
                    "agent_name": "10_27_76_40",
                    "total_count": 247,
                    "total_size": 9973419
                },
                {
                    "agent_id": 100033,
                    "agent_name": "10_27_77_11",
                    "total_count": 242,
                    "total_size": 85260232
                },
                {
                    "agent_id": 100006,
                    "agent_name": "10_27_76_44",
                    "total_count": 1051,
                    "total_size": 76524651
                },
                {
                    "agent_id": 100016,
                    "agent_name": "10_27_76_49",
                    "total_count": 13,
                    "total_size": 69534
                },
                {
                    "agent_id": 100032,
                    "agent_name": "10_27_77_6",
                    "total_count": 2532,
                    "total_size": 855340521
                },
                {
                    "agent_id": 100004,
                    "agent_name": "10_27_76_42",
                    "total_count": 324,
                    "total_size": 20409752
                },
                {
                    "agent_id": 100020,
                    "agent_name": "10_27_76_48",
                    "total_count": 136,
                    "total_size": 1581171
                },
                {
                    "agent_id": 100017,
                    "agent_name": "10_27_76_47",
                    "total_count": 11,
                    "total_size": 1241680
                },
                {
                    "agent_id": 100022,
                    "agent_name": "10_27_77_4",
                    "total_count": 2528,
                    "total_size": 860676005
                },
                {
                    "agent_id": 100036,
                    "agent_name": "10_27_77_7",
                    "total_count": 144,
                    "total_size": 194487945
                },
                {
                    "agent_id": 100026,
                    "agent_name": "10_27_77_18",
                    "total_count": 16,
                    "total_size": 11105527
                },
                {
                    "agent_id": 100024,
                    "agent_name": "10_27_77_19",
                    "total_count": 132,
                    "total_size": 100573722
                },
                {
                    "agent_id": 100003,
                    "agent_name": "10_27_76_41",
                    "total_count": 27,
                    "total_size": 3552417
                },
                {
                    "agent_id": 100035,
                    "agent_name": "10_27_77_12",
                    "total_count": 320,
                    "total_size": 570802110
                },
                {
                    "agent_id": 100018,
                    "agent_name": "10_27_76_39",
                    "total_count": 145,
                    "total_size": 1293680
                },
                {
                    "agent_id": 100023,
                    "agent_name": "10_27_77_8",
                    "total_count": 192,
                    "total_size": 424589323
                },
                {
                    "agent_id": 100029,
                    "agent_name": "10_27_77_13",
                    "total_count": 160,
                    "total_size": 1176255584
                },
                {
                    "agent_id": 100039,
                    "agent_name": "10_27_70_117",
                    "total_count": 93,
                    "total_size": 176788653
                },
                {
                    "agent_id": 100027,
                    "agent_name": "10_27_77_16",
                    "total_count": 68,
                    "total_size": 35251409
                },
                {
                    "agent_id": 100037,
                    "agent_name": "10_27_77_5",
                    "total_count": 138,
                    "total_size": 180486043
                },
                {
                    "agent_id": 100028,
                    "agent_name": "10_27_77_15",
                    "total_count": 2310,
                    "total_size": 1330481166
                },
                {
                    "agent_id": 100025,
                    "agent_name": "10_27_77_20",
                    "total_count": 727,
                    "total_size": 882557771
                },
                {
                    "agent_id": 100034,
                    "agent_name": "10_27_77_14",
                    "total_count": 40,
                    "total_size": 21821860
                },
                {
                    "agent_id": 100019,
                    "agent_name": "10_27_76_37",
                    "total_count": 791,
                    "total_size": 444505654
                },
                {
                    "agent_id": 100021,
                    "agent_name": "10_27_77_3",
                    "total_count": 360,
                    "total_size": 131668730
                },
                {
                    "agent_id": 100030,
                    "agent_name": "10_27_77_10",
                    "total_count": 120,
                    "total_size": 49541932
                }
            ]
        }
    ]`
	result := make([]*es.AggregateByResource, 0)
	err := json.Unmarshal([]byte(j), &result)
	if err != nil {
		log.Errorf("unmarshal error: %v", err)
		return
	}
	for _, r := range result {
		for _, a := range r.Agents {
			if a.AgentName == "10_27_76_39" {
				fmt.Printf("%s   %f\n", r.TimeDisplay, a.TotalCount)
			}
		}
	}
}

type User struct {
	Name string
}

func grokTest() {
	key := "CL03_预焊裁切_铝极长度_J0093B0243740_20240730174233-20240730-174233-469.jpg"
	regex := ".*/?(?P<jzsn2>J0.{3}(A|B)\\d{7}).*$"
	g, err := grok.New()
	if err != nil {
		log.Fatalf("new grok error: %v", err)
	}
	parse, err := g.Parse(regex, key)
	if err != nil {
		log.Fatalf("key: %s, ParseGrok error: %v", key, err)
	}
	for k, v := range parse {
		println(k, v)
	}
}
