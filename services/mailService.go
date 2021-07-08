package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendMail(toMail string, subject string, body string) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	from := os.Getenv("FROM_MAIL")
	password := os.Getenv("MAIL_PASS")

	to := []string{
		toMail,
	}

	smtpHost := os.Getenv("MAIL_HOST")
	smtpPort := os.Getenv("MAIL_PORT")

	message := []byte("From: Demo<" + from + ">\r\n" +
		"Subject:" + subject + "\r\n" +
		"\r\n" +
		body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
