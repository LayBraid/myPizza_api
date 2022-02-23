package api

import (
	"github.com/gofiber/fiber/v2"
)

type App interface {
	Start() error
	RegisterRoutes(routes ...Route)
	RegisterControllers(controllers ...Controller)
}

type HttpApp struct {
	*fiber.App
}

type Route struct {
	Path    string
	Method  Method
	Handler func(c *fiber.Ctx) error
}

type Controller struct {
	BasePath string
	Routes   []Route
}
