package main

import "xsky.com/ocpf/pkg/util"

func main() {
	encryptedPwd := util.Md5Encrypt("Xsky@123", "xsky")
	encryptedPwd = util.Md5Encrypt(encryptedPwd, "1716889410")
	println(encryptedPwd)
	pwd, _ := util.AesEncrypt("xsky", "Xsky@123")
	println(pwd)
}
