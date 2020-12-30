// @Author: abbeymart | Abi Akindele | @Created: 2020-12-29 | @Updated: 2020-12-29
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcmail

import (
	"github.com/abbeymart/mctest"
	"testing"
)

func TestResMessage(t *testing.T) {
	// test-data
	to := []string{"abbeya1@yahoo.com", "abiakindele@gmail.com"}
	msg := []byte("To: abbeya1@yahoo.com\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n Enjoy it!!! +" +
		" \r\n" +
		" \r\n" +
		"\r\n" +
		" With Regards,\r\n" +
		"mConnect.biz Support Team\r\n")

	mctest.McTest(mctest.OptionValue{
		Name: "should return success code for sending email message",
		TestFunc: func() {
			//req := GetResMessage(msgType, options)
			//mctest.AssertEquals(t, req.Code, res.Code, "response-code should be: " + res.Code)
			//mctest.AssertEquals(t, req.Message, res.Message, "response-message should be: "+res.Message)
		},
	})
	mctest.PostTestResult()
}
