package controllers

import (
	"api/api"
	"api/http_server/routes"
)

func PizzaControllers() api.Controller {
	return api.Controller{
		BasePath: "/pizzas",
		Routes: []api.Route{
			routes.GetPizzas(),
		},
	}
}
