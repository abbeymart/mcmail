// @Author: abbeymart | Abi Akindele | @Created: 2020-12-29 | @Updated: 2020-12-29
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcmail

import (
	"github.com/abbeymart/mctest"
	"strings"
	"testing"
)

func TestResMessage(t *testing.T) {
	// test-data
	to := []string{"abbeya1@yahoo.com", "abiakindele@gmail.com"}
	//msg := `"To: abbeya1@yahoo.com\r\n" +
	//	"Subject: discount Gophers!\r\n" +
	//	"\r\n" +
	//	"This is the email body.\r\n Enjoy it!!! +" +
	//	" \r\n" +
	//	" \r\n" +
	//	"\r\n" +
	//	" With Regards,\r\n" +
	//	"mConnect.biz Support Team\r\n"`

	//msgText := "Test message - it's a wonderful world!!!"
	msgHtml := "Hello <b>Guest</b> and <i>Welcome to mConnect</i>!"
	subject :=  "mConnect Go Universal!!!"

	// email server information/instance
	mailer := EmailConfigType{
		Username:  sendinblue.Username,
		Password:  sendinblue.Password,
		Port:      sendinblue.Port,
		ServerUrl: sendinblue.ServerUrl,
		MsgFrom:   sendinblue.MsgFrom,
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
