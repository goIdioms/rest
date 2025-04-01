package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goIdioms/rest/pkg/models"
)

func (h Handler) singUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h Handler) singIn(c *gin.Context) {

}
