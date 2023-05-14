package repository

import (
	"api/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(login, password string) (models.User, error)
}

type Organization interface {
	Create(organization models.Organization) (models.Organization, error)
	GetAll() ([]models.Organization, error)
	GetById(id int) (models.Organization, error)
	GetByKey(key string) (models.Organization, error)
	Delete(id int) error
	Update(id int, organization models.Organization) (models.Organization, error)
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
	Delete(id int, idorg int) error
	Update(id int, department models.Department, idorg int) (models.Department, error)
}

type Post interface {
	Create(post models.Post, iddep int) (models.Post, error)
	GetAll(iddep int) ([]models.Post, error)
	GetById(id int, iddep int) (models.Post, error)
	Delete(id int, iddep int) error
	Update(id int, post models.Post, iddep int) (models.Post, error)
}

type Goal interface {
	Create(goal models.Goal, iddep int) (models.Goal, error)
	GetAll(iddep int) ([]models.Goal, error)
	GetById(id int, iddep int) (models.Goal, error)
	Delete(id int, iddep int) error
	Update(id int, goal models.Goal, iddep int) (models.Goal, error)
}

type Empl_post interface {
	Create(emplpost models.Employee_Post, idpost int) (models.Employee_Post, error)
	GetAll(idpost int) ([]models.Employee_Post, error)
	GetById(id int, idpost int) (models.Employee_Post, error)
	Delete(idempl int) error
}

type Employee interface {
	Create(employee models.Employee, iddep int) (models.Employee, error)
	GetAll(iddep int) ([]models.Employee, error)
	GetById(id int, iddep int) (models.Employee, error)
	Delete(id int, iddep int) error
	Update(id int, employee models.Employee, iddep int) (models.Employee, error)
}

type Task interface {
	Create(task models.Task, idempl int) (models.Task, error)
	GetAll(idempl int) ([]models.Task, error)
	GetById(id int, idempl int) (models.Task, error)
	Delete(id int, idempl int) error
	Update(id int, task models.Task, idempl int) (models.Task, error)
}

type User interface {
	Create(user models.User, idempl int) (models.User, error)
	GetAll(idempl int) ([]models.User, error)
	GetById(id int, idempl int) (models.User, error)
	Delete(id int, idempl int) error
	Update(id int, user models.User, idempl int) (models.User, error)
}

type Repository struct {
	Authorization
	Organization
	Strategy
	Finances_Operations
	User
	Task
	Employee
	Post
	Empl_post
	Department
	Goal
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:       NewAuthPostgres(db),
		Organization:        NewOrganizationPostgres(db),
		Strategy:            NewStrategyPostgres(db),
		Finances_Operations: NewOperationPostgres(db),
		User:                NewUserPostgres(db),
		Task:                NewTaskPostgres(db),
		Employee:            NewEmployeePostgres(db),
		Post:                NewPostPostgres(db),
		Empl_post:           NewEmployee_PostPostgres(db),
		Department:          NewDepartmentPostgres(db),
		Goal:                NewGoalPostgres(db),
	}
}
