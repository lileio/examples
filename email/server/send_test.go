package server

import (
	"bytes"
	"io"
	"testing"

	"context"

	"github.com/arbarlow/gomail"
	"github.com/lileio/examples/email"
	"github.com/lileio/examples/email/smtp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var result string

type MockSender struct {
	mock.Mock
}

func (m MockSender) Close() error {
	return nil
}

func (m MockSender) Send(from string, to []string, msg io.WriterTo) error {
	m.Called()

	buf := new(bytes.Buffer)
	_, err := msg.WriteTo(buf)
	if err != nil {
		return err
	}
	s := buf.String()
	result = s

	return nil
}

type MockDialer struct {
	mock.Mock
	Sender MockSender
}

func (m MockDialer) Dial() (gomail.SendCloser, error) {
	m.Called()
	return m.Sender, nil
}

func TestSend(t *testing.T) {
	sender := MockSender{}
	sender.On("Send").Return(nil)

	dialer := MockDialer{Sender: sender}
	dialer.On("Dial").Return(sender, nil)

	smtp.SetDialerAndSender(&dialer, sender)
	go smtp.Start()

	ctx := context.Background()

	req := &email.EmailRequest{
		From:          "someone@gmail.com",
		To:            []string{"alex@lile.io"},
		Subject:       "Welcome to lile",
		PlainText:     "Hi there, welcome",
		HtmlAlternate: "<p>Hi there html peeps</p>",
		Attachments: []*email.Attachment{
			{
				Filename: "file.txt",
				Body:     []byte("Some text file"),
			},
		},
	}

	res, err := s.Send(ctx, req)

	assert.Contains(t, result, "From: someone@gmail.com")
	assert.Contains(t, result, "To: alex@lile.io")
	assert.Contains(t, result, "Subject: Welcome to lile")
	assert.Contains(t, result, "Content-Type: text/plain; charset=UTF-8\r\n\r\nHi there, welcome")
	assert.Contains(t, result, "Content-Type: text/html; charset=UTF-8\r\n\r\n<p>Hi there html peeps</p>")
	assert.Contains(t, result, "Content-Disposition: attachment; filename=\"file.txt\"")

	dialer.AssertExpectations(t)
	sender.AssertExpectations(t)

	assert.Nil(t, err)
	assert.NotNil(t, res)

}
