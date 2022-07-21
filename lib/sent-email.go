package lib

import (
	"fmt"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"net/smtp"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, fmt.Errorf("unknown from server")
		}
	}
	return nil, nil
}

func SendEmail(message entities.SendCustomer) error {
	server := "smtp.office365.com"
	port := 587
	username := "invoinesia-cs@outlook.com"
	from := username
	password := "Mathilda234"
	dest := message.To

	auth := LoginAuth(username, password)

	to := []string{dest}

	messages := []byte("From: " + from + "\n" +
		"To: " + dest + "\n" +
		"Subject: " + message.Subject + "\n\n" +
		message.Body)

	endpoint := fmt.Sprintf("%v:%v", server, port)

	err := smtp.SendMail(endpoint, auth, from, to, messages)
	if err != nil {
		fmt.Println("error from send email : ", err)
	}

	return nil
}
