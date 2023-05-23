package mail

import (
	"bytes"
	"fmt"
	"github.com/mdcaceres/doctest/models/dto"
	"net/smtp"
	"os"
	"text/template"
)

var (
	simplePath     = "C:\\Users\\mdcac\\go\\src\\github.com\\mdcaceres\\doctest\\services\\mail\\templates\\ActionTemplate.html"
	invitationPath = "C:\\Users\\mdcac\\go\\src\\github.com\\mdcaceres\\doctest\\services\\mail\\templates\\Invitation.html"
)

type IMailService interface {
	SendSimple(payload *dto.MailData) error
}

type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (m *EmailService) SendSimple(payload *dto.MailData) error {
	from := os.Getenv("SMTP_USER")
	addr := os.Getenv("SMTP_ADDR")
	auth := getAuth()
	to := payload.To
	body, err := getSimpleBody(payload)
	if err != nil {
		return err
	}
	msg := getMsg(payload.Subject, body)
	err = smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		return err
	}
	return nil
}

func (m *EmailService) SendInvitation(payload *dto.InviteMailData) error {
	from := os.Getenv("SMTP_USER")
	addr := os.Getenv("SMTP_ADDR")
	auth := getAuth()
	to := payload.To
	body, err := getInvitationBody(payload)
	if err != nil {
		return err
	}
	msg := getMsg(payload.Subject, body)
	err = smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		return err
	}
	return nil
}

func getSimpleBody(payload *dto.MailData) (string, error) {
	var body bytes.Buffer
	t, err := template.ParseFiles(simplePath)
	if err != nil {
		return "", err
	}
	err = t.Execute(&body, struct {
		Name   string
		Action string
		Url    string
	}{Name: payload.Name, Action: payload.Action, Url: payload.Url})
	if err != nil {
		return "", err
	}
	return body.String(), nil
}

func getInvitationBody(payload *dto.InviteMailData) (string, error) {
	var body bytes.Buffer
	t, err := template.ParseFiles(invitationPath)
	if err != nil {
		return "", err
	}
	err = t.Execute(&body, struct {
		Name        string
		InvitedId   string
		ProjectId   string
		ProjectName string
		Url         string
	}{Name: payload.Name, InvitedId: payload.InvitedId, ProjectId: payload.ProjectId, ProjectName: payload.ProjectName, Url: payload.Url})
	if err != nil {
		return "", err
	}
	return body.String(), nil
}

func getMsg(subject string, body string) []byte {
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	return []byte(fmt.Sprintf("Subject:%s\n%s\n\n%s", subject, headers, body))
}

func getAuth() smtp.Auth {
	host := os.Getenv("SMTP_HOST")
	user := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")
	auth := smtp.PlainAuth(
		"",
		user,
		password,
		host)
	return auth
}
