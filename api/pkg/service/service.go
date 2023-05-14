package service

import (
	"api/models"
	"api/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	Authorize(login, password string, idempl int) (string, error)
	EmailCheck(email, code string) (string, error)
	GetEmployee(email string) (int, error)
	ParseToken(token string) (int, error)
	GetAll(id int) (models.UserAllData, error)
}

type Organization interface {
	Create(organization models.Organization) (models.Organization, error)
	GetAll() ([]models.Organization, error)
	GetById(id int) (models.Organization, error)
	GetDirector(id int) (models.Employee, error)
	GetByKey(key string) (models.Organization, error)
	Delete(id int) error
	Update(id int, organization models.Organization) (models.Organization, error)
	UpdateDirector(director models.Employee, oldruc, newruc int) (string, error)
}

type Finances_Operations interface {
	Create(operation models.Finances_Operations, idorg int) (models.Finances_Operations, error)
	GetAll(idorg int) ([]models.Finances_Operations, error)
	GetById(id int, idorg int) (models.Finances_Operations, error)
	Delete(id int, idorg int) error
	Update(id int, operation models.Finances_Operations, idorg int) (models.Finances_Operations, error)
}

type Strategy interface {
	Create(strategy models.Strategy, idorg int) (models.Strategy, error)
	GetAll(idorg int) ([]models.Strategy, error)
	GetById(id int, idorg int) (models.Strategy, error)
	Delete(id int, idorg int) error
	Update(id int, strategy models.Strategy, idorg int) (models.Strategy, error)
}

type Department interface {
	Create(department models.Department, idorg int) (models.Department, error)
	GetAll(idorg int) ([]models.Department, error)
	GetById(id int, idorg int) (models.Department, error)
	GetRucovoditel(id int) (models.Employee, error)
	Delete(id int, idorg int) error
	Update(id int, department models.Department, idorg int) (models.Department, error)
	UpdateRucovoditel(ruc models.Employee, oldruc, newruc int) (string, error)
}

type Post interface {
	Create(post models.Post, iddep int, idorg int) (models.Post, error)
	GetAll(iddep int, idorg int) ([]models.Post, error)
	GetOrganizationAll(idorg int) ([]models.Post, error)
	GetById(id int, iddep int, idorg int) (models.Post, error)
	Delete(id int, iddep int, idorg int) error
	Update(id int, post models.Post, iddep int, idorg int) (models.Post, error)
}

type Goal interface {
	Create(goal models.Goal, iddep int, idorg int) (models.Goal, error)
	GetAll(iddep int, idorg int) ([]models.Goal, error)
	GetById(id int, iddep int, idorg int) (models.Goal, error)
	Delete(id int, iddep int, idorg int) error
	Update(id int, goal models.Goal, iddep int, idorg int) (models.Goal, error)
}

type EmplPost interface {
	Create(emplpost models.Employee_Post, idpost int, iddep int) (models.Employee_Post, error)
	GetAll(idpost int, iddep int) ([]models.Employee_Post, error)
	GetById(id int, idpost int, iddep int) (models.Employee_Post, error)
	Delete(id int, iddep int) error
}

type Employee interface {
	Create(employee models.Employee, iddep int, idorg int) (models.Employee, error)
	GetAll(iddep int, idorg int) ([]models.Employee, error)
	GetOrganizationAll(idorg int) ([]models.Employee, error)
	GetById(id int, iddep int, idorg int) (models.Employee, error)
	Delete(id int, iddep int, idorg int) error
	Update(id int, employee models.Employee, iddep int, idorg int) (models.Employee, error)
}

type Task interface {
	Create(task models.Task, idempl int, iddep int) (models.Task, error)
	GetAll(idempl int, iddep int) ([]models.Task, error)
	GetById(id int, idempl int, iddep int) (models.Task, error)
	Delete(id int, idempl int, iddep int) error
	Update(id int, task models.Task, idempl int, iddep int) (models.Task, error)
}

type User interface {
	Create(user models.User, idempl int, iddep int) (models.User, error)
	GetAll(idempl int, iddep int) ([]models.User, error)
	GetById(id int, idempl int, iddep int) (models.User, error)
	Delete(id int, idempl int, iddep int) error
	Update(id int, user models.User, idempl int, iddep int) (models.User, error)
	UpdateProfile(id int, user models.UserAllData) (models.User, error)
	UpdateEmployee(empl models.Employee) (models.Employee, error)
	UpdatePassword(id int, oldpassword string, newpassword string, userall models.UserAllData) (string, error)

	CreateTask(task models.Task) (models.Task, error)
	GetAllTask(idempl int) ([]models.Task, error)
	GetTaskById(id int) (models.Task, error)
	DeleteTask(id int) error
	DeleteLogTask(id int) error
	UpdateTask(id int, task models.Task) (models.Task, error)

	CreateGoal(goal models.Goal) (models.Goal, error)
	GetAllGoal(idempl int) ([]models.Goal, error)
	GetGoalById(id int) (models.Goal, error)
	DeleteGoal(id int) error
	DeleteLogGoal(id int) error
	UpdateGoal(id int, goal models.Goal) (models.Goal, error)

	CreateStrat(strat models.Strategy) (models.Strategy, error)
	GetAllStrat(idempl int) ([]models.Strategy, error)
	GetStratById(id int) (models.Strategy, error)
	DeleteStrat(id int) error
	DeleteLogStrat(id int) error
	UpdateStrat(id int, strat models.Strategy) (models.Strategy, error)

	SendMail(email, title, message string) (string, error)
}

type Service struct {
	Authorization
	Organization
	Strategy
	Finances_Operations
	User
	Task
	Employee
	Post
	EmplPost
	Department
	Goal
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:       NewAuthService(repos.Authorization),
		Organization:        NewOrganizationService(repos.Organization, repos.Employee),
		Strategy:            NewStrategyService(repos.Strategy, repos.Organization),
		Finances_Operations: NewOperationService(repos.Finances_Operations, repos.Organization),
		Department:          NewDepartmentService(repos.Department, repos.Organization, repos.Employee),
		Post:                NewPostService(repos.Post, repos.Department),
		Employee:            NewEmployeeService(repos.Employee, repos.Department),
		EmplPost:            NewEmpl_postService(repos.Empl_post, repos.Employee),
		Goal:                NewGoalService(repos.Goal, repos.Department),
		Task:                NewTaskService(repos.Task, repos.Employee),
		User:                NewUserService(repos.User, repos.Employee),
	}
}
