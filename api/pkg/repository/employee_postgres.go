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

func (r *EmployeePostgres) GetDep(idempl int) (models.Department, error) {
	var department models.Department
	var depid int
	query := fmt.Sprintf("SELECT department_id FROM %s WHERE id_employee=$1", apiEmployeeTable)

	err := r.db.Get(&depid, query, idempl)
	if err != nil {
		return models.Department{}, err
	}

	query2 := fmt.Sprintf("SELECT * FROM %s WHERE id_department=$1", apiDepartmentTable)

	err2 := r.db.Get(&department, query2, depid)

	return department, err2
}

func (r *EmployeePostgres) GetOrg(idempl int) (models.Organization, error) {
	var org models.Organization
	var depid, orgid int
	query := fmt.Sprintf("SELECT department_id FROM %s WHERE id_employee=$1", apiEmployeeTable)

	err := r.db.Get(&depid, query, idempl)
	if err != nil {
		return models.Organization{}, err
	}

	query2 := fmt.Sprintf("SELECT organization_id FROM %s WHERE id_department=$1", apiDepartmentTable)

	err2 := r.db.Get(&orgid, query2, depid)
	if err2 != nil {
		return models.Organization{}, err2
	}

	query3 := fmt.Sprintf("SELECT * FROM %s WHERE id_organization=$1", apiOrganizationTable)

	err3 := r.db.Get(&org, query3, orgid)

	return org, err3
}

func (r *EmployeePostgres) GetOrganizationAll(idorg int) ([]models.Employee, error) {
	var deps []models.Department
	query := fmt.Sprintf("SELECT * FROM %s WHERE organization_id=$1", apiDepartmentTable)

	err := r.db.Select(&deps, query, idorg)

	var empls []models.Employee
	for i := 0; i < len(deps); i++ {
		var tempempls []models.Employee
		query2 := fmt.Sprintf("SELECT * FROM %s WHERE department_id=$1", apiEmployeeTable)

		err2 := r.db.Select(&tempempls, query2, deps[i].ID_Department)
		if err2 != nil {
			return []models.Employee{}, err
		}
		empls = append(empls, tempempls...)
	}

	return empls, err
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
