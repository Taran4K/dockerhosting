package repository

import (
	"api/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(user models.User, idempl int) (models.User, error) {
	var usr models.User
	tx, err := r.db.Begin()
	if err != nil {
		return models.User{}, err
	}

	var UserId int
	query := fmt.Sprintf("SELECT insert_user($1, $2, $3)")

	row := tx.QueryRow(query, user.Login, user.Password, idempl)

	err = row.Scan(&UserId)
	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}
	tx.Commit()

	usr, err = r.GetById(UserId, idempl)
	if err != nil {
		return models.User{}, err
	}

	return usr, err
}

func (r *UserPostgres) GetAll(idempl int) ([]models.User, error) {
	var usr []models.User
	query := fmt.Sprintf("SELECT * FROM \"%s\" WHERE employee_id=$1", usersTable)

	err := r.db.Select(&usr, query, idempl)

	return usr, err
}

func (r *UserPostgres) GetById(id int, idempl int) (models.User, error) {
	var User models.User
	query := fmt.Sprintf("SELECT * FROM \"%s\" WHERE id_User=$1 AND employee_id=$2", usersTable)

	err := r.db.Get(&User, query, id, idempl)

	return User, err
}

func (r *UserPostgres) Delete(id int, idempl int) error {
	iduser := "id_User"
	query := fmt.Sprintf("SELECT delete_row($1, $2, $3)")

	_, err := r.db.Exec(query, usersTable, iduser, id)

	return err
}

func (r *UserPostgres) Update(id int, User models.User, idempl int) (models.User, error) {
	var usr models.User

	query := fmt.Sprintf("SELECT update_user($1, $2, $3, $4)")

	_, err := r.db.Exec(query, id, User.Login, User.Password, idempl)

	usr, _ = r.GetById(id, idempl)

	return usr, err
}

func (r *UserPostgres) UpdateProfile(id int, User models.UserAllData) (models.User, error) {
	var usr models.User

	query := fmt.Sprintf("SELECT update_user($1, $2, $3, $4)")

	_, err := r.db.Exec(query, id, User.User.Login, User.User.Password, User.User.Employee_ID)

	if err != nil {
		return models.User{}, err
	}

	query2 := fmt.Sprintf("SELECT update_employee($1, $2, $3, $4, $5, $6, $7, $8, $9)")

	_, err2 := r.db.Exec(query2, User.Employee.ID_Employee, User.Employee.Surname, User.Employee.Name, User.Employee.SecondName, User.Employee.Date_Birth, User.Employee.SeriaPasp, User.Employee.NumberPasp, User.Employee.Email, User.Employee.Department_ID)

	usr, _ = r.GetById(id, User.User.Employee_ID)

	return usr, err2
}

func (r *UserPostgres) UpdateEmployee(empl models.Employee) (models.Employee, error) {
	query := fmt.Sprintf("SELECT update_employee($1, $2, $3, $4, $5, $6, $7, $8, $9)")

	_, err := r.db.Exec(query, empl.ID_Employee, empl.Surname, empl.Name, empl.SecondName, empl.Date_Birth, empl.SeriaPasp, empl.NumberPasp, empl.Email, empl.Department_ID)

	return empl, err
}

func (r *UserPostgres) UpdatePassword(id int, newpass string) (string, error) {
	query := fmt.Sprintf("update \"%s\" set Password = $1 where id_user = $2", usersTable)

	_, err := r.db.Exec(query, newpass, id)

	if err != nil {
		return "", err
	}

	return "Успешное сохранение", err
}

func (r *UserPostgres) CreateTask(Task models.Task) (models.Task, error) {
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

	task, err := r.GetTaskById(TaskId)
	if err != nil {
		return models.Task{}, err
	}

	return task, err
}

func (r *UserPostgres) GetAllTask(idempl int) ([]models.Task, error) {
	var Tasks []models.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE employee_id=$1 AND log_del=false", apiTaskTable)

	err := r.db.Select(&Tasks, query, idempl)

	return Tasks, err
}

func (r *UserPostgres) GetTaskById(id int) (models.Task, error) {
	var task models.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE id_Task=$1 AND log_del=false", apiTaskTable)

	err := r.db.Get(&task, query, id)

	return task, err
}

