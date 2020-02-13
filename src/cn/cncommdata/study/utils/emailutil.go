package utils

import (
	"github.com/go-gomail/gomail"
	"message/src/cn/cncommdata/study/model"
	"strings"
)

/**
email参数结构体
*/

/**
 * 1。全局变量，因为发件人账号，密码，需要在发送时才能指定，注意，由于是小写，外面的包无法使用
 */
var serverHost, fromEmail, fromPasswd string
var serverPort int

var m *gomail.Message

/**
 * 初始化email结构体
 */
func InitEmail(ep *model.EmailParam) {

	/**
	 * 声明切片，用来装邮件接收者
	 */
	var toers []string

	/**
	 * 设置邮件发送的基本信息：发信协议、端口号、发件人账号、发件人账号密码
	 */
	serverHost = ep.ServerHost
	serverPort = ep.ServerPort
	fromEmail = ep.FromEmail
	fromPasswd = ep.FromPasswd

	m = gomail.NewMessage()

	/**
	 * 如果邮件接收者为空，则不发送
	 */
	if len(ep.Toers) == 0 {
		return
	}

	//生成切片
	split := strings.Split(ep.Toers, ",")
	for _, v := range split {
		toers = append(toers, strings.TrimSpace(v))
	}

	/**
	 * 收件人可以有多个，故用此方式
	 */
	m.SetHeader("To", toers...)

	/**
	 * 抄送列表
	 */
	if len(ep.CCers) != 0 {
		for _, tmp := range strings.Split(ep.CCers, ",") {
			toers = append(toers, strings.TrimSpace(tmp))
		}
		m.SetHeader("Cc", toers...)
	}

	/**
	 * 第三个参数为发件人别名。如："牛娃娃"，可以为空（此时则为邮箱名称）
	 */
	m.SetAddressHeader("From", fromEmail, "LAZY")
}

// SendEmail body支持html格式字符串
func SendEmail(subject, body string) {

	// 主题
	m.SetHeader("Subject", subject)

	// 正文
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer(serverHost, serverPort, fromEmail, fromPasswd)
	// 发送
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
}

func Send() {

	serverHost := "smtp.exmail.qq.com"
	serverPort := 465
	fromEmail := "libing.niu@cncommdata.cn" //发件人邮箱
	fromPasswd := "20190324Cd"              //授权码
	myToers := "libing_niu@163.com"         // 收件人邮箱，逗号隔开
	myCCers := "luhao.liu@cncommdata.cn"    //"readchy@163.com"

	subject := "Go语言测试"
	body := `你好，Go语言。很高兴能够有机会认识你，希望今后我们能够相亲相爱成为一家人。`
	// 结构体赋值
	myEmail := &model.EmailParam{
		ServerHost: serverHost,
		ServerPort: serverPort,
		FromEmail:  fromEmail,
		FromPasswd: fromPasswd,
		Toers:      myToers,
		CCers:      myCCers,
	}

	InitEmail(myEmail)
	SendEmail(subject, body)
}
