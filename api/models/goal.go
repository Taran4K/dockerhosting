package models

type Goal struct {
	ID_Goal       int    `json:"id_goal" db:"id_goal"`
	Name          string `json:"name" db:"name" binding:"required"`
	Description   string `json:"description" db:"description" binding:"required"`
	Date_start    string `json:"date_start" db:"date_start" binding:"required"`
	Date_end      string `json:"date_end" db:"date_end" binding:"required"`
	Done          bool   `json:"done" db:"done"`
	Date_done     string `json:"date_done" db:"date_done"`
	Log_del       bool   `json:"log_del" db:"log_del"`
	Department_ID int    `json:"department_id" db:"department_id"`
}
