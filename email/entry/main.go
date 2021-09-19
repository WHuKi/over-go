package main

/*
@brief 使用企业服务器邮件发送
*/
func main() {
	////3. 邮箱信息发送
	//emailInfo := email.EmailInfo{
	//	AuthInfo: email.EmailFormAuthInfo{
	//		EmailHost:     epidemicEmail.Host,
	//		EmailPort:     int(epidemicEmail.Port),
	//		EmailUserName: epidemicEmail.UserName,
	//		EmailAuthCode: epidemicEmail.Password,
	//	},
	//	From:           epidemicEmail.From,
	//	To:             epidemicEmailTo,
	//	Cc:             epidemicEmailCc,
	//	Subject:        epidemicEmail.Subject,
	//	Text:           []byte(epidemicEmail.Text),
	//	AttachFilePath: fmt.Sprintf("excel/epidemic%v.xlsx", time.Now().Format(constant.TIME_TYPE_ONE)),
	//}
	//if err = emailInfo.Send(); err != nil {
	//	utils.Logger().Error("email send fail", zap.Error(err))
	//	return err
	//}
}
