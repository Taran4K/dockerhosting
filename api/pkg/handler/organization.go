package handler

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createOrganization(c *gin.Context) {

	var input models.Organization
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	org, err := h.services.Organization.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, org)
}

func (h *Handler) getAllOrganization(c *gin.Context) {
	list, err := h.services.Organization.GetAll()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllOrgsResponse{
		Data: list,
	})
}

func (h *Handler) getOrganization(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	org, err := h.services.Organization.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, org)
}

func (h *Handler) updateOrganization(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	_, err = h.services.Organization.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.Organization
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	org, err := h.services.Organization.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOrgAndmessage{
		Message: "Успешное изменение данных",
		Data:    org,
	})
}

func (h *Handler) deleteOrganization(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	org, err := h.services.Organization.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Organization.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOrgAndmessage{
		Message: "Успешное удаление данных",
		Data:    org,
	})
}

func (h *Handler) getOrganizationByKey(c *gin.Context) {
	key := c.Param("id")

	org, err := h.services.Organization.GetByKey(key)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, org)
}

func (h *Handler) getDirector(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	empl, err := h.services.Organization.GetDirector(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, empl)
}

func (h *Handler) updateDirector(c *gin.Context) {
	oldempl, _ := strconv.Atoi(c.Query("oldempl"))
	newempl, _ := strconv.Atoi(c.Query("newempl"))

	id, _ := strconv.Atoi(c.Param("id"))
	empl, err := h.services.Organization.GetDirector(id)
	if err != nil {
		empl.ID_Employee = 0
		msg, err := h.services.Organization.UpdateDirector(empl, oldempl, newempl)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": msg,
		})
	}

	msg, err := h.services.Organization.UpdateDirector(empl, oldempl, newempl)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": msg,
		"error":   err,
	})
}
