package handler

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, error{message})
}

type getAllOperationsResponse struct {
	Data []models.Finances_Operations `json:"data"`
}

type getOperAndmessage struct {
	Message string                     `json:"message"`
	Data    models.Finances_Operations `json:"data"`
}

type getAllDepartmentsResponse struct {
	Data []models.Department `json:"data"`
}

type getDepAndmessage struct {
	Message string            `json:"message"`
	Data    models.Department `json:"data"`
}

type getAllOrgsResponse struct {
	Data []models.Organization `json:"data"`
}

type getOrgAndmessage struct {
	Message string              `json:"message"`
	Data    models.Organization `json:"data"`
}

type getAllStrategysResponse struct {
	Data []models.Strategy `json:"data"`
}

type getStratAndmessage struct {
	Message string          `json:"message"`
	Data    models.Strategy `json:"data"`
}

type getAllPostsResponse struct {
	Data []models.Post `json:"data"`
}

type getPostAndmessage struct {
	Message string      `json:"message"`
	Data    models.Post `json:"data"`
}

type getAllGoalsResponse struct {
	Data []models.Goal `json:"data"`
}

type getGoalAndmessage struct {
	Message string      `json:"message"`
	Data    models.Goal `json:"data"`
}

type getAllEmplPostsResponse struct {
	Data []models.Employee_Post `json:"data"`
}

type getEmplPostAndmessage struct {
	Message string               `json:"message"`
	Data    models.Employee_Post `json:"data"`
}

type getAllEmployeesResponse struct {
	Data []models.Employee `json:"data"`
}

type getEmployeeAndmessage struct {
	Message string          `json:"message"`
	Data    models.Employee `json:"data"`
}

type getAllTasksResponse struct {
	Data []models.Task `json:"data"`
}

type getTaskAndmessage struct {
	Message string      `json:"message"`
	Data    models.Task `json:"data"`
}

type getAllUsersResponse struct {
	Data []models.User `json:"data"`
}

type getUserAndmessage struct {
	Message string      `json:"message"`
	Data    models.User `json:"data"`
}
