package sendgrid

import (
	"github.com/sendgrid/sendgrid-go"
	email "github.com/sendgrid/sendgrid-go/helpers/mail"
	"liar/pkg/mail"
	"log"
)

type SendGrid struct {
	ApiKey string
}

func (s *SendGrid) Send(msg mail.Message) error {
	m := email.NewV3Mail()

	from := email.NewEmail(msg.SenderName, msg.SenderEmail)
	m.SetFrom(from)

	m.Subject = msg.Subject

	p := email.NewPersonalization()
	tos := make([]*email.Email, len(msg.Receiver))

	for _, e := range msg.Receiver {
		tos = append(tos, email.NewEmail(e, e))
	}

	p.AddTos(tos...)
	m.AddPersonalizations(p)

	content := email.NewContent("text/html", msg.Content)
	m.AddContent(content)

	client := sendgrid.NewSendClient(s.ApiKey)

	response, err := client.Send(m)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}

	return err
}
