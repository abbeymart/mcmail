// @Author: abbeymart | Abi Akindele | @Created: 2020-12-25 | @Updated: 2020-12-25
// @Company: mConnect.biz | @License: MIT
// @Description: types

package mcmail

type EmailConfigType struct {
	Username  string
	Password  string
	Port      int
	ServerUrl string
	MsgFrom   string
	ApiKey    string
}

type EmailPropsType map[string]interface{}

type EmailSubjectFunc func(props EmailPropsType) string

type EmailFunc func(props EmailPropsType) interface{}

type TemplateDataType map[string]interface{}

type EmailTemplateType struct {
	Subject EmailSubjectFunc
	Text    EmailFunc
	Html    EmailFunc
}

type MessageObject map[string]string
