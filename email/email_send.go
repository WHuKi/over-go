package email

import (
	"fmt"
)

/*
@brief 邮件发送
*/

//EmailInfo 邮件信息
type EmailInfo struct {
	// 授权信息
	AuthInfo EmailFormAuthInfo
	// 发送者
	From string `json:"from"` // xxx@163.com
	// 接收者
	To []string `json:"to"` // xxx@163.com
	// 抄送者
	Cc []string `json:"cc"` // xxx@163.com
	// 邮件主题
	Subject string `json:"subject"` // 请假
	// 邮件内容
	Text []byte `json:"text"` // 因事。。
	// 邮件附件路径
	AttachFilePath string `json:"attach_file_path"` // xxx.txt
}

//EmailFormAuthInfo 发送者授权信息
type EmailFormAuthInfo struct {
	// 邮箱服务器
	EmailHost string // smtp.163.com
	// 端口号
	EmailPort int // 25
	// 用户名
	EmailUserName string // xxx@163.com
	// 授权码
	EmailAuthCode string // WMOZABGULTPPP
}

/*
Send
@brief 邮件发送
*/
func (ei *EmailInfo) Send() (err error) {
	e := NewEmail()
	e.From = fmt.Sprintf("大群 <%s>", ei.From) //"大群 <daqunchat@bianfeng.com>"
	e.To = ei.To
	e.Cc = ei.Cc
	e.Subject = ei.Subject
	e.Text = ei.Text
	_, err = e.AttachFile(ei.AttachFilePath)
	if err != nil {
		return err
	}

	err = e.Send(fmt.Sprintf("%s:%d", ei.AuthInfo.EmailHost, ei.AuthInfo.EmailPort), LoginAuth(ei.AuthInfo.EmailUserName, ei.AuthInfo.EmailAuthCode))
	if err != nil {
		return err
	}

	return
}
