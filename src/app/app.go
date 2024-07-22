package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lib_management/config"
	"lib_management/container"
	"lib_management/controller"
	"lib_management/database"
	"lib_management/domain"
	"lib_management/repository"
)

type App struct {
	db     *gorm.DB
	config *config.Config
}

func Run(cfg *config.Config) {
	app := new(App)
	app.db = database.Connent(cfg)
	app.config = cfg

	author := repository.GetRepository[domain.Author](app.db)
	c := container.GetContainer(author)

	r := gin.Default()

	router := r.Group("api")
	authorController := controller.NewAuthorController(c)
	router.GET("/", authorController.GetAll)

	r.Run()
}
