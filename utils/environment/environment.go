package environment

import (
	"api/utils"
	"os"
)

func GetPort() string {
	port := os.Getenv("PORT")
	return utils.IfThenElse(len(port) == 0, "7878", port).(string)
}
