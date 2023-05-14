package repository

import (
	"api/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type GoalPostgres struct {
	db *sqlx.DB
}

func NewGoalPostgres(db *sqlx.DB) *GoalPostgres {
	return &GoalPostgres{db: db}
}

const foreignkeyGoal string = "department_id"
const primarykeyGoal string = "id_goal"

func (r *GoalPostgres) Create(Goal models.Goal, idorg int) (models.Goal, error) {

	var org models.Goal
	tx, err := r.db.Begin()
	if err != nil {
		return models.Goal{}, err
	}
	var GoalId int
	query := fmt.Sprintf("SELECT insert_SGT($1, $2, $3, $4, $5, $6, $7, $8, $9)")

	row := tx.QueryRow(query, Goal.Name, Goal.Description, Goal.Date_start, Goal.Date_end, Goal.Done, Goal.Department_ID, foreignkeyGoal, apiGoalTable, primarykeyGoal)

	err = row.Scan(&GoalId)
	if err != nil {
		tx.Rollback()
		return models.Goal{}, err
	}
	tx.Commit()

	org, err = r.GetById(GoalId, Goal.Department_ID)
	if err != nil {
		return models.Goal{}, err
	}

	return org, err
}

func (r *GoalPostgres) GetAll(idorg int) ([]models.Goal, error) {
	var Goals []models.Goal
	query := fmt.Sprintf("SELECT * FROM %s WHERE department_id=$1", apiGoalTable)

	err := r.db.Select(&Goals, query, idorg)

	return Goals, err
}

func (r *GoalPostgres) GetById(id int, idorg int) (models.Goal, error) {
	var org models.Goal
	query := fmt.Sprintf("SELECT * FROM %s WHERE id_Goal=$1 AND department_id=$2", apiGoalTable)

	err := r.db.Get(&org, query, id, idorg)

	return org, err
}

func (r *GoalPostgres) Delete(id int, idorg int) error {
	idstr := "id_Goal"
	query := fmt.Sprintf("SELECT delete_row($1, $2, $3)")

	_, err := r.db.Exec(query, apiGoalTable, idstr, id)

	return err
}

func (r *GoalPostgres) Update(id int, Goal models.Goal, idorg int) (models.Goal, error) {
	var org models.Goal

	query := fmt.Sprintf("SELECT update_SGT($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")

	_, err := r.db.Exec(query, id, Goal.Name, Goal.Description, Goal.Date_start, Goal.Date_end, Goal.Done, Goal.Department_ID, foreignkeyGoal, apiGoalTable, primarykeyGoal)

	org, _ = r.GetById(id, Goal.Department_ID)

	return org, err
}
