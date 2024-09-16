package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		pcs := api.Group("/pcs")
		{
			pcs.GET("/", h.getAllPCs)
			pcs.POST("/", h.createPC)
			pcs.GET("/:id", h.getPCById)
			pcs.PUT("/:id", h.updatePCById)
			pcs.DELETE("/:id", h.deletePCById)

			items := pcs.Group("/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItemById)
				items.DELETE("/:item_id", h.deleteItemById)
			}
		}
	}

	return router
}
