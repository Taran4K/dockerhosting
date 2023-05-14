package repository

import (
	"api/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (r *PostPostgres) Create(Post models.Post, iddep int) (models.Post, error) {
	var post models.Post
	tx, err := r.db.Begin()
	if err != nil {
		return models.Post{}, err
	}

	var PostId int
	query := fmt.Sprintf("SELECT insert_Post($1, $2, $3)")

	row := tx.QueryRow(query, Post.Name, Post.Salary, iddep)

	err = row.Scan(&PostId)
	if err != nil {
		tx.Rollback()
		return models.Post{}, err
	}
	tx.Commit()

	post, err = r.GetById(PostId, iddep)
	if err != nil {
		return models.Post{}, err
	}

	return post, err
}

func (r *PostPostgres) GetAll(iddep int) ([]models.Post, error) {
	var Post []models.Post
	query := fmt.Sprintf("SELECT * FROM %s WHERE department_id=$1", apiPostTable)

	err := r.db.Select(&Post, query, iddep)

	return Post, err
}

func (r *PostPostgres) GetOrganizationAll(idorg int) ([]models.Post, error) {
	var deps []models.Department
	query := fmt.Sprintf("SELECT * FROM %s WHERE organization_id=$1", apiDepartmentTable)

	err := r.db.Select(&deps, query, idorg)

	var posts []models.Post
	for i := 0; i < len(deps); i++ {
		var tempposts []models.Post
		query2 := fmt.Sprintf("SELECT * FROM %s WHERE department_id=$1", apiPostTable)

		err2 := r.db.Select(&tempposts, query2, deps[i].ID_Department)
		if err2 != nil {
			return []models.Post{}, err
		}
		posts = append(posts, tempposts...)
	}

	return posts, err
}

func (r *PostPostgres) GetById(id int, iddep int) (models.Post, error) {
	var Post models.Post
	query := fmt.Sprintf("SELECT * FROM %s WHERE id_Post=$1 AND department_id=$2", apiPostTable)

	err := r.db.Get(&Post, query, id, iddep)

	return Post, err
}

func (r *PostPostgres) Delete(id int, iddep int) error {
	idpost := "id_Post"
	query := fmt.Sprintf("SELECT delete_row($1, $2, $3)")

	_, err := r.db.Exec(query, apiPostTable, idpost, id)

	return err
}

func (r *PostPostgres) Update(id int, Post models.Post, iddep int) (models.Post, error) {
	var post models.Post

	query := fmt.Sprintf("SELECT update_Post($1, $2, $3, $4)")

	_, err := r.db.Exec(query, id, Post.Name, Post.Salary, Post.Department_ID)

	post, _ = r.GetById(id, Post.Department_ID)

	return post, err
}
