package repository

import (
	"api/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

const foreignkeyTask string = "employee_id"
const primarykeyTask string = "id_Task"

func (r *TaskPostgres) Create(Task models.Task, idorg int) (models.Task, error) {

	var org models.Task
	tx, err := r.db.Begin()
	if err != nil {
		return models.Task{}, err
	}
	var TaskId int
	query := fmt.Sprintf("SELECT insert_SGT($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")

	if Task.Done == true {
		datetemp := time.Now().Add(time.Hour * time.Duration(3)).Format("2006-01-02 15:04:05")
		var row = tx.QueryRow(query, Task.Name, Task.Description, Task.Date_start, Task.Date_end, Task.Done, datetemp, Task.Employee_ID, foreignkeyTask, apiTaskTable, primarykeyTask)
		err = row.Scan(&TaskId)
		if err != nil {
			tx.Rollback()
			return models.Task{}, err
		}
		tx.Commit()
	} else {
		var row = tx.QueryRow(query, Task.Name, Task.Description, Task.Date_start, Task.Date_end, Task.Done, time.Date(0, 0, 0, 0, 0, 0, 0, time.Local), Task.Employee_ID, foreignkeyTask, apiTaskTable, primarykeyTask)
		err = row.Scan(&TaskId)
		if err != nil {
			tx.Rollback()
			return models.Task{}, err
		}
		tx.Commit()
	}

	org, err = r.GetById(TaskId, Task.Employee_ID)
	if err != nil {
		return models.Task{}, err
	}

	return org, err
}

func (r *TaskPostgres) GetAll(idorg int) ([]models.Task, error) {
	var Tasks []models.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE employee_id=$1", apiTaskTable)

	err := r.db.Select(&Tasks, query, idorg)

	return Tasks, err
}

func (r *TaskPostgres) GetById(id int, idorg int) (models.Task, error) {
	var org models.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE id_Task=$1 AND employee_id=$2", apiTaskTable)

	err := r.db.Get(&org, query, id, &idorg)

	return org, err
}

func (r *TaskPostgres) Delete(id int, idorg int) error {
	idstr := "id_Task"
	query := fmt.Sprintf("SELECT delete_row($1, $2, $3)")

	_, err := r.db.Exec(query, apiTaskTable, idstr, id)

	return err
}

func (r *TaskPostgres) Update(id int, Task models.Task, idorg int) (models.Task, error) {
	var org models.Task

	query := fmt.Sprintf("SELECT update_SGT($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")

	if Task.Done == true {
		datetemp := time.Now().Add(time.Hour * time.Duration(3)).Format("2006-01-02 15:04:05")
		_, err := r.db.Exec(query, id, Task.Name, Task.Description, Task.Date_start, Task.Date_end, Task.Done, datetemp, Task.Employee_ID, foreignkeyTask, apiTaskTable, primarykeyTask)
		org, _ = r.GetById(id, Task.Employee_ID)

		return org, err
	} else {
		_, err := r.db.Exec(query, id, Task.Name, Task.Description, Task.Date_start, Task.Date_end, Task.Done, time.Date(0, 0, 0, 0, 0, 0, 0, time.Local), Task.Employee_ID, foreignkeyTask, apiTaskTable, primarykeyTask)
		org, _ = r.GetById(id, Task.Employee_ID)

		return org, err
	}
}
