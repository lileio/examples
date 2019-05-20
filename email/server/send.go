package server

import (
	"bytes"
	"context"

	"github.com/arbarlow/gomail"
	"github.com/lileio/examples/email"
	"github.com/lileio/examples/email/smtp"
)

func (s EmailServer) Send(ctx context.Context, r *email.EmailRequest) (*email.EmailResponse, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", r.From)
	m.SetHeader("To", r.To...)
	m.SetHeader("Subject", r.Subject)
	m.SetBody("text/plain", r.PlainText)

	if r.HtmlAlternate != "" {
		m.AddAlternative("text/html", r.HtmlAlternate)
	}

	for _, a := range r.Attachments {
		m.AttachReader(a.Filename, bytes.NewReader(a.Body))
	}

	for k, v := range r.Headers {
		m.SetHeader(k, v)
	}

	err := smtp.SendMessage(m)
	if err != nil {
		return nil, err
	}

	return &email.EmailResponse{}, nil
}
