package main

import "Calculation/internal/application"

func main() {
	app := application.New()
	app.RunServer()
}
