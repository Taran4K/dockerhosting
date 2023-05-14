package handler

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createUser(c *gin.Context) {
	iddep, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	idempl, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	var input models.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	org, err := h.services.User.Create(input, idempl, iddep)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, org)
}

func (h *Handler) getAllUser(c *gin.Context) {
	iddep, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	id, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	list, err := h.services.User.GetAll(id, iddep)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllUsersResponse{
		Data: list,
	})
}

func (h *Handler) getUser(c *gin.Context) {
	iddep, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	idempl, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	org, err := h.services.User.GetById(id, idempl, iddep)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, org)
}

func (h *Handler) updateUser(c *gin.Context) {
	iddep, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	idempl, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	_, err = h.services.User.GetById(id, idempl, iddep)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	org, err := h.services.User.Update(id, input, idempl, iddep)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getUserAndmessage{
		Message: "Успешное изменение данных",
		Data:    org,
	})
}

func (h *Handler) updateProfileUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("iduser"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	userall, err := h.services.Authorization.GetAll(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userall.User.Login = input.Login
	userall.Employee.Email = input.Email

	user, err := h.services.User.UpdateProfile(id, userall)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getUserAndmessage{
		Message: "Успешное изменение данных",
		Data:    user,
	})
}

func (h *Handler) updateProfileEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("iduser"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	_, err = h.services.Authorization.GetAll(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.Employee
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User.UpdateEmployee(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getEmployeeAndmessage{
		Message: "Успешное изменение данных",
		Data:    user,
	})
}

func (h *Handler) updatePasswordUser(c *gin.Context) {
	oldpass := c.Query("oldpassword")
	newpass := c.Query("newpassword")

	id, err := strconv.Atoi(c.Param("iduser"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	userall, err := h.services.Authorization.GetAll(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	msg, err := h.services.User.UpdatePassword(id, oldpass, newpass, userall)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": msg,
	})
}

func (h *Handler) deleteUser(c *gin.Context) {
	iddep, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	idempl, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	org, err := h.services.User.GetById(id, idempl, iddep)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.User.Delete(id, idempl, iddep)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getUserAndmessage{
		Message: "Успешное удаление данных",
		Data:    org,
	})
}

func (h *Handler) createUserTask(c *gin.Context) {
	var input models.Task
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	task, err := h.services.User.CreateTask(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) getAllUserTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("iduser"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	userall, err := h.services.Authorization.GetAll(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, err := h.services.User.GetAllTask(userall.Employee.ID_Employee)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTasksResponse{
		Data: list,
	})
}

func (h *Handler) getUserTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idtask"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	task, err := h.services.User.GetTaskById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) updateUserTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idtask"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	_, err = h.services.User.GetTaskById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.Task
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	task, err := h.services.User.UpdateTask(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getTaskAndmessage{
		Message: "Успешное изменение данных",
		Data:    task,
	})
}

func (h *Handler) deleteUserTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idtask"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	task, err := h.services.User.GetTaskById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.User.DeleteTask(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getTaskAndmessage{
		Message: "Успешное удаление данных",
		Data:    task,
	})
}

func (h *Handler) deleteLogUserTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idtask"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	task, err := h.services.User.GetTaskById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.User.DeleteLogTask(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getTaskAndmessage{
		Message: "Успешное удаление данных",
		Data:    task,
	})
}

func (h *Handler) createUserGoal(c *gin.Context) {
	var input models.Goal
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	goal, err := h.services.User.CreateGoal(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, goal)
}

func (h *Handler) getAllUserGoal(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("iduser"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	userall, err := h.services.Authorization.GetAll(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, err := h.services.User.GetAllGoal(userall.Department.ID_Department)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllGoalsResponse{
		Data: list,
	})
}

func (h *Handler) getUserGoal(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idgoal"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	goal, err := h.services.User.GetGoalById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, goal)
}

func (h *Handler) updateUserGoal(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idgoal"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	_, err = h.services.User.GetGoalById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.Goal
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	goal, err := h.services.User.UpdateGoal(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getGoalAndmessage{
		Message: "Успешное изменение данных",
		Data:    goal,
	})
}

func (h *Handler) deleteUserGoal(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idgoal"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	goal, err := h.services.User.GetGoalById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.User.DeleteGoal(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getGoalAndmessage{
		Message: "Успешное удаление данных",
		Data:    goal,
	})
}

func (h *Handler) deleteLogUserGoal(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idgoal"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	goal, err := h.services.User.GetGoalById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.User.DeleteLogGoal(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getGoalAndmessage{
		Message: "Успешное удаление данных",
		Data:    goal,
	})
}

func (h *Handler) createUserStrat(c *gin.Context) {
	var input models.Strategy
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	strat, err := h.services.User.CreateStrat(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, strat)
}

func (h *Handler) getAllUserStrat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("iduser"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	userall, err := h.services.Authorization.GetAll(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, err := h.services.User.GetAllStrat(userall.Organization.Id_Organization)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllStrategysResponse{
		Data: list,
	})
}

func (h *Handler) getUserStrat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idstrat"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	strat, err := h.services.User.GetStratById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, strat)
}

func (h *Handler) updateUserStrat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idstrat"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	_, err = h.services.User.GetStratById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.Strategy
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	strat, err := h.services.User.UpdateStrat(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getStratAndmessage{
		Message: "Успешное изменение данных",
		Data:    strat,
	})
}

func (h *Handler) deleteUserStrat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idstrat"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	strat, err := h.services.User.GetStratById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.User.DeleteStrat(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getStratAndmessage{
		Message: "Успешное удаление данных",
		Data:    strat,
	})
}

func (h *Handler) deleteLogUserStrat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idstrat"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный ключ")
	}

	strat, err := h.services.User.GetStratById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.User.DeleteLogStrat(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getStratAndmessage{
		Message: "Успешное удаление данных",
		Data:    strat,
	})
}
