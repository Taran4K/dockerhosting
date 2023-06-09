package repository

import (
	"api/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

const (
	PW_SALT_BYTES = 32
	PW_HASH_BYTES = 32
)

type OrganizationPostgres struct {
	db *sqlx.DB
}

func NewOrganizationPostgres(db *sqlx.DB) *OrganizationPostgres {
	return &OrganizationPostgres{db: db}
}

func (r *OrganizationPostgres) Create(organization models.Organization) (models.Organization, error) {
	var org models.Organization
	tx, err := r.db.Begin()
	if err != nil {
		return models.Organization{}, err
	}

	var organizationId int
	info := []byte(organization.INN + organization.Name)
	key, err := bcrypt.GenerateFromPassword(info, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	keystring := string(key)
	keystring = strings.Replace(keystring, "/", "", -1)

	query := fmt.Sprintf("SELECT insert_organization($1, $2, $3, $4, $5, $6)")

	row := tx.QueryRow(query, organization.Name, organization.Addres, organization.INN, organization.Budget, organization.Date_Foundation, keystring)

	err = row.Scan(&organizationId)
	log.Print(err)

	if err != nil {
		tx.Rollback()
		return models.Organization{}, err
	}

	tx.Commit()

	org, err = r.GetById(organizationId)
	if err != nil {
		return models.Organization{}, err
	}

	return org, err
}

func (r *OrganizationPostgres) GetAll() ([]models.Organization, error) {
	var orgs []models.Organization
	query := fmt.Sprintf("SELECT * FROM %s", apiOrganizationTable)

	err := r.db.Select(&orgs, query)

	return orgs, err
}

func (r *OrganizationPostgres) GetById(id int) (models.Organization, error) {
	var org models.Organization
	query := fmt.Sprintf("SELECT * FROM %s WHERE id_organization=$1", apiOrganizationTable)

	err := r.db.Get(&org, query, id)

	return org, err
}

func (r *OrganizationPostgres) GetDirector(empls []models.Employee) (models.Employee, error) {
	for i := 0; i < len(empls); i++ {
		var usr models.User
		query := fmt.Sprintf("SELECT * FROM \"%s\" WHERE employee_id=$1", usersTable)

		err := r.db.Get(&usr, query, empls[i].ID_Employee)

		if usr.Roles_ID == 3 {
			return empls[i], err
		}
	}
	return models.Employee{}, nil
}

func (r *OrganizationPostgres) GetByKey(key string) (models.Organization, error) {
	print(key)
	var org models.Organization
	query := fmt.Sprintf("SELECT * FROM %s WHERE auth_key=$1", apiOrganizationTable)

	err := r.db.Get(&org, query, key)

	return org, err
}

func (r *OrganizationPostgres) Delete(id int) error {
	idorg := "id_organization"
	query := fmt.Sprintf("SELECT delete_row($1, $2, $3)")

	_, err := r.db.Exec(query, apiOrganizationTable, idorg, id)

	return err
}

func (r *OrganizationPostgres) Update(id int, organization models.Organization) (models.Organization, error) {
	var org models.Organization

	query := fmt.Sprintf("SELECT update_organization($1, $2, $3, $4, $5, $6)")

	_, err := r.db.Exec(query, id, organization.Name, organization.Addres, organization.INN, organization.Budget, organization.Date_Foundation)

	org, _ = r.GetById(id)

	return org, err
}

func (r *OrganizationPostgres) UpdateDirector(olddir, newdir int) (string, error) {
	query := fmt.Sprintf("SELECT swap_roles($1, $2)")

	id, err := r.db.Exec(query, newdir, 3)
	fmt.Print(id)
	if id == nil {
		return "", err
	}

	query1 := fmt.Sprintf("SELECT swap_roles($1, $2)")

	id, err1 := r.db.Exec(query1, olddir, 1)
	fmt.Print(id)
	if id == nil {
		return "", err1
	}

	return "Успешное изменение данных", err
}

func (r *OrganizationPostgres) CreateDirector(newdir int) (string, error) {
	query := fmt.Sprintf("SELECT swap_roles($1, $2)")

	_, err := r.db.Exec(query, newdir, 3)
	if err != nil {
		return "", err
	}

	return "Успешное изменение данных", err
}
