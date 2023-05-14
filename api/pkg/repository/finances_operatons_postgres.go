package repository

import (
	"api/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type OperationPostgres struct {
	db *sqlx.DB
}

func NewOperationPostgres(db *sqlx.DB) *OperationPostgres {
	return &OperationPostgres{db: db}
}

func (r *OperationPostgres) Create(operation models.Finances_Operations, idorg int) (models.Finances_Operations, error) {
	var opr models.Finances_Operations
	tx, err := r.db.Begin()
	if err != nil {
		return models.Finances_Operations{}, err
	}

	var operationId int
	query := fmt.Sprintf("SELECT insert_finances_operations($1, $2, $3, $4)")

	row := tx.QueryRow(query, operation.Summ, operation.Date_Operation, operation.Description, idorg)

	err = row.Scan(&operationId)
	if err != nil {
		tx.Rollback()
		return models.Finances_Operations{}, err
	}
	tx.Commit()

	opr, err = r.GetById(operationId, idorg)
	if err != nil {
		return models.Finances_Operations{}, err
	}

	return opr, err
}

func (r *OperationPostgres) GetAll(idorg int) ([]models.Finances_Operations, error) {
	var operation []models.Finances_Operations
	query := fmt.Sprintf("SELECT * FROM %s WHERE organization_id=$1", apiOperationTable)

	err := r.db.Select(&operation, query, idorg)

	return operation, err
}

func (r *OperationPostgres) GetById(id int, idorg int) (models.Finances_Operations, error) {
	var operation models.Finances_Operations
	query := fmt.Sprintf("SELECT * FROM %s WHERE id_operations=$1 AND organization_id=$2", apiOperationTable)

	err := r.db.Get(&operation, query, id, idorg)

	return operation, err
}

func (r *OperationPostgres) Delete(id int, idorg int) error {
	idstr := "id_operations"
	query := fmt.Sprintf("SELECT delete_row($1, $2, $3)")

	_, err := r.db.Exec(query, apiOperationTable, idstr, id)

	return err
}

func (r *OperationPostgres) Update(id int, operation models.Finances_Operations, idorg int) (models.Finances_Operations, error) {
	var oper models.Finances_Operations

	query := fmt.Sprintf("SELECT update_finances_operations($1, $2, $3, $4, $5)")

	_, err := r.db.Exec(query, id, operation.Summ, operation.Date_Operation, operation.Description, idorg)

	oper, _ = r.GetById(id, idorg)

	return oper, err
}
