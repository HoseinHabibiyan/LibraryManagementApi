package controller

import (
	"github.com/gin-gonic/gin"
	"lib_management/container"
	"lib_management/domain"
	"net/http"
	"strconv"
)

type AuthorController struct {
	container container.Container[domain.Author]
}

func NewAuthorController(c container.Container[domain.Author]) *AuthorController {
	return &AuthorController{
		container: c,
	}
}

// @Schemes
// @Tags Author
// @Accept json
// @Produce json
// @Router /author [get]
func (c *AuthorController) GetAll(context *gin.Context) {
	repo := c.container.GetRepository()
	data := repo.GetAll()
	context.JSON(http.StatusOK, data)
}

// @Schemes
// @Tags Author
// @Accept json
// @Produce json
// @param id path int true "id"
// @Router /author/{id} [get]
func (c *AuthorController) GetById(context *gin.Context) {
	repo := c.container.GetRepository()
	id, _ := strconv.Atoi(context.Param("id"))
	data := repo.GetById(int32(id))
	context.JSON(http.StatusOK, data)
}

// @Schemes
// @Tags Author
// @Accept json
// @Produce json
// @Param author body domain.Author true "author"
// @Router /author/add [post]
func (c *AuthorController) Add(context *gin.Context) {
	repo := c.container.GetRepository()
	var author domain.Author
	err := context.Bind(&author)
	if err != nil {
		return
	}
	err = repo.Create(&author)
	if err != nil {
		panic("add author error!")
	}
}
