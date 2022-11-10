package storage

import (
	"database/sql"
	"fmt"
	"time"
)

type DBManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) *DBManager {
	return &DBManager{db: db}
}

type Blog struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetBlogsQueryParam struct {
	Author string
	Title  string
	Page   int32
	Limit  int32
}

func (b *DBManager) Create(blog *Blog) (*Blog, error) {
	query := `
		INSERT INTO blogs(
			title,
			description,
			author
		) VALUES($1, $2, $3)
		RETURNING id, title, description, author, created_at
	`

	// id, title, description, author, created_at
	row := b.db.QueryRow(
		query,
		blog.Title,
		blog.Description,
		blog.Author,
	)

	var result Blog
	err := row.Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.Author,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (b *DBManager) Get(id int64) (*Blog, error) {
	var result Blog

	query := `
		SELECT
			id, 
			title, 
			description, 
			author, 
			created_at
		FROM blogs
		WHERE id=$1
	`

	row := b.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.Author,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (b *DBManager) GetAll(params *GetBlogsQueryParam) ([]*Blog, error) {
	var blogs []*Blog

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := " WHERE true "
	if params.Author != "" {
		filter += " AND author ilike '%" + params.Author + "%' "
	}

	if params.Title != "" {
		filter += " AND title ilike '%" + params.Title + "%' "
	}

	query := `
		SELECT
			id, 
			title, 
			description, 
			author, 
			created_at
		FROM blogs
		` + filter + `
		ORDER BY created_at desc
		` + limit

	rows, err := b.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var blog Blog

		err := rows.Scan(
			&blog.ID,
			&blog.Title,
			&blog.Description,
			&blog.Author,
			&blog.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		blogs = append(blogs, &blog)
	}

	return blogs, nil
}

func (b *DBManager) Update(blog *Blog) (*Blog, error) {
	query := `
		UPDATE blogs SET
			title=$1,
			description=$2,
			author=$3
		WHERE id=$4
		RETURNING id, title, description, author, created_at
	`

	row := b.db.QueryRow(
		query,
		blog.Title,
		blog.Description,
		blog.Author,
		blog.ID,
	)

	var result Blog
	err := row.Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.Author,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (b *DBManager) Delete(id int64) error {
	query := "DELETE FROM blogs WHERE id=$1"

	result, err := b.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsCount == 0 {
		return sql.ErrNoRows
	}

	return nil
}
