package mailer

import (
	"errors"

	gomail "gopkg.in/mail.v2"
)

type mailTrapClient struct {
	fromEmail string
	apiKey    string
}

func NewMailTrapClient(apiKey, fromEmail string) (mailTrapClient, error) {
	if apiKey == "" {
		return mailTrapClient{}, errors.New("mailtrap api key is required")
	}

	return mailTrapClient{
		apiKey:    apiKey,
		fromEmail: fromEmail,
	}, nil

}

func (m mailTrapClient) Send(templateFile, username, email string, data any, isSandBox bool) (int, error) {
	// template parsing

	template, err := TemplateParsing(templateFile, data)
	if err != nil {
		return -1, err
	}

	message := gomail.NewMessage()

	message.SetHeader("From", m.fromEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", template.subject.String())

	message.AddAlternative("text/html", template.body.String())

	// setup SMTP dialer
	dialer := gomail.NewDialer("live.smtp.mailtrap.io", 587, "api", m.apiKey)

	// send and retries
	statusCode, err := Retries(MaxRetries, func() (any, error) {
		return 200, dialer.DialAndSend(message)
	})

	return statusCode, err
}
