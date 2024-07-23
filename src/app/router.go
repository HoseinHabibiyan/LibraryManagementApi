package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lib_management/container"
	"lib_management/controller"
	"lib_management/domain"
	"lib_management/repository"
)

func MapRouter(g *gin.Engine, db *gorm.DB) {
	router := g.Group("api")

	author := repository.GetRepository[domain.Author](db)
	c := container.NewContainer(author)

	// author
	authorController := controller.NewAuthorController(c)
	router.GET("/", authorController.GetAll)
}
