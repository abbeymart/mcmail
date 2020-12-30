// @Author: abbeymart | Abi Akindele | @Created: 2020-12-29 | @Updated: 2020-12-29
// @Company: mConnect.biz | @License: MIT
// @Description: send-mail function

package mcmail

import (
	"github.com/abbeymart/mcresponse"
	"net/smtp"
)

// SendEmail method sends simple email messages to multiple recipients
// The message parameter should be an RFC 822-style email with headers first, a blank line, and then the message body.
// The lines of msg should be CRLF terminated.
// The msg headers should usually include fields such as "From", "To", "Subject", and "Cc".
// Sending "Bcc" messages is accomplished by including an email address in the to parameter
// but not including it in the msg headers.
// The SendMail function and the net/smtp package are low-level mechanisms and provide no support for DKIM signing,
// MIME attachments (see the mime/multipart package), or other mail functionality.
// Higher-level packages exist outside of the standard library
func (mailer EmailConfigType) SendEmail(recipients []string, message string) mcresponse.ResponseMessage {
	// TODO: validate toEmailAddress as slice of strings

	// Authenticate email server
	auth := smtp.PlainAuth("", mailer.Username, mailer.Password, mailer.ServerUrl)

	// Send email to toAddress(es)
	err := smtp.SendMail(mailer.ServerUrl+":"+mailer.Port, auth, mailer.MsgFrom, recipients, []byte(message))

	// Handle email error, if any
	if err != nil {
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
