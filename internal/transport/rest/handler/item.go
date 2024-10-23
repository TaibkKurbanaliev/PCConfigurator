package handler

import (
	"net/http"
	"pcbuilder/internal/domain"
	"pcbuilder/internal/domain/items"
	"strings"

	"github.com/gin-gonic/gin"
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
	itemType := c.Param("item_type")
	input, err := DecodeByType(c, itemType)

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
	itemType := c.Param("item_type")

	items, err := h.service.PCItem.GetAllItems(c, itemType)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
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
	id := c.Param("item_id")
	itemType := c.Param("item_type")

	input, err := DecodeByType(c, itemType)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	returnID, err := h.service.PCItem.UpdateByID(c, id, input)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": returnID,
	})
}

func (h *Handler) deleteItemById(c *gin.Context) {

}

func DecodeByType(c *gin.Context, itemType string) (interface{}, error) {
	switch strings.ToLower(itemType) {
	case "cpu":
		var item items.CPU
		err := c.BindJSON(&item)
		return item, err
	case "cpu_cooler":
		var item items.CPUCooler
		err := c.BindJSON(&item)
		return item, err
	case "case":
		var item items.Case
		err := c.BindJSON(&item)
		return item, err
	case "gpu":
		var item items.GPU
		err := c.BindJSON(&item)
		return item, err
	case "memory":
		var item items.Memory
		err := c.BindJSON(&item)
		return item, err
	case "motherboard":
		var item items.Motherboard
		err := c.BindJSON(&item)
		return item, err
	case "power_supply":
		var item items.PowerSupply
		err := c.BindJSON(&item)
		return item, err
	case "storage":
		var item items.Storage
		err := c.BindJSON(&item)
		return item, err
	}

	return nil, domain.ErrCollectioNotFound
}
