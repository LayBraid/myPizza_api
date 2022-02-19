package main

import (
	"api/http_server"
	"api/http_server/controllers"
	log "github.com/sirupsen/logrus"
	log2 "log"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		ForceColors:     true,
		TimestampFormat: "02-01-2006 15:04:05",
	})

	log.SetLevel(log.DebugLevel)
}

func main() {
	app := http_server.New()
	app.RegisterControllers(controllers.PizzaControllers())

	log2.Fatal(app.Start())
}
