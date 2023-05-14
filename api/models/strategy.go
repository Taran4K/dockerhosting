package models

type Strategy struct {
	ID_Strategy     int    `json:"id_strategy" db:"id_strategy"`
	Name            string `json:"name" db:"name" binding:"required"`
	Description     string `json:"description" db:"description" binding:"required"`
	Date_start      string `json:"date_start" db:"date_start" binding:"required"`
	Date_end        string `json:"date_end" db:"date_end" binding:"required"`
	Done            bool   `json:"done" db:"done"`
	Date_done       string `json:"date_done" db:"date_done"`
	Log_del         bool   `json:"log_del" db:"log_del"`
	Organization_ID int    `json:"organization_id" db:"organization_id"`
}
