package handler

import (
	"api/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	c := cors.New(cors.Config{
		AllowMethods:           []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:           []string{"*"},
		AllowCredentials:       true,
		ExposeHeaders:          []string{"*"},
		MaxAge:                 6000,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             true,
		AllowOrigins:           []string{"*"},
	})
	router.Use(c)

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		orgs := api.Group("/organization")
		{
			orgs.POST("/", h.createOrganization)
			orgs.GET("/", h.getAllOrganization)
			orgs.PATCH("/:id", h.getOrganizationByKey)
			orgs.GET("/:id", h.getOrganization)
			orgs.PUT("/:id", h.updateOrganization)
			orgs.DELETE("/:id", h.deleteOrganization)

			strategy := orgs.Group(":id/strategy")
			{
				strategy.POST("/", h.createStrategy)
				strategy.GET("/", h.getAllStrategies)
				strategy.GET("/:strategy_id", h.getStrategyById)
				strategy.PUT("/:strategy_id", h.updateStrategy)
				strategy.DELETE("/:strategy_id", h.deleteStrategy)
			}

			operations := orgs.Group(":id/operation")
			{
				operations.POST("/", h.createOperation)
				operations.GET("/", h.getAllOperation)
				operations.GET("/:operation_id", h.getOperation)
				operations.PUT("/:operation_id", h.updateOperation)
				operations.DELETE("/:operation_id", h.deleteOperation)
			}

			department := orgs.Group(":id/department")
			{
				department.POST("/", h.createDepartment)
				department.GET("/", h.getAllDepartment)
				department.GET("/:department_id", h.getDepartment)
				department.PUT("/:department_id", h.updateDepartment)
				department.DELETE("/:department_id", h.deleteDepartment)

				goals := department.Group(":department_id/goal")
				{
					goals.POST("/", h.createGoal)
					goals.GET("/", h.getAllGoal)
					goals.GET("/:goal_id", h.getGoal)
					goals.PUT("/:goal_id", h.updateGoal)
					goals.DELETE("/:goal_id", h.deleteGoal)
				}

				post := department.Group(":department_id/post")
				{
					post.POST("/", h.createPost)
					post.GET("/", h.getAllPost)
					post.GET("/:post_id", h.getPost)
					post.PUT("/:post_id", h.updatePost)
					post.DELETE("/:post_id", h.deletePost)
				}

				employee := department.Group(":department_id/employee")
				{
					employee.POST("/", h.createEmployee)
					employee.GET("/", h.getAllEmployee)
					employee.GET("/:employee_id", h.getEmployee)
					employee.PUT("/:employee_id", h.updateEmployee)
					employee.DELETE("/:employee_id", h.deleteEmployee)

					task := employee.Group(":employee_id/task")
					{
						task.POST("/", h.createTask)
						task.GET("/", h.getAllTask)
						task.GET("/:task_id", h.getTask)
						task.PUT("/:task_id", h.updateTask)
						task.DELETE("/:task_id", h.deleteTask)
					}

					user := employee.Group(":employee_id/user")
					{
						user.POST("/", h.createUser)
						user.GET("/", h.getAllUser)
						user.GET("/:user_id", h.getUser)
						user.PUT("/:user_id", h.updateUser)
						user.DELETE("/:user_id", h.deleteUser)
					}

					empl_post := employee.Group(":employee_id/emplpost")
					{
						empl_post.POST("/", h.createEmployee_Post)
						empl_post.GET("/", h.getAllEmployee_Post)
						empl_post.GET("/:emplpost_id", h.getEmployee_Post)
						empl_post.DELETE("/", h.deleteEmployee_Post)
					}
				}
			}
		}
	}

	return router
}
