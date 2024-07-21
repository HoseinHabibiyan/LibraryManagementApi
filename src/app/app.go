package app

import (
	"gorm.io/gorm"
	"lib_management/config"
	"lib_management/database"
)

type App struct {
	db     *gorm.DB
	config *config.Config
}

func Run(cfg *config.Config) {
	app := new(App)
	app.db = database.Connent(cfg)
	app.config = cfg

	//r := gin.Default()
	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//	})
	//})
	//r.Run()
}
