package models

type Employee_Post struct {
	ID_Employee_Post int `json:"id_employee_post" db:"id_employee_post"`
	Post_ID          int `json:"post_id" db:"post_id"`
	Employee_ID      int `json:"employee_id" db:"employee_id"`
}
