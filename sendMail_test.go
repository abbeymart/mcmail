// @Author: abbeymart | Abi Akindele | @Created: 2020-12-29 | @Updated: 2020-12-29
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcmail

import (
	"fmt"
	"strings"
	"testing"

	"github.com/abbeymart/mcmail/config"
	"github.com/abbeymart/mctest"
)

func TestResMessage(t *testing.T) {
	// encrypt appConfig
	// export MCAPP_SECURE_KEY=32-bytes-alphanumeric-value | must match the encryption secureKey 32-bytes-alphanumeric-value
	config.LoadConfig()

	msgHtml := "Hello <b>Guest</b>:<br/><br/><hr/> <h3>Welcome to mConnect Marketplace!</h3><br/><br/><hr/>"
	msgHtml += fmt.Sprintf("%v", config.AppConfig.AppContact)
	subject := "mConnect Go Universal - TESTING [html]!!!"

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

	var results []mctest.UnitTestResult

	// HTML message
	test1 := mctest.NewTest(mctest.ParamsType{
		Name: "should return success code for sending email message [html]",
	})
	res := mailer.SendEmail(config.ToEmail, msgHtml, subject, "html")
	test1.SetTestFunction(func() {
		test1.AssertEquals(res.Code, "success", "response-code should be: success")
		test1.AssertEquals(strings.Contains(res.Message, responseMessage), true, "response-message should includes/contains"+responseMessage)
	})
	test1Result := test1.RunTest()
	results = append(results, test1Result)

	mctest.TestResult(results)
}
