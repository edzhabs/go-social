package mailer

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"
	"time"
)

const (
	FromName            = "GopherSocial"
	MaxRetries          = 3
	UserWelcomeTemplate = "user_invitation.tmpl"
)

//go:embed "templates"
var FS embed.FS

type Client interface {
	Send(templateFile, username, email string, data any, isSandbox bool) (int, error)
}

type templateType struct {
	subject *bytes.Buffer
	body    *bytes.Buffer
}

func TemplateParsing(templateFile string, data any) (templateType, error) {
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return templateType{}, err
	}

	template := templateType{
		subject: new(bytes.Buffer),
		body:    new(bytes.Buffer),
	}

	if err := tmpl.ExecuteTemplate(template.subject, "subject", data); err != nil {
		return templateType{}, err
	}

	if err := tmpl.ExecuteTemplate(template.body, "body", data); err != nil {
		return templateType{}, err
	}
	return template, nil
}

func Retries(maxRetries int, fn func() (any, error)) (int, error) {
	var retryErr error

	for i := range maxRetries {
		_, retryErr = fn()
		if retryErr != nil {
			// exponential backoff
			time.Sleep(time.Second * time.Duration(i))
			continue
		}

		return 200, nil
	}

	return -1, fmt.Errorf("failed to sent email after %d attempts, err: %v", MaxRetries, retryErr)
}
