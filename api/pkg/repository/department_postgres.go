package repository

import (
	"api/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DepartmentPostgres struct {
	db *sqlx.DB
}

func NewDepartmentPostgres(db *sqlx.DB) *DepartmentPostgres {
	return &DepartmentPostgres{db: db}
}

func (r *DepartmentPostgres) Create(department models.Department, idorg int) (models.Department, error) {
	var opr models.Department
	tx, err := r.db.Begin()
	if err != nil {
		return models.Department{}, err
	}

	var departmentId int
	query := fmt.Sprintf("SELECT insert_department($1, $2, $3)")

	row := tx.QueryRow(query, department.Name, department.Description, idorg)

	err = row.Scan(&departmentId)
	if err != nil {
		tx.Rollback()
		return models.Department{}, err
	}
	tx.Commit()

	opr, err = r.GetById(departmentId, idorg)
	if err != nil {
		return models.Department{}, err
	}

	return opr, err
}

func (r *DepartmentPostgres) GetAll(idorg int) ([]models.Department, error) {
	var department []models.Department
	query := fmt.Sprintf("SELECT * FROM %s WHERE organization_id=$1", apiDepartmentTable)

	err := r.db.Select(&department, query, idorg)

	return department, err
}

func (r *DepartmentPostgres) GetById(id int, idorg int) (models.Department, error) {
	var department models.Department
	query := fmt.Sprintf("SELECT * FROM %s WHERE id_department=$1 AND organization_id=$2", apiDepartmentTable)

	err := r.db.Get(&department, query, id, idorg)

	return department, err
}

func (r *DepartmentPostgres) GetRucovoditel(id int) (models.Employee, error) {
	var empls []models.Employee
	query := fmt.Sprintf("SELECT * FROM %s WHERE department_id=$1", apiEmployeeTable)

	err := r.db.Select(&empls, query, id)
	if err != nil {
		return models.Employee{}, err
	}
	for i := 0; i < len(empls); i++ {
		var usr models.User
		query1 := fmt.Sprintf("SELECT * FROM \"%s\" WHERE employee_id=$1", usersTable)

		err1 := r.db.Get(&usr, query1, empls[i].ID_Employee)

		if usr.Roles_ID == 2 {
			return empls[i], err1
		}
	}
	return models.Employee{}, err
}

func (r *DepartmentPostgres) Delete(id int, idorg int) error {
	iddep := "id_department"
	query := fmt.Sprintf("SELECT delete_row($1, $2, $3)")

	_, err := r.db.Exec(query, apiDepartmentTable, iddep, id)

	return err
}

func (r *DepartmentPostgres) Update(id int, department models.Department, idorg int) (models.Department, error) {
	var dep models.Department

	query := fmt.Sprintf("SELECT update_department($1, $2, $3, $4)")

	_, err := r.db.Exec(query, id, department.Name, department.Description, idorg)

	dep, _ = r.GetById(id, idorg)

	return dep, err
}

func (r *DepartmentPostgres) UpdateRucovoditel(olddir, newdir int) (string, error) {
	query := fmt.Sprintf("SELECT swap_roles($1, $2)")

	id, err := r.db.Exec(query, newdir, 2)
	if id == nil {
		return "", err
	}

	query1 := fmt.Sprintf("SELECT swap_roles($1, $2)")

	id, err1 := r.db.Exec(query1, olddir, 1)
	if id == nil {
		return "", err1
	}

	return "Успешное изменение данных", err1
}

func (r *DepartmentPostgres) CreateRucovoditel(newruc int) (string, error) {
	query := fmt.Sprintf("SELECT swap_roles($1, $2)")

	_, err := r.db.Exec(query, newruc, 2)
	if err != nil {
		return "", err
	}
	return "Успешное изменение данных", err
}
