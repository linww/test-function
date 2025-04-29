package main

import (
	"fmt"
	"xsky.com/ocpf/pkg/util"
)

func main() {
	//usernameTpl := "xsky-sub01"
	//pwd, _ := util.AesEncrypt(usernameTpl, "Xsky@2023")
	//fmt.Println(pwd)
	MD5()
}

func MD5() {
	fmt.Println(util.Md5Encrypt("d580b460f5a86c2a1d6ddac60d5cb260", "1716889410"))
}
