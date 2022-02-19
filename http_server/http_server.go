package http_server

import (
	"api/api"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	*fiber.App
}

type Response struct {
	Code        string      `json:"code"`
	Data        interface{} `json:"data"`
	ErrorResult interface{} `json:"errorResult"`
}

func New() api.App {
	return &api.HttpApp{
		App: fiber.New(fiber.Config{
			AppName:               "myPizzaAPI",
			Prefork:               false,
			BodyLimit:             200 * 1024 * 1024,
			DisableStartupMessage: true,
		}),
	}
}
