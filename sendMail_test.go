// @Author: abbeymart | Abi Akindele | @Created: 2020-12-29 | @Updated: 2020-12-29
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcmail

import (
	"fmt"
	"github.com/abbeymart/mcmail/config"
	"github.com/abbeymart/mctest"
	"strings"
	"testing"
)

func TestResMessage(t *testing.T) {
	// encrypt appConfig
	// export MCAPP_SECURE_KEY=32-bytes-alphanumeric-value | must match the encryption secureKey 32-bytes-alphanumeric-value
	config.LoadConfig()

	//msgText := "Test message - it's a wonderful world!!!"
	msgHtml := "Hello <b>Guest</b>:<br/><br/><hr/> <h3>Welcome to mConnect Marketplace!</h3><br/><br/><hr/>"
	msgHtml += fmt.Sprintf("%v", config.AppConfig.AppContact)
	subject := "mConnect Go Universal - TESTING!!!"

	// email server information/instance
	mailer := EmailConfigType{
		Username:           config.AppConfig.EmailService.EmailUsername,
		Password:           config.AppConfig.EmailService.EmailPassword,
		Port:               config.AppConfig.EmailService.EmailPort,
		ServerUrl:          config.AppConfig.EmailService.EmailServerUrl,
		MsgFrom:            config.AppConfig.EmailService.EmailMsgFrom,
		InsecureSkipVerify: true,
	}

	responseMessage := "Email message successfully sent"

	mctest.McTest(mctest.OptionValue{
		Name: "should return success code for sending email message",
		TestFunc: func() {
			res := mailer.SendEmail(config.ToEmail, msgHtml, subject, "html")
			mctest.AssertEquals(t, res.Code, "success", "response-code should be: success")
			mctest.AssertEquals(t, strings.Contains(res.Message, responseMessage), true, "response-message should includes/contains"+responseMessage)
		},
	})
	mctest.PostTestResult()
}
