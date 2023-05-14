package models

import "github.com/shopspring/decimal"

type Finances_Operations struct {
	ID_Operations   int             `json:"id_operations" db:"id_operations"`
	Summ            decimal.Decimal `json:"summ" db:"summ" binding:"required"`
	Date_Operation  string          `json:"date_operation" db:"date_operation" binding:"required"`
	Description     string          `json:"description" db:"description" binding:"required"`
	Organization_ID int             `json:"organization_id" db:"organization_id"`
}
