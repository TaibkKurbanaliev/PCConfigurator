package handler

import (
	"net/http"
	"pcbuilder/internal/domain"

	"github.com/gin-gonic/gin"
)

// @Summary User SignUp
// @Tags users-auth
// @Description create user account
// @ModuleID SignUp
// @Accept  json
// @Produce  json
// @Param input body domain.User true "sign up info"
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input domain.User

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(c.Request.Context(), input)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"name" binding:"required" bson:"name"`
	Password string `json:"password" binding:"required" bson:"password"`
}

// @Summary User SignIn
// @Tags users-auth
// @Description user sign in
// @ModuleID SignIn
// @Accept  json
// @Produce  json
// @Param input body signInInput true "sign up info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(c.Request.Context(), input.Username, input.Password)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
