// @Author: abbeymart | Abi Akindele | @Created: 2020-12-25 | @Updated: 2020-12-25
// @Company: mConnect.biz | @License: MIT
// @Description: types

package mcmail

type EmailConfigType struct {
	Username           string `json:"username"`
	Password           string `json:"password"`
	Port               int    `json:"port"`
	ServerUrl          string `json:"serverUrl"`
	MsgFrom            string `json:"msgFrom"`
	ApiKey             string `json:"apiKey"`
	Tls                string `json:"tls"`
	InsecureSkipVerify bool   `json:"insecureSkipVerify"`
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
