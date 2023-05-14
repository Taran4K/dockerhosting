package models

import "github.com/shopspring/decimal"

type Post struct {
	ID_Post       int             `json:"id_post" db:"id_post"`
	Name          string          `json:"name" db:"name" binding:"required"`
	Salary        decimal.Decimal `json:"salary" db:"salary" binding:"required"`
	Department_ID int             `json:"department_id" db:"department_id"`
}
