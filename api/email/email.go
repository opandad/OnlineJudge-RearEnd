/*
	@Title
	api/email.go

	@Description
	封装邮件发送，通过STMP协议

	@Func List（这个需打开函数检查）

	| func name          | develop  | unit test |

	|-------------------------------------------|

	| SendMailByQQ  |    ok    |    ok	    |

	@config
	email path => ~/configs/email.go
*/
package email

import (
	"OnlineJudge-RearEnd/configs"
	"fmt"
	"net/smtp"
	"strings"
)

/*
@Title
SendMailByQQ

@description
发送qq邮件给收件人，可以单发和群发，可用于验证码单发、比赛信息群发

@param
to, nickname, subject, msg ([]string, string, string, string)

param e.g.

单发：SendMailByQQ([]string{"收件人"}，"自己发送的昵称"，"主题"，"消息")

群发：SendMailByQQ({"收件人1","收件人2",...,"收件人n"}，"自己发送的昵称"，"主题"，"消息")

@return
成功或失败 bool
*/
func SendMailByQQ(to []string, nickname string, subject string, msg string) bool {
	auth := smtp.PlainAuth("", configs.EMAIL_STMP_ACCOUNT, configs.EMAIL_STMP_PASSWORD, configs.EMAIL_STMP_SERVER_HOSTNAME)
	/*
		to 发送给
		from 从谁发送 nickname+xxx（谁发，从哪个地方发）
		subject 主题
		contentType 发送的编码格式
	*/
	contentType := "Content-Type: text/plain; charset=UTF-8"
	finalMsg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + configs.EMAIL_STMP_ACCOUNT + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + msg)

	err := smtp.SendMail(configs.EMAIL_STMP_SERVER_HOSTNAME+":"+configs.EMAIL_STMP_SERVER_PORT, auth, configs.EMAIL_STMP_ACCOUNT, to, finalMsg)

	if err != nil {
		fmt.Println("发送给：" + strings.Join(to, ",") + "的邮件失败！")
		fmt.Println(err)
		return false
	}
	return true
}
