package models

import "github.com/shopspring/decimal"

type Organization struct {
	Id_Organization int             `json:"id_organization" db:"id_organization"`
	Name            string          `json:"name" db:"name" binding:"required"`
	Addres          string          `json:"addres" db:"addres" binding:"required"`
	INN             string          `json:"inn" db:"inn" binding:"required"`
	Budget          decimal.Decimal `json:"budget" db:"budget" binding:"required"`
	Date_Foundation string          `json:"date_foundation" db:"date_foundation" binding:"required"`
	Auth_key        string          `json:"auth_key" db:"auth_key"`
}
