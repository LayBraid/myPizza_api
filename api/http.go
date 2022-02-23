package api

import (
	"api/utils/environment"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Method int

const (
	GET Method = iota
	POST
	DELETE
	PUT
	ALL
)

func (h *HttpApp) Start() error {
	log.Println("Is start :)")
	return h.Listen(":" + environment.GetPort())
}

func (h *HttpApp) RegisterControllers(controllers ...Controller) {
	for _, controller := range controllers {
		for _, route := range controller.Routes {
			route.Path = controller.BasePath + route.Path
		}
		h.RegisterRoutes(controller.Routes...)
	}
}

func (m Method) ToFunction(app *HttpApp) interface{} {
	return []interface{}{app.Get, app.Post, app.Delete, app.Put, app.All}[m]
}

func (h *HttpApp) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		route.Method.ToFunction(h).(func(path string, handlers ...fiber.Handler) fiber.Router)(route.Path, route.Handler)
	}
}