func (r *UserPostgres) DeleteTask(id int) error {
	idstr := "id_Task"
	query := fmt.Sprintf("SELECT delete_row($1, $2, $3)")

	_, err := r.db.Exec(query, apiTaskTable, idstr, id)

	return err
}

func (r *UserPostgres) DeleteLogTask(id int) error {
	idtask := "id_task"
	query := fmt.Sprintf("SELECT delete_logrow($1, $2, $3)")

	_, err := r.db.Exec(query, apiTaskTable, idtask, id)

	return err
}

func (r *UserPostgres) UpdateTask(id int, Task models.Task) (models.Task, error) {
	var org models.Task

	query := fmt.Sprintf("SELECT update_SGT($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")

	if Task.Done == true {
		datetemp := time.Now().Add(time.Hour * time.Duration(3)).Format("2006-01-02 15:04:05")
		_, err := r.db.Exec(query, id, Task.Name, Task.Description, Task.Date_start, Task.Date_end, Task.Done, datetemp, Task.Employee_ID, foreignkeyTask, apiTaskTable, primarykeyTask)
		org, _ = r.GetTaskById(id)

		return org, err
	} else {
		_, err := r.db.Exec(query, id, Task.Name, Task.Description, Task.Date_start, Task.Date_end, Task.Done, time.Date(0, 0, 0, 0, 0, 0, 0, time.Local), Task.Employee_ID, foreignkeyTask, apiTaskTable, primarykeyTask)
		org, _ = r.GetTaskById(id)

		return org, err
	}
}

func (r *UserPostgres) CreateGoal(goal models.Goal) (models.Goal, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return models.Goal{}, err
	}
	var GoalId int
	query := fmt.Sprintf("SELECT insert_SGT($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")

	if goal.Done == true {
		datetemp := time.Now().Add(time.Hour * time.Duration(3)).Format("2006-01-02 15:04:05")
		var row = tx.QueryRow(query, goal.Name, goal.Description, goal.Date_start, goal.Date_end, goal.Done, datetemp, goal.Department_ID, foreignkeyGoal, apiGoalTable, primarykeyGoal)
		err = row.Scan(&GoalId)
		if err != nil {
			tx.Rollback()
			return models.Goal{}, err
		}
		tx.Commit()
	} else {
		var row = tx.QueryRow(query, goal.Name, goal.Description, goal.Date_start, goal.Date_end, goal.Done, time.Date(0, 0, 0, 0, 0, 0, 0, time.Local), goal.Department_ID, foreignkeyGoal, apiGoalTable, primarykeyGoal)
		err = row.Scan(&GoalId)
		if err != nil {
			tx.Rollback()
			return models.Goal{}, err
		}
		tx.Commit()
	}

	task, err := r.GetGoalById(GoalId)
	if err != nil {
		return models.Goal{}, err
	}

	return task, err
}

func (r *UserPostgres) GetAllGoal(idempl int) ([]models.Goal, error) {
	var Goals []models.Goal
	query := fmt.Sprintf("SELECT * FROM %s WHERE department_id=$1 AND log_del=false", apiGoalTable)

	err := r.db.Select(&Goals, query, idempl)

	return Goals, err
}

func (r *UserPostgres) GetGoalById(id int) (models.Goal, error) {
	var org models.Goal
	query := fmt.Sprintf("SELECT * FROM %s WHERE id_Goal=$1 AND log_del=false", apiGoalTable)

	err := r.db.Get(&org, query, id)

	return org, err
}

func (r *UserPostgres) DeleteGoal(id int) error {
	idgoal := "id_Goal"
	query := fmt.Sprintf("SELECT delete_row($1, $2, $3)")

	_, err := r.db.Exec(query, apiGoalTable, idgoal, id)

	return err
}

func (r *UserPostgres) DeleteLogGoal(id int) error {
	idgoal := "id_goal"
	query := fmt.Sprintf("SELECT delete_logrow($1, $2, $3)")

	_, err := r.db.Exec(query, apiGoalTable, idgoal, id)

	return err
}

