package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
	"gopkg.in/yaml.v3"
)

func MergeSlice(s1 []string, s2 []string) []string {
	slice := make([]string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}

func sendToMailHTML(user, password, host, port, subject, body string, to, cc, bcc []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("Cloud Disk<%s>", user)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	e.Subject = subject
	e.HTML = []byte(body)
	err := e.SendWithTLS(fmt.Sprintf("%s:%s", host, port), smtp.PlainAuth("", user, password, host), &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	})
	return NewErrWrapper(err,"sendToMailHTML")
}

type emailConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     string `json:"port"`
}

func SendEmailFromConfig(configPath, subject, body string, to, cc, bcc []string) error {
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		return NewErrWrapper(err, "SendEmailFromConfig")
	}
	var config emailConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return NewErrWrapper(err, "SendEmailFromConfig")
	}
	return sendToMailHTML(config.User, config.Password, config.Host, config.Port, subject, body, to, cc, bcc)
}

func SendEmailCaptcha(email string, code string) error {
	bodyPattern := `
		<html>
        <body>
        <h3>
        %s
        </h3>
        </body>
        </html>
	`
	body := fmt.Sprintf(bodyPattern, code)
	to := []string{email}
	yamlPath := "/home/dengxinlin/golang/cloud_disk/core/etc/email.yaml"
	subject := "[cloud disk captcha]"
	return SendEmailFromConfig(yamlPath, subject, body, to, nil, nil)
}
