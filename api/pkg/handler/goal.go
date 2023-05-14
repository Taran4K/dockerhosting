package handler

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createGoal(c *gin.Context) {
	idorg, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	iddep, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	var input models.Goal
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	goal, err := h.services.Goal.Create(input, iddep, idorg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	idempl, err3 := h.services.Department.GetRucovoditel(iddep)
	if err3 != nil && idempl.ID_Employee != 0 {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if idempl.EmailVerified == true {
		dep, err3 := h.services.Department.GetById(iddep, idorg)
		if err3 != nil && idempl.ID_Employee != 0 {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		var message = "Создана цель: " + goal.Name + ", для отдела: " + dep.Name + ", которую необходимо выполнить до: " + goal.Date_end[:len(goal.Date_end)-10] + ". Описание цели: " + goal.Description
		_, err2 := h.services.User.SendMail(idempl.Email, "Добавлена цель", message)
		if err2 != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, goal)
}

func (h *Handler) getAllGoal(c *gin.Context) {
	idorg, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	id, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	list, err := h.services.Goal.GetAll(id, idorg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllGoalsResponse{
		Data: list,
	})
}

func (h *Handler) getGoal(c *gin.Context) {
	idorg, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	iddep, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	id, err := strconv.Atoi(c.Param("goal_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	org, err := h.services.Goal.GetById(id, iddep, idorg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, org)
}

func (h *Handler) updateGoal(c *gin.Context) {
	idorg, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	iddep, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	id, err := strconv.Atoi(c.Param("goal_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	_, err = h.services.Goal.GetById(id, iddep, idorg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.Goal
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	goal, err := h.services.Goal.Update(id, input, iddep, idorg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	idempl, err3 := h.services.Department.GetRucovoditel(iddep)
	if err3 != nil && idempl.ID_Employee != 0 {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if idempl.EmailVerified == true {
		dep, err3 := h.services.Department.GetById(iddep, idorg)
		if err3 != nil && idempl.ID_Employee != 0 {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		var message = "Обновлена цель: " + goal.Name + ", для отдела: " + dep.Name + ", которую необходимо выполнить до: " + goal.Date_end[:len(goal.Date_end)-10] + ". Описание цели: " + goal.Description
		_, err2 := h.services.User.SendMail(idempl.Email, "Цель обновлена", message)
		if err2 != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, getGoalAndmessage{
		Message: "Успешное изменение данных",
		Data:    goal,
	})
}

func (h *Handler) deleteGoal(c *gin.Context) {
	idorg, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	iddep, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	id, err := strconv.Atoi(c.Param("goal_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	org, err := h.services.Goal.GetById(id, iddep, idorg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Goal.Delete(id, iddep, idorg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getGoalAndmessage{
		Message: "Успешное удаление данных",
		Data:    org,
	})
}
