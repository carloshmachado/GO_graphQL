package database

import (
	"database/sql"

	"github.com/dgryski/trifles/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryDB  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description, categoryDB string) (*Course, error) {
	id := uuid.UUIDv4()
	_, err := c.db.Exec(`INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)`,
		id, name, description, categoryDB)
	if err != nil {
		return nil, err
	}

	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryDB:  categoryDB,
	}, nil
}

func (c *Course) FindAll() ([]*Course, error) {
	rows, err := c.db.Query(`SELECT id, name, description, category_id FROM courses`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []*Course{}
	for rows.Next() {
		var id, name, description, categoryDB string
		if err := rows.Scan(&id, &name, &description, &categoryDB); err != nil {
			return nil, err
		}
		courses = append(courses, &Course{ID: id, Name: name, Description: description, CategoryDB: categoryDB})
	}
	return courses, nil
}

func (c *Course) FindCategoryByID(categoryID string) ([]*Course, error) {
	rows, err := c.db.Query(`SELECT id, name, description, category_id FROM courses WHERE category_id = $1`, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []*Course{}
	for rows.Next() {
		var id, name, description, categoryDB string
		if err := rows.Scan(&id, &name, &description, &categoryDB); err != nil {
			return nil, err
		}
		courses = append(courses, &Course{ID: id, Name: name, Description: description, CategoryDB: categoryDB})
	}
	return courses, nil
}
