package app

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"lib_management/config"
	"lib_management/database"
	"lib_management/docs"
	"log"
)

func Run(cfg *config.Config) {

	db := database.Connent(cfg)

	engine := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	MapRouter(engine, db)

	err := engine.Run()

	if err != nil {
		log.Fatal(err)
	}
}
