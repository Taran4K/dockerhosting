package models

type User struct {
	Id_user     int    `json:"id_user"`
	Login       string `json:"login" db:"login" binding:"required"`
	Password    string `json:"password" db:"password" binding:"required"`
	Employee_ID int    `json:"employee_id" db:"employee_id"`
	Roles_ID    int    `json:"roles_id" db:"roles_id"`
}

type UserAllData struct {
	User         User         `json:"user" binding:"required"`
	Employee     Employee     `json:"employee" binding:"required"`
	Department   Department   `json:"department" binding:"required"`
	Organization Organization `json:"organization" binding:"required"`
}
