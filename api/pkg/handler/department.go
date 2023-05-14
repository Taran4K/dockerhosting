package handler

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createDepartment(c *gin.Context) {
	idorg, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	var input models.Department
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	org, err := h.services.Department.Create(input, idorg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, org)
}

func (h *Handler) getAllDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	list, err := h.services.Department.GetAll(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllDepartmentsResponse{
		Data: list,
	})
}

func (h *Handler) getDepartment(c *gin.Context) {
	idorg, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	id, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	org, err := h.services.Department.GetById(id, idorg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, org)
}

func (h *Handler) updateDepartment(c *gin.Context) {
	idorg, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	id, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	_, err = h.services.Department.GetById(id, idorg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.Department
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	org, err := h.services.Department.Update(id, input, idorg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getDepAndmessage{
		Message: "Успешное изменение данных",
		Data:    org,
	})
}

func (h *Handler) getRucovoditel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("department_id"))

	empl, err := h.services.Department.GetRucovoditel(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, empl)
}

func (h *Handler) deleteDepartment(c *gin.Context) {
	idorg, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	id, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	org, err := h.services.Department.GetById(id, idorg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Department.Delete(id, idorg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getDepAndmessage{
		Message: "Успешное удаление данных",
		Data:    org,
	})
}

func (h *Handler) updateRucovoditel(c *gin.Context) {
	oldempl, _ := strconv.Atoi(c.Query("oldempl"))
	newempl, _ := strconv.Atoi(c.Query("newempl"))

	id, _ := strconv.Atoi(c.Param("department_id"))
	empl, err := h.services.Department.GetRucovoditel(id)
	if err != nil {
		empl.ID_Employee = 0
		msg1, err1 := h.services.Department.UpdateRucovoditel(empl, oldempl, newempl)
		if err1 != nil {
			newErrorResponse(c, http.StatusInternalServerError, err1.Error())
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": msg1,
			"error":   err1,
		})
	}

	msg, err := h.services.Department.UpdateRucovoditel(empl, oldempl, newempl)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": msg,
		"error":   err,
	})
}
