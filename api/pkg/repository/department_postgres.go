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
