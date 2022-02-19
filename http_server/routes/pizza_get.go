package routes

import (
	"api/api"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"os"
)

type MyList []Pizza

type Pizza struct {
	Name        string  `json:"name"`
	Ingredients string  `json:"ingredients"`
	Price       float64 `json:"price"`
}

func GetPizzas() *api.Route {
	return &api.Route{
		Path:   "/get_all",
		Method: api.GET,
		Handler: func(c *fiber.Ctx) (err error) {
			open, err := os.Open("./assets/pizzas.json")
			if err != nil {
				return
			}
			all, err := ioutil.ReadAll(open)
			if err != nil {
				return
			}
			c.Set("Content_Type", "application/json")
			return c.Send(all)
		},
	}
}
