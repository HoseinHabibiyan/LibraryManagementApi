package app

import (
	"github.com/gin-gonic/gin"
	"lib_management/config"
	"lib_management/database"
	"log"
)

func Run(cfg *config.Config) {

	db := database.Connent(cfg)

	engine := gin.Default()

	MapRouter(engine, db)

	err := engine.Run()

	if err != nil {
		log.Fatal(err)
	}
}
