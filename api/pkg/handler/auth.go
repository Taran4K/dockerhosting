package handler

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input signInInput
	var user models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user.Login = input.Login
	user.Password = input.Password

	idempl, err := h.services.Authorization.GetEmployee(input.Email)

	user.Employee_ID = idempl

	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type codeVerifyInput struct {
	Code  string `json:"code" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	idempl, err := h.services.Authorization.GetEmployee(input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.services.Authorization.Authorize(input.Login, input.Password, idempl)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if token == "Почта не подходит к данному пользователю" {
		newErrorResponse(c, http.StatusInternalServerError, "Почта не подходит к данному пользователю")
		return
	}

	iduser, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userall, err := h.services.Authorization.GetAll(iduser)

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"data":  userall,
	})
}

func (h *Handler) emailCheck(c *gin.Context) {
	var input codeVerifyInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	check, err := h.services.Authorization.EmailCheck(input.Email, input.Code)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": check,
	})
}
