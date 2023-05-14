package repository

import (
	"api/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(user models.User, idempl int) (models.User, error) {
	var usr models.User
	tx, err := r.db.Begin()
	if err != nil {
		return models.User{}, err
	}

	var UserId int
	query := fmt.Sprintf("SELECT insert_user($1, $2, $3)")

	row := tx.QueryRow(query, user.Login, user.Password, idempl)

	err = row.Scan(&UserId)
	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}
	tx.Commit()

	usr, err = r.GetById(UserId, idempl)
	if err != nil {
		return models.User{}, err
	}

	return usr, err
}

func (r *UserPostgres) GetAll(idempl int) ([]models.User, error) {
	var usr []models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE employee_id=$1", usersTable)

	err := r.db.Select(&usr, query, idempl)

	return usr, err
}

func (r *UserPostgres) GetById(id int, idempl int) (models.User, error) {
	var User models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id_User=$1 AND employee_id=$2", usersTable)

	err := r.db.Get(&User, query, id, idempl)

	return User, err
}

func (r *UserPostgres) Delete(id int, idempl int) error {
	iduser := "id_User"
	query := fmt.Sprintf("SELECT delete_row($1, $2, $3)")

	_, err := r.db.Exec(query, usersTable, iduser, id)

	return err
}

func (r *UserPostgres) Update(id int, User models.User, idempl int) (models.User, error) {
	var usr models.User

	query := fmt.Sprintf("SELECT update_user($1, $2, $3, $4)")

	_, err := r.db.Exec(query, id, User.Login, User.Password, idempl)

	usr, _ = r.GetById(id, idempl)

	return usr, err
}
