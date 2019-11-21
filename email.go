package main

import "fmt"

type EmailParam struct {
	//邮箱服务地址
	ServerHost string
	//邮箱服务器端口
	ServerPort int
	//发件人邮箱地址
	FromEmail string
	//发件人邮箱密码
	FromPasswd string
	//接收邮件者，如有多个，以英文逗号隔开，不能为空
	Toers string
	//抄送者邮件，如有多个，以英文逗号隔开，可以为空
	CCers string
}

func Map() {

	//创建一个map

	strings := make(map[string]string)

	strings["name"] = "niulibing"
	strings["sex"] = "nan"
	fmt.Printf("map的值为%v", strings)
}
