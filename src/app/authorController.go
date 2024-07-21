package app

import (
	"github.com/gin-gonic/gin"
	"lib_management/domain"
	"lib_management/repository"
	"net/http"
)

func (app *App) Get(c *gin.Context) {
	repo := repository.NewRepository[domain.Author](app.db)

	data := repo.Get()

	c.JSON(http.StatusOK, data)
}
