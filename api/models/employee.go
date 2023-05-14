package models

type Employee struct {
	ID_Employee   int    `json:"id_employee" db:"id_employee"`
	Surname       string `json:"surname" db:"surname" binding:"required"`
	Name          string `json:"name" db:"name" binding:"required"`
	SecondName    string `json:"secondname" db:"secondname"`
	Date_Birth    string `json:"date_birth" db:"date_birth" binding:"required"`
	SeriaPasp     string `json:"seriapasp" db:"seriapasp" binding:"required"`
	NumberPasp    string `json:"numberpasp" db:"numberpasp" binding:"required"`
	Email         string `json:"email" db:"email" binding:"required"`
	Department_ID int    `json:"department_id" db:"department_id"`
}
