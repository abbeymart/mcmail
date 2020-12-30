// @Author: abbeymart | Abi Akindele | @Created: 2020-12-29 | @Updated: 2020-12-29
// @Company: mConnect.biz | @License: MIT
// @Description: send-mail function

package mcmail

import (
	"crypto/tls"
	"fmt"
	"github.com/abbeymart/mcresponse"
	"gopkg.in/gomail.v2"
	"net/smtp"
)

// SendTextEmail method sends simple email messages to multiple recipients
// The message parameter should be an RFC 822-style email with headers first, a blank line, and then the message body.
// The lines of msg should be CRLF terminated.
// The msg headers should usually include fields such as "From", "To", "Subject", and "Cc".
// Sending "Bcc" messages is accomplished by including an email address in the to parameter
// but not including it in the msg headers.
// The SendMail function and the net/smtp package are low-level mechanisms and provide no support for DKIM signing,
// MIME attachments (see the mime/multipart package), or other mail functionality.
// Higher-level packages exist outside of the standard library
func (mailer EmailConfigType) SendTextEmail(recipients []string, message string) mcresponse.ResponseMessage {
	// Authenticate email server
	auth := smtp.PlainAuth("", mailer.Username, mailer.Password, mailer.ServerUrl)

	fmt.Printf("auth-response: %v\n\n", auth)
	// Send email to toAddress(es)
	err := smtp.SendMail(mailer.ServerUrl+":"+fmt.Sprintf("%v", mailer.Port), auth, mailer.MsgFrom, recipients, []byte(message))

	// Handle email error, if any
	if err != nil {
		fmt.Printf("Mail error: %v", err.Error())
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

// SendEmail sends text and html messages, attachment etc.
func (mailer EmailConfigType) SendEmail(recipients []string, message string, subject string, emailType string) mcresponse.ResponseMessage {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", mailer.MsgFrom)

	// Set E-Mail receivers
	m.SetHeader("To", recipients...)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
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
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
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
