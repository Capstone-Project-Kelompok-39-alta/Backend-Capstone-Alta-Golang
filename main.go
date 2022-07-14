package main

import (
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/server"
)

func main() {
	app := server.Server()
	err := app.Start(":443")

	if err != nil {
		panic(err)
	}
}
