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

	// fmt.Println("Sender Email Address: " + SenderEmailAddress)
	// fmt.Println("Email Host: " + EMAIL_HOST + ":" + os.Getenv("SMTP_PORT"))

	for _, recipientEmailAddress := range recipientEmails {
		message := gomail.NewMessage()
		message.SetHeader("From", SenderEmailAddress)
		message.SetHeader("To", recipientEmailAddress)
		message.SetHeader("Subject", subject)
		message.SetBody("text/html", body)

		dialer := gomail.NewDialer(EMAIL_HOST, SMTP_PORT, SenderEmailAddress, SenderEmailPassword)

		if err := dialer.DialAndSend(message); err != nil {
			fmt.Println("Error when sending the email to: " + recipientEmailAddress + "\nError: ")
			fmt.Println(err)
		}

		fmt.Println("Email has been sent to: " + recipientEmailAddress + "\nfrom: " + SenderEmailAddress)
	}

	return nil
}
