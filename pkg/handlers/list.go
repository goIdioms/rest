package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goIdioms/rest/pkg/models"
)

func (h Handler) createList(c *gin.Context) {
	userID, err := GetUserId(c)
	if err != nil {
		return
	}

	var input models.TodoList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userID, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "fail to create list")
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h Handler) getAllLists(c *gin.Context) {

}

func (h Handler) getListById(c *gin.Context) {

}
func (h Handler) updateList(c *gin.Context) {

}

func (h Handler) deleteList(c *gin.Context) {

}
