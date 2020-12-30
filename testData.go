// @Author: abbeymart | Abi Akindele | @Created: 2020-12-29 | @Updated: 2020-12-29
// @Company: mConnect.biz | @License: MIT
// @Description: test-data

package mcmail

const EmailUser = "AKIA5G3PE5NSNHPU7F6L"

const EmailPass = "BPdnjRodZlDVDxA8IPY+KI94W2h+AGXvfl8frpqDMl9O"

const EmailServer = "email-smtp.us-west-2.amazonaws.com"

const EmailPort = 587

const TestFromAddress = "abbeya1@yahoo.com"
const TestFromAddress2 = "abiakindele@gmail.com"

const TestToAddress = "abbeya1@yahoo.com"

const ContactInfo = "mconnect.biz Support | support@mconnect.biz"

var amazonaws = EmailConfigType{
	Username:  "AKIA5G3PE5NSNHPU7F6L",
	Password:  "BPdnjRodZlDVDxA8IPY+KI94W2h+AGXvfl8frpqDMl9O",
	Port:      587,
	ServerUrl: "email-smtp.us-west-2.amazonaws.com",
	MsgFrom:   TestFromAddress,
}

var sendinblue = EmailConfigType{
	Username:  "abiakindele@gmail.com",
	Password:  "SvGjR4hUD5xtAQH9",
	Port:      587,
	ServerUrl: "smtp-relay.sendinblue.com",
	MsgFrom:   TestFromAddress2,
	ApiKey:    "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",
}
