package models

type Department struct {
	ID_Department   int    `json:"id_department" db:"id_department"`
	Name            string `json:"name" db:"name" binding:"required"`
	Description     string `json:"description" db:"description" binding:"required"`
	Organization_ID int    `json:"organization_id" db:"organization_id"`
}
