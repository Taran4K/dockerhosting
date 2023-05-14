package repository

import (
	"api/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var UserId int
	query := fmt.Sprintf("SELECT insert_user($1, $2, $3)")

	row := tx.QueryRow(query, user.Login, user.Password, user.Employee_ID)

	err = row.Scan(&UserId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()

	return UserId, err
}

func (r *AuthPostgres) GetUser(login, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id_user FROM %s WHERE login=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}
