package repository

import (
	"api/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var UserId int
	query := fmt.Sprintf("SELECT insert_user($1, $2, $3)")

	row := tx.QueryRow(query, user.Login, user.Password, user.Employee_ID)

	err = row.Scan(&UserId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()

	return UserId, err
}

func (r *AuthPostgres) GetUser(login, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM \"%s\" WHERE login=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}

func (r *AuthPostgres) GetAll(id int) (models.UserAllData, error) {
	var user models.User
	query1 := fmt.Sprintf("SELECT * FROM \"%s\" WHERE id_user=$1", usersTable)
	r.db.Get(&user, query1, id)

	var employee models.Employee
	query2 := fmt.Sprintf("SELECT * FROM %s WHERE id_employee=$1", apiEmployeeTable)
	r.db.Get(&employee, query2, user.Employee_ID)

	var department models.Department
	query3 := fmt.Sprintf("SELECT * FROM \"%s\" WHERE id_department=$1", apiDepartmentTable)
	r.db.Get(&department, query3, employee.Department_ID)

	var organization models.Organization
	query4 := fmt.Sprintf("SELECT * FROM \"%s\" WHERE id_organization=$1", apiOrganizationTable)
	err := r.db.Get(&organization, query4, department.Organization_ID)

	if err != nil {
		return models.UserAllData{}, err
	}

	var userall models.UserAllData

	userall.User = user
	userall.Employee = employee
	userall.Department = department
	userall.Organization = organization

	return userall, err
}

func (r *AuthPostgres) GetEmployee(email string) (int, error) {
	var employeeid int
	query := fmt.Sprintf("SELECT update_EmailVerification($1)")
	err := r.db.Get(&employeeid, query, email)

	return employeeid, err
}
