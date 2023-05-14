package repository

import (
	"api/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type EmployeePostgres struct {
	db *sqlx.DB
}

func NewEmployeePostgres(db *sqlx.DB) *EmployeePostgres {
	return &EmployeePostgres{db: db}
}

func (r *EmployeePostgres) Create(employee models.Employee, iddep int) (models.Employee, error) {
	var empl models.Employee
	tx, err := r.db.Begin()
	if err != nil {
		return models.Employee{}, err
	}

	var EmployeeId int
	query := fmt.Sprintf("SELECT insert_employee($1, $2, $3, $4, $5, $6, $7, $8)")

	row := tx.QueryRow(query, employee.Surname, employee.Name, employee.SecondName, employee.Date_Birth, employee.SeriaPasp, employee.NumberPasp, employee.Email, iddep)

	err = row.Scan(&EmployeeId)
	if err != nil {
		tx.Rollback()
		return models.Employee{}, err
	}
	tx.Commit()

	empl, err = r.GetById(EmployeeId, iddep)
	if err != nil {
		return models.Employee{}, err
	}

	return empl, err
}

func (r *EmployeePostgres) GetAll(iddep int) ([]models.Employee, error) {
	var employee []models.Employee
	query := fmt.Sprintf("SELECT * FROM %s WHERE department_id=$1", apiEmployeeTable)

	err := r.db.Select(&employee, query, iddep)

	return employee, err
}

func (r *EmployeePostgres) GetById(id int, iddep int) (models.Employee, error) {
	var employee models.Employee
	query := fmt.Sprintf("SELECT * FROM %s WHERE id_employee=$1 AND department_id=$2", apiEmployeeTable)

	err := r.db.Get(&employee, query, id, iddep)

	return employee, err
}

func (r *EmployeePostgres) Delete(id int, iddep int) error {
	idempl := "id_Employee"
	query := fmt.Sprintf("SELECT delete_row($1, $2, $3)")

	_, err := r.db.Exec(query, apiEmployeeTable, idempl, id)

	return err
}

func (r *EmployeePostgres) Update(id int, employee models.Employee, iddep int) (models.Employee, error) {
	var empl models.Employee

	query := fmt.Sprintf("SELECT update_employee($1, $2, $3, $4, $5, $6, $7, $8, $9)")

	_, err := r.db.Exec(query, id, employee.Surname, employee.Name, employee.SecondName, employee.Date_Birth, employee.SeriaPasp, employee.NumberPasp, employee.Email, employee.Department_ID)

	empl, _ = r.GetById(id, employee.Department_ID)

	return empl, err
}
