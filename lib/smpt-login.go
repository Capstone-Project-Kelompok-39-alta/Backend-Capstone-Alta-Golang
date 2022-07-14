package lib

import (
	"errors"
	"net/smtp"
)

type loginAuth struct {
	username, password string
}

func (l *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (l *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username":
			return []byte(l.username), nil
		case "Password":
			return []byte(l.password), nil
		default:
			return nil, errors.New("unknown from server")

		}
	}
	return nil, nil
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}
