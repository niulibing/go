package model

import "fmt"

/**
 * 邮件结构体
 */
type EmailParam struct {
	/**
	 * 邮箱服务器端口，如腾讯邮箱为465
	 */
	ServerHost string
	/**
	 * 邮件发送服务器端口号
	 */
	ServerPort int
	/**
	 * 发件人邮箱地址
	 */
	FromEmail string

	/**
	 * 发件人邮箱密码
	 */
	FromPasswd string
	/**
	 * 邮件接收者 ，如果有多个，则以英文的逗号隔开，不能为空
	 */
	Toers string
	/**
	 * 抄送着邮件地址，如果有多个，则以英文的逗号隔开，可以为空
	 */
	CCers string
}

type File struct {
	CanOrNo bool
}

func FileConstruct(bo bool) *File {

	return &File{CanOrNo: bo}
}

func (f *File) CanWrite(bo interface{}) bool {

	if f.CanOrNo {
		fmt.Printf("允许数据写入。。。。")
		return true
	}
	fmt.Printf("不允许数据写入....")
	return false
}

func (f *File) WriteData(data interface{}) error {

	//模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}
