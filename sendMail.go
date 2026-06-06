// @Author: abbeymart | Abi Akindele | @Created: 2020-12-29 | @Updated: 2020-12-29, 2026-06-06
// @Company: mConnect.biz | @License: MIT
// @Description: send-mail function

package mcmail

import (
	"crypto/tls"
	"fmt"

	"github.com/abbeymart/mcresponse"
	"gopkg.in/gomail.v2"
)

// SendEmail sends text and HTML messages, attachment etc.
func (mailer EmailConfigType) SendEmail(recipients []string, message string, subject string, emailType string) mcresponse.ResponseMessage {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", mailer.MsgFrom)

	// Set E-Mail receivers
	m.SetHeader("To", recipients...)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or HTML with text/html
	// m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	// m.Attach("/home/Alex/lolcat.jpg") | TODO: implement attachment feature
	switch emailType {
	case "text":
		m.SetBody("text/plain", message)
	case "html":
		m.SetBody("text/html", message)
	}

	// Settings for SMTP server
	d := gomail.NewDialer(mailer.ServerUrl, mailer.Port, mailer.Username, mailer.Password)

	// needed for invalid SSL/TLS certificate | should be set to false in PROD.
	d.TLSConfig = &tls.Config{
		ServerName:         mailer.ServerUrl,
		InsecureSkipVerify: mailer.InsecureSkipVerify,
	}

	// Now send E-Mail
	//fmt.Println("before-email-dial-and-send")
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return mcresponse.GetResMessage("sendmailError", mcresponse.ResponseMessageOptions{
			Message: "Unable to send email message: " + err.Error(),
			Value:   nil,
		})
	}

	// Handle successful email delivery
	return mcresponse.GetResMessage("success", mcresponse.ResponseMessageOptions{
		Message: "Email message successfully sent",
		Value:   nil,
	})
}
