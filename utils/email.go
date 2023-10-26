package utils

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(recipientEmails []string, subject string, body string) error {
	SenderEmailAddress := os.Getenv("SENDER_EMAIL_ADDRESS")
	SenderEmailPassword := os.Getenv("SENDER_EMAIL_PASSWORD")
	SMTP_PORT, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	EMAIL_HOST := os.Getenv("EMAIL_HOST")

	for _, recipientEmailAddress := range recipientEmails {
		message := gomail.NewMessage()
		message.SetHeader("From", SenderEmailAddress)
		message.SetHeader("To", recipientEmailAddress)
		message.SetHeader("Subject", subject)
		message.SetBody("text/html", body)

		dialer := gomail.NewDialer(EMAIL_HOST, SMTP_PORT, SenderEmailAddress, SenderEmailPassword)

		if err := dialer.DialAndSend(message); err != nil {
			fmt.Println(err)
		}
	}

	return nil
}
