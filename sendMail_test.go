// @Author: abbeymart | Abi Akindele | @Created: 2020-12-29 | @Updated: 2020-12-29
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcmail

import (
	"github.com/abbeymart/mcmail/secure"
	"github.com/abbeymart/mctest"
	"strings"
	"testing"
)

func TestResMessage(t *testing.T) {
	// test-data
	to := secure.ToEmail
	//msg := `"To: example@yahoo.com\r\n" +
	//	"Subject: discount Gophers!\r\n" +
	//	"\r\n" +
	//	"This is the email body.\r\n Enjoy it!!! +" +
	//	" \r\n" +
	//	" \r\n" +
	//	"\r\n" +
	//	" With Regards,\r\n" +
	//	"mac.com Support Team\r\n"`

	//msgText := "Test message - it's a wonderful world!!!"
	msgHtml := "Hello <b>Guest</b> and <i>Welcome to mac</i>!"
	subject :=  "mConnect Go Universal!!!"

	// email server information/instance
	mailer := EmailConfigType{
		Username:  secure.Sendinblue.Username,
		Password:  secure.Sendinblue.Password,
		Port:      secure.Sendinblue.Port,
		ServerUrl: secure.Sendinblue.ServerUrl,
		MsgFrom:   secure.Sendinblue.MsgFrom,
	}

	responseMessage := "Email message successfully sent"

	mctest.McTest(mctest.OptionValue{
		Name: "should return success code for sending email message",
		TestFunc: func() {
			res := mailer.SendEmail(to, msgHtml, subject, "html")
			mctest.AssertEquals(t, res.Code, "success", "response-code should be: success")
			mctest.AssertEquals(t, strings.Contains(res.Message, "Email message successfully sent"), true, "response-message should includes/contains"+responseMessage)
			// gomail: could not send email 1: 554 Message rejected:
			// Sending paused for this account. For more information, please check the inbox
			// of the email address associated with your AWS account.
		},
	})
	mctest.PostTestResult()
}
