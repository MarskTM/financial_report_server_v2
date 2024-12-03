package service

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"phenikaa/infrastructure"

	"gorm.io/gorm"

	"github.com/jordan-wright/email"
)

type EmailService interface {
	SendEmail(recipients []string, subject, content string) error
}

type emailService struct {
	db *gorm.DB
}

func (s *emailService) SendEmail(recipients []string, subject, content string) error {
	mailServerHost, mailPort, mailAccount, mailPass := infrastructure.GetMailParam()
	doneChannel := make(chan interface{})

	mail := email.NewEmail()
	mail.From = mailAccount
	mail.To = recipients
	mail.Subject = subject
	mail.HTML = []byte(content)

	go func(c chan interface{}) {
		err := mail.Send(mailServerHost+":"+mailPort, smtp.PlainAuth("", mailAccount, mailPass, mailServerHost))
		if err != nil {
			log.Println(err)
			c <- err
		} else {
			c <- "Send email successfully."
		}
	}(doneChannel)

	done := <-doneChannel
	switch v := done.(type) {
	default:
		errorStr := fmt.Sprintf("Unexpected type %T", v)
		return errors.New("can't send email. " + errorStr)
	case error:
		return v
	case string:
		return nil
	}
}

// Việc trả về luôn dịa chỉ của 
func NewEmailService() *emailService {
	db := infrastructure.GetDB()
	return &emailService{
		db: db,
	}
}
