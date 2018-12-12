package main

import (
	"software_experiment/pkg/web/router"
)

func main() {
	app := router.GetRouter()

	app.Run("0.0.0.0:8085")
}
