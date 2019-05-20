package smtp

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/arbarlow/gomail"
	log "github.com/sirupsen/logrus"
)

type Dialer interface {
	Dial() (gomail.SendCloser, error)
}

type email struct {
	message  *gomail.Message
	response chan response
}

type response struct {
	error error
}

var (
	CloseTime = time.Duration(30)

	sender gomail.SendCloser
	dialer Dialer
	port   int

	queue = make(chan email)
	open  = false
)

func init() {
	if os.Getenv("SMTP_HOST") == "" {
		os.Setenv("SMTP_HOST", "localhost")
	}

	if os.Getenv("SMTP_PORT") == "" {
		os.Setenv("SMTP_PORT", "25")
	}

	var err error
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatal("Could not parse SMTP_PORT into an integer")
	}

	dialer = gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port,
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"))
}

func Start() {
	for {
		select {
		case m := <-queue:
			err := connect()
			if err != nil {
				log.Errorf("SMTP connection error: %v", err)

				m.response <- response{error: err}
				break
			}

			t := time.Now()
			if err := gomail.Send(sender, m.message); err != nil {
				log.Errorf("Sending email failed: %v", err)

				// Close the connection, so that the next send can reconnect
				if strings.Contains(err.Error(), "timed out") {
					closeConn()
				}

				m.response <- response{error: err}
				break
			}
			log.Debugf("Sent email in %s", time.Since(t))

			m.response <- response{}
		case <-time.After(CloseTime * time.Second):
			closeConn()
		}
	}
}

func SendMessage(m *gomail.Message) error {
	res := make(chan response)

	queue <- email{
		message:  m,
		response: res,
	}

	r := <-res
	return r.error
}

func SetDialerAndSender(d Dialer, s gomail.SendCloser) {
	dialer = d
	sender = s
}

func connect() error {
	if !open {
		var err error
		if sender, err = dialer.Dial(); err != nil {
			return err
		}

		open = true
	}
	return nil
}

func closeConn() {
	if open {
		log.Debug("Closing STMP connection")
		if err := sender.Close(); err != nil {
			if !strings.Contains(err.Error(), "broken pipe") {
				log.Error(err)
			}
		}
		open = false
	}
}
