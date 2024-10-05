package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pcbuilder/internal/domain/items"
)

// @Summary      PC item Add
// @Description  create item
// @Tags         createItem
// @Accept       json
// @Produce      json
// @Param        input   body      items.Memory  true  "Item"
// @Success      200  {object}
// @Failure      400  {object}  response
// @Failure      404  {object}  response
// @Failure      500  {object}  response
// @Router       /accounts/{id} [post]
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
	id := c.Param("item_id")
	itemType := c.Param("item_type")

	item, err := h.service.PCItem.GetByID(c, id, itemType)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItemById(c *gin.Context) {

}

func (h *Handler) deleteItemById(c *gin.Context) {

}
