package main

import (
	"liar/pkg/config"
	"liar/pkg/mail"
	"liar/pkg/sendgrid"
	"log"
)

func main() {
	var cf config.Conf
	cf.LoadConf()

	var sender mail.Sender

	sender = sendgrid.NewSendGrid(cf.SendGrid.Key)

	err := sender.Send(&mail.Message{
		Subject:     "Demo api",
		SenderEmail: cf.Sender.Mail,
		SenderName:  cf.Sender.Name,
		Receiver:    []string{"haind.vng@gmail.com"},
		Content:     "Say hi",
	})

	log.Fatal(err)
}
