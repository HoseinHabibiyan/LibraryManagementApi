package controller

import (
	"github.com/gin-gonic/gin"
	"lib_management/container"
	"lib_management/domain"
	"net/http"
)

type AuthorController struct {
	container container.Container[domain.Author]
}

func NewAuthorController(c container.Container[domain.Author]) *AuthorController {
	return &AuthorController{
		container: c,
	}
}

func (c *AuthorController) GetAll(context *gin.Context) {
	repo := c.container.GetRepository()
	data := repo.GetAll()
	context.JSON(http.StatusOK, data)
}
