package main

import (
	"./internal/application"
)

func main() {
	app := application.New()
	app.RunServer()
}
