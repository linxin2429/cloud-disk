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
	fmt.Println(user,password,fmt.Sprintf("%s:%s", host, port))
	err := e.SendWithTLS(fmt.Sprintf("%s:%s", host, port), smtp.PlainAuth("", user, password, host), &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	})
	return err
}

type emailConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     string `json:"port"`
}

func SendEmailFromConfig(configPath, subject, body, mailtype string, to, cc, bcc []string) error {
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
