package utils

import (
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(id int, url string) error {
	// load env vars
	host := os.Getenv("MAS_EMAIL_HOST")
	//port := os.Getenv("MAS_EMAIL_PORT")
	user  := os.Getenv("MAS_EMAIL_USER")
	password := os.Getenv("MAS_EMAIL_PASSWORD")
	to := os.Getenv("MAS_EMAIL_TO")

	// set email sender
	mail := gomail.NewMessage()
	mail.SetHeader("From", user)
	mail.SetHeader("To", to)

	mail.SetHeader("Subject", "Hello!")
	mailBody := "<h1> Upsee Service alert </h1> <p>" + url + " is down!"
	mail.SetBody("text/html", mailBody)

	dialer := gomail.NewDialer(host, 465, user, password)

	if err := dialer.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}