package handler

import (
	"net/http"
	"pcbuilder/internal/domain/items"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	var input items.Memory
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.PCItem.Create(c, input)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {

}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItemById(c *gin.Context) {

}

func (h *Handler) deleteItemById(c *gin.Context) {

}
