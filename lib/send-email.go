package lib

import (
	"fmt"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"net/smtp"
)

func SendEmail(message entities.SendCustomer) error {
	server := "smtp.office365.com"
	port := 587
	user := "invoinesia-cs@outlook.com"
	from := user
	pass := "Mathilda234"
	dest := message.To

	auth := LoginAuth(user, pass)

	to := []string{dest}

	messages := []byte("From: " + from + "\n" +
		"To: " + dest + "\n" +
		"Subject: " + message.Subject + "\n\n" +
		message.Body)

	endpoint := fmt.Sprintf("%v:%v", server, port)

	err := smtp.SendMail(endpoint, auth, from, to, messages)
	if err != nil {
		fmt.Println("this is error guys :", err)
	}
	return nil
}
