package config

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type EmailConfigType struct {
	EmailUsername  string `json:"emailUsername"`
	EmailPassword  string `json:"emailPassword"`
	EmailPort      int    `json:"emailPort"`
	EmailServerUrl string `json:"emailServerUrl"`
	EmailMsgFrom   string `json:"emailMsgFrom"`
	EmailApiKey    string `json:"emailApiKey"`
	AccountName    string `json:"accountName"`
}

type AuditDbType struct {
	DbType                string `json:"dbType"`
	DbHost                string `json:"dbHost"`
	DbUsername            string `json:"dbUsername"`
	DbPassword            string `json:"dbPassword"`
	DbName                string `json:"dbName"`
	DbFilename            string `json:"dbFilename"`
	DbLocation            string `json:"dbLocation"`
	DbPort                uint32 `json:"dbPort"`
	DbPoolSize            uint   `json:"dbPoolSize"`
	DbUrl                 string `json:"dbUrl"`
	DbOptionsSecureAccess bool   `json:"dbOptionsSecureAccess"`
	DbOptionsSecureCert   string `json:"dbOptionsSecureCert"`
	DbOptionsSecureKey    string `json:"dbOptionsSecureKey"`
}

type AppDbType struct {
	DbType                string `json:"dbType"`
	DbHost                string `json:"dbHost"`
	DbUsername            string `json:"dbUsername"`
	DbPassword            string `json:"dbPassword"`
	DbName                string `json:"dbName"`
	DbFilename            string `json:"dbFilename"`
	DbLocation            string `json:"dbLocation"`
	DbPort                uint32 `json:"dbPort"`
	DbPoolSize            uint   `json:"dbPoolSize"`
	DbUrl                 string `json:"dbUrl"`
	DbOptionsSecureAccess bool   `json:"dbOptionsSecureAccess"`
	DbOptionsSecureCert   string `json:"dbOptionsSecureCert"`
	DbOptionsSecureKey    string `json:"dbOptionsSecureKey"`
}

type AppConfigType struct {
	AppMsgFrom         string          `json:"appMsgFrom"`
	AppContact         string          `json:"appContact"`
	InsecureSkipVerify bool            `json:"insecureSkipVerify"`
	ToEmail            []string        `json:"toEmail"`
	EmailService       EmailConfigType `json:"emailService"`
	SendinblueEmail    EmailConfigType `json:"sendinblueEmail"`
	AuditDb            AuditDbType     `json:"auditDb"`
}

const secureFile = "config/config.aes"

var AppConfig AppConfigType

func DecryptConfig(secureFile string, appConfig *AppConfigType) {
	// secure-key-var
	secureKey := ""
	sKeyEnv := os.Getenv("MCAPP_SECURE_KEY")
	if sKeyEnv != "" && len(sKeyEnv) == 32 {
		secureKey = sKeyEnv
	} else {
		log.Fatalln("Valid secureKey environment-variable (MCAPP_SECURE_KEY) is required and must match the encryption secureKey 32-bytes-alphanumeric-value")
	}
	// *secure secureKey, using environment variable or use default at own-risk* OR
	// exclude program-file from GitHub repo-commit and merge privately with the CI/CD/deployment-phase - not necessary
	// secureKey environment-variable (MCAPP_SECURE_KEY) is required and must match the encryption secureKey
	// 32-bytes-alphanumeric-value used to generate the secure config-file(s)
	key := []byte(secureKey)
	// read secure file/content
	cipherData, err := os.ReadFile(secureFile)
	if err != nil {
		log.Fatalln("Could not read secure/cipher file content", err)
	}
	// decrypt secure content
	plaintext, err := decrypt(cipherData, key)
	if err != nil {
		log.Fatalln("Could not decrypt secure input file content - secureKey environment-variable (MCAPP_SECURE_KEY) must match the encryption secureKey 32-bytes-alphanumeric-value ", err)
	}
	// compute app-config value
	err = json.Unmarshal(plaintext, &appConfig)
	if err != nil {
		log.Fatalln("Could not compute app-config value", err)
	}
	fmt.Printf("*****SUCCESS*****\n App-Config value computed:\n\n")
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func LoadConfig() {
	// transform secured/encrypted json-config to AppConfigType type
	DecryptConfig(secureFile, &AppConfig)
}
