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

func (r *OrganizationPostgres) GetByKey(key string) (models.Organization, error) {
	var org models.Organization
	query := fmt.Sprintf("SELECT * FROM %s WHERE Auth_Key=$1", apiOrganizationTable)

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