func (r *UserPostgres) UpdateGoal(id int, goal models.Goal) (models.Goal, error) {
	var org models.Goal

	query := fmt.Sprintf("SELECT update_SGT($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")

	if goal.Done == true {
		datetemp := time.Now().Add(time.Hour * time.Duration(3)).Format("2006-01-02 15:04:05")
		_, err := r.db.Exec(query, id, goal.Name, goal.Description, goal.Date_start, goal.Date_end, goal.Done, datetemp, goal.Department_ID, foreignkeyGoal, apiGoalTable, primarykeyGoal)
		org, _ = r.GetGoalById(id)

		return org, err
	} else {
		_, err := r.db.Exec(query, id, goal.Name, goal.Description, goal.Date_start, goal.Date_end, goal.Done, time.Date(0, 0, 0, 0, 0, 0, 0, time.Local), goal.Department_ID, foreignkeyGoal, apiGoalTable, primarykeyGoal)
		org, _ = r.GetGoalById(id)

		return org, err
	}
}

func (r *UserPostgres) CreateStrat(strat models.Strategy) (models.Strategy, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return models.Strategy{}, err
	}
	var stratId int
	query := fmt.Sprintf("SELECT insert_SGT($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")

	if strat.Done == true {
		datetemp := time.Now().Add(time.Hour * time.Duration(3)).Format("2006-01-02 15:04:05")
		var row = tx.QueryRow(query, strat.Name, strat.Description, strat.Date_start, strat.Date_end, strat.Done, datetemp, strat.Organization_ID, foreignkey, apiStrategyTable, primarykey)
		err = row.Scan(&stratId)
		if err != nil {
			tx.Rollback()
			return models.Strategy{}, err
		}
		tx.Commit()
	} else {
		var row = tx.QueryRow(query, strat.Name, strat.Description, strat.Date_start, strat.Date_end, strat.Done, time.Date(0, 0, 0, 0, 0, 0, 0, time.Local), strat.Organization_ID, foreignkey, apiStrategyTable, primarykey)
		err = row.Scan(&stratId)
		if err != nil {
			tx.Rollback()
			return models.Strategy{}, err
		}
		tx.Commit()
	}

	stratget, err := r.GetStratById(stratId)
	if err != nil {
		return models.Strategy{}, err
	}

	return stratget, err
}

func (r *UserPostgres) GetAllStrat(idempl int) ([]models.Strategy, error) {
	var strats []models.Strategy
	query := fmt.Sprintf("SELECT * FROM %s WHERE organization_id=$1 AND log_del=false", apiStrategyTable)

	err := r.db.Select(&strats, query, idempl)

	return strats, err
}

func (r *UserPostgres) GetStratById(id int) (models.Strategy, error) {
	var org models.Strategy
	query := fmt.Sprintf("SELECT * FROM %s WHERE id_Strategy=$1 AND log_del=false", apiStrategyTable)

	err := r.db.Get(&org, query, id)

	return org, err
}

func (r *UserPostgres) DeleteStrat(id int) error {
	idstr := "id_strategy"
	query := fmt.Sprintf("SELECT delete_row($1, $2, $3)")

	_, err := r.db.Exec(query, apiStrategyTable, idstr, id)

	return err
}

func (r *UserPostgres) DeleteLogStrat(id int) error {
	idstr := "id_strategy"
	query := fmt.Sprintf("SELECT delete_logrow($1, $2, $3)")

	_, err := r.db.Exec(query, apiStrategyTable, idstr, id)

	return err
}

func (r *UserPostgres) UpdateStrat(id int, strat models.Strategy) (models.Strategy, error) {
	var org models.Strategy

	query := fmt.Sprintf("SELECT update_SGT($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")

	if strat.Done == true {
		datetemp := time.Now().Add(time.Hour * time.Duration(3)).Format("2006-01-02 15:04:05")
		_, err := r.db.Exec(query, id, strat.Name, strat.Description, strat.Date_start, strat.Date_end, strat.Done, datetemp, strat.Organization_ID, foreignkey, apiStrategyTable, primarykey)
		org, _ = r.GetStratById(id)

		return org, err
	} else {
		_, err := r.db.Exec(query, id, strat.Name, strat.Description, strat.Date_start, strat.Date_end, strat.Done, time.Date(0, 0, 0, 0, 0, 0, 0, time.Local), strat.Organization_ID, foreignkey, apiStrategyTable, primarykey)
		org, _ = r.GetStratById(id)

		return org, err
	}
}
