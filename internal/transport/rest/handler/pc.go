package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createPC(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllPCs(c *gin.Context) {

}

func (h *Handler) getPCById(c *gin.Context) {

}

func (h *Handler) updatePCById(c *gin.Context) {

}

func (h *Handler) deletePCById(c *gin.Context) {

}
