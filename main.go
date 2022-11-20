package main

import (
	"github.com/jaimera/poc-services/application"
	"github.com/jaimera/poc-services/server"
)

func main() {
	app := application.BuildApp()

	server.Run(app)
}
