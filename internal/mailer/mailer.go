package mailer

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"log"
	"sync"
	"time"

	"github.com/go-mail/mail/v2"
)

//go:embed templates
var mailTemplateFS embed.FS

type Mailer struct {
	Dialer   *mail.Dialer
	Sender   string
	MsgChan  chan *MailMsgPayload
	ErrChan  chan error
	DoneChan chan bool
	WG       *sync.WaitGroup
}

type MailMsgPayload struct {
	Recipient string
	Template  string
	Data      any
}

func NewMailMsgPayload() *MailMsgPayload {
	return &MailMsgPayload{}
}

// New Creates a new mailer

func New(host string, port int, username, password, sender string) Mailer {
	dialer := mail.NewDialer(host, port, username, password)
	dialer.Timeout = time.Second * 5

	return Mailer{
		Dialer:   dialer,
		Sender:   sender,
		MsgChan:  make(chan *MailMsgPayload),
		ErrChan:  make(chan error),
		DoneChan: make(chan bool),
	}
}

func (m Mailer) ListenForIncomingEmail() {
	fmt.Println("Listening for mail")
	for {
		select {
		case msg := <-m.MsgChan:
			fmt.Println("hit the route")
			m.SendEmail(msg.Recipient, msg.Template, msg.Data)
		case err := <-m.ErrChan:
			log.Println(err.Error())
		case <-m.DoneChan:
			close(m.MsgChan)
			close(m.ErrChan)
			close(m.DoneChan)
			return
		}
	}

}

func (m Mailer) SendEmail(recipient, templateFile string, data any) error {
	// parse the template
	tmpl, err := template.New("email").ParseFS(mailTemplateFS, "templates/"+templateFile)
	if err != nil {
		return err
	}

	subject := &bytes.Buffer{}

	err = tmpl.ExecuteTemplate(subject, "subject", data)

	if err != nil {
		return err
	}

	plainBody := &bytes.Buffer{}

	err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)

	if err != nil {
		return err
	}

	htmlBody := &bytes.Buffer{}

	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)

	if err != nil {
		return err
	}

	msg := mail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", m.Sender)
	msg.SetHeader("Subject", subject.String())
	msg.SetBody("text/plain", plainBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	err = m.Dialer.DialAndSend(msg)
	if err != nil {
		return err
	}

	m.DoneChan <- true

	return nil

}
