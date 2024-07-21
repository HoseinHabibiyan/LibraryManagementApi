package main

import (
	"lib_management/app"
	"lib_management/config"
)

func main() {

	cfg := config.GetConfig()
	app.Run(cfg)
}
