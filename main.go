package main

import (
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/server"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	app := server.Server()
	err := app.Start(":8080")

	app.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	if err != nil {
		panic(err)
	}

	app.Logger.Fatal(app.StartAutoTLS(":8080"))
}
