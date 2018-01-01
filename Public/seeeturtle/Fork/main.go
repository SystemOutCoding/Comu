package main

import (
	"github.com/seeeturtle/Fork/app"
	"github.com/seeeturtle/Fork/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
