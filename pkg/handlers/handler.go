package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/goIdioms/rest/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-in", h.singIn)
		auth.POST("/sing-up", h.singUp)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAllLists)
			lists.POST("/", h.createList)
			lists.PUT("/:id", h.updateList)
			lists.GET("/:id", h.getListById)
			lists.DELETE("/:id", h.deleteList)

		}

		items := api.Group("/:id/items")
		{
			items.GET("/", h.getAllItems)
			items.POST("/", h.createItem)
			items.PUT("/:item_id", h.updateItem)
			items.GET("/:item_id", h.getItemById)
			items.DELETE("/:item_id", h.deleteItem)
		}
	}

	return router
}
