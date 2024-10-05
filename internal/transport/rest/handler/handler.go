package handler

import (
	"pcbuilder/internal/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "pcbuilder/docs"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
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
				items.GET("/:item_type/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItemById)
				items.DELETE("/:item_id", h.deleteItemById)
			}
		}
	}

	return router
}
