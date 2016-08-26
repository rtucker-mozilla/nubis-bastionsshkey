package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendWelcomeMail(config Configuration, dest string, message []byte) {
	dialString := fmt.Sprintf("%s:%s", config.AWS.SMTPHostname, config.AWS.SMTPPort)
	fmt.Println(dialString)
	fmt.Println(dest)
	fmt.Println(config.AWS.SMTPFromAddress)
	// Set the sender and recipient.
	auth := smtp.PlainAuth(
		"",
		config.AWS.SMTPUsername,
		config.AWS.SMTPPassword,
		config.AWS.SMTPHostname,
	)
	sendErr := smtp.SendMail(
		dialString,
		auth,
		config.AWS.SMTPFromAddress,
		[]string{dest},
		message,
	)
	if sendErr != nil {
		log.Fatal(sendErr)
	}
}
