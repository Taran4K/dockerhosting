package repository

import (
	"api/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type StrategyPostgres struct {
	db *sqlx.DB
}

func NewStrategyPostgres(db *sqlx.DB) *StrategyPostgres {
	return &StrategyPostgres{db: db}
}

const foreignkey string = "organization_id"
const primarykey string = "id_strategy"

func (r *StrategyPostgres) Create(strategy models.Strategy, idorg int) (models.Strategy, error) {

	var org models.Strategy
	tx, err := r.db.Begin()
	if err != nil {
		return models.Strategy{}, err
	}
	var strategyId int
	query := fmt.Sprintf("SELECT insert_SGT($1, $2, $3, $4, $5, $6, $7, $8, $9)")

	row := tx.QueryRow(query, strategy.Name, strategy.Description, strategy.Date_start, strategy.Date_end, strategy.Done, idorg, foreignkey, apiStrategyTable, primarykey)

	err = row.Scan(&strategyId)
	if err != nil {
		tx.Rollback()
		return models.Strategy{}, err
	}
	tx.Commit()

	org, err = r.GetById(strategyId, idorg)
	if err != nil {
		return models.Strategy{}, err
	}

	return org, err
}

func (r *StrategyPostgres) GetAll(idorg int) ([]models.Strategy, error) {
	var strategys []models.Strategy
	query := fmt.Sprintf("SELECT * FROM %s WHERE organization_id=$1", apiStrategyTable)

	err := r.db.Select(&strategys, query, idorg)

	return strategys, err
}

func (r *StrategyPostgres) GetById(id int, idorg int) (models.Strategy, error) {
	var org models.Strategy
	query := fmt.Sprintf("SELECT * FROM %s WHERE id_strategy=$1 AND organization_id=$2", apiStrategyTable)

	err := r.db.Get(&org, query, id, idorg)

	return org, err
}

func (r *StrategyPostgres) Delete(id int, idorg int) error {
	idstr := "id_strategy"
	query := fmt.Sprintf("SELECT delete_row($1, $2, $3)")

	_, err := r.db.Exec(query, apiStrategyTable, idstr, id)

	return err
}

func (r *StrategyPostgres) Update(id int, strategy models.Strategy, idorg int) (models.Strategy, error) {
	var org models.Strategy

	query := fmt.Sprintf("SELECT update_SGT($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")

	_, err := r.db.Exec(query, id, strategy.Name, strategy.Description, strategy.Date_start, strategy.Date_end, strategy.Done, idorg, foreignkey, apiStrategyTable, primarykey)

	org, _ = r.GetById(id, idorg)

	return org, err
}
